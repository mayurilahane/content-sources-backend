BEGIN;

ALTER TABLE REPOSITORY_CONFIGURATIONS DROP COLUMN IF EXISTS failed_snapshot_count;

COMMIT;
