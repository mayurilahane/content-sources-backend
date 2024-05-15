package candlepin_client

import (
	"context"
	"fmt"

	caliri "github.com/content-services/caliri/release/v4"
	"github.com/content-services/content-sources-backend/pkg/config"
	"github.com/openlyinc/pointy"
)

func (c *cpClientImpl) fetchOwner(ctx context.Context, key string) (*caliri.OwnerDTO, error) {
	ctx, client, err := getCandlepinClient(ctx)
	if err != nil {
		return nil, err
	}

	found, httpResp, err := client.OwnerAPI.GetOwner(ctx, key).Execute()
	if httpResp != nil {
		defer httpResp.Body.Close()
	}
	if httpResp.StatusCode == 404 {
		return nil, nil
	}
	if err != nil {
		return nil, errorWithResponseBody("couldn't fetch org", httpResp, err)
	}
	return found, nil
}

func (c *cpClientImpl) CreateOwner(ctx context.Context) error {
	if !config.Get().Clients.Candlepin.DevelOrg {
		return fmt.Errorf("cannot create an org with devel org turned off")
	}

	ctx, client, err := getCandlepinClient(ctx)
	if err != nil {
		return err
	}

	found, err := c.fetchOwner(ctx, DevelOrgKey)
	if found != nil || err != nil {
		return err
	}

	_, httpResp, err := client.OwnerAPI.CreateOwner(ctx).OwnerDTO(caliri.OwnerDTO{
		DisplayName:       pointy.Pointer("ContentSourcesTest"),
		Key:               pointy.Pointer(DevelOrgKey),
		ContentAccessMode: pointy.Pointer("org_environment"),
	}).Execute()
	if httpResp != nil {
		defer httpResp.Body.Close()
	}
	if err != nil {
		return errorWithResponseBody("couldn't create org", httpResp, err)
	}
	return nil
}
