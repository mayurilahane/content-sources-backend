package external_repos

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/content-services/content-sources-backend/pkg/config"
	"github.com/content-services/content-sources-backend/pkg/dao"
	"github.com/content-services/content-sources-backend/pkg/db"
	"github.com/content-services/yummy/pkg/yum"
	"github.com/rs/zerolog/log"
)

const (
	RhCdnHost = "cdn.redhat.com"
)

// IntrospectUrl Fetch the metadata of a url and insert RPM data
//  Returns the number of new RPMs inserted system-wide and any error encountered
func IntrospectUrl(url string) (int64, error) {
	err, repo := dao.GetRepositoryDao(db.DB).FetchForUrl(url)
	rpmDao := dao.GetRpmDao(db.DB, nil)
	repoDao := dao.GetRepositoryDao(db.DB)
	if err != nil {
		return 0, err
	}

	return Introspect(repo, repoDao, rpmDao)
}

// IsRedHat returns if the url is a 'cdn.redhat.com' url
func IsRedHat(url string) bool {
	return strings.Contains(url, RhCdnHost)
}

// Introspect introspects a dao.Repository with the given RpmDao
// inserting any needed RPMs and adding and removing associations to the repository
// Returns the number of new RPMs inserted system-wide and any error encountered
func Introspect(repo dao.Repository, repoDao dao.RepositoryDao, rpm dao.RpmDao) (int64, error) {
	var (
		client http.Client
		err    error
		pkgs   []yum.Package
		repomd yum.Repomd
		total  int64
	)
	log.Debug().Msg("Introspecting " + repo.URL)

	if client, err = httpClient(IsRedHat(repo.URL)); err != nil {
		return 0, err
	}

	if repomd, err = yum.GetRepomdXML(client, repo.URL); err != nil {
		return 0, err
	}

	if repo.Revision != "" && repomd.Revision != "" && repomd.Revision == repo.Revision {
		// If repository hasn't changed, no need to update
		return 0, nil
	}

	if pkgs, err = yum.ExtractPackageData(client, repo.URL); err != nil {
		return 0, err
	}

	if total, err = rpm.InsertForRepository(repo.UUID, pkgs); err != nil {
		return 0, err
	}

	repo.Revision = repomd.Revision
	if err = repoDao.Update(repo); err != nil {
		return 0, err
	}

	return total, nil
}

// IntrospectAll introspects all repositories
//  Returns the number of new RPMs inserted system-wide and all errors encountered
func IntrospectAll() (int64, []error) {
	var errors []error
	var total int64
	var count int64
	rpmDao := dao.GetRpmDao(db.DB, nil)
	repoDao := dao.GetRepositoryDao(db.DB)
	err, repos := repoDao.List()

	if err != nil {
		return 0, []error{err}
	}
	for i := 0; i < len(repos); i++ {
		count, err = Introspect(repos[i], repoDao, rpmDao)
		total += count
		if err != nil {
			errors = append(errors, err)
		}
	}
	return total, errors
}

func httpClient(useCert bool) (http.Client, error) {
	timeout := 90 * time.Second
	if useCert {
		var (
			cert   *tls.Certificate
			caCert []byte
			err    error
		)

		cert = config.Get().Certs.CdnCertPair
		if cert == nil {
			return http.Client{}, errors.New("no certificate loaded")
		}
		if caCert, err = LoadCA(); err != nil {
			return http.Client{}, err
		}
		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(caCert)

		tlsConfig := &tls.Config{
			Certificates: []tls.Certificate{*cert},
			RootCAs:      caCertPool,
		}

		transport := &http.Transport{TLSClientConfig: tlsConfig, ResponseHeaderTimeout: timeout}
		return http.Client{Transport: transport, Timeout: timeout}, nil
	} else {
		return http.Client{}, nil
	}
}