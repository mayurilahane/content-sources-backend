BEGIN;

ALTER TABLE uploads ADD COLUMN IF NOT EXISTS size int DEFAULT 0 NOT NULL;

COMMIT;
