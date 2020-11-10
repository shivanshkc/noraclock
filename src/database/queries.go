package database

import (
	"fmt"
	"noraclock/v2/src/configs"
)

func createUpdatedAtFunc() string {
	return `
CREATE OR REPLACE FUNCTION "setUpdatedAt"()
RETURNS TRIGGER AS $$
BEGIN
  NEW."updatedAt" = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
`
}

func createMemoryTableIfNotExists() string {
	return fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %s (
	"id"         VARCHAR     PRIMARY KEY NOT NULL,
	"title"      VARCHAR     NOT NULL,
	"body"       TEXT     	 NOT NULL,
	"createdAt"  TIMESTAMP   NOT NULL DEFAULT NOW(),
	"updatedAt"  TIMESTAMP   NOT NULL DEFAULT NOW()
)
`,
		configs.Postgres.MemoryTableName,
	)
}

func createUpdatedAtTrigger() string {
	return fmt.Sprintf(`
DROP TRIGGER IF EXISTS "setUpdatedAt" ON %s;
CREATE TRIGGER "setUpdatedAt"
BEFORE UPDATE ON %s
FOR EACH ROW
EXECUTE PROCEDURE "setUpdatedAt"();
`,
		configs.Postgres.MemoryTableName,
		configs.Postgres.MemoryTableName,
	)
}
