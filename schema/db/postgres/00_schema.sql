-- migrate up

CREATE SCHEMA IF NOT EXISTS cron;

CREATE TABLE cron."job" (
    id_job              SERIAL          NOT NULL,
	"name"              TEXT            NOT NULL,
	"key"               TEXT            NOT NULL,
	running             BOOLEAN         NOT NULL DEFAULT FALSE,
	"settings"          TEXT            NOT NULL,
	"position"          INT4            NOT NULL UNIQUE,
	created_at			TIMESTAMP       NOT NULL DEFAULT NOW(),
	updated_at			TIMESTAMP       NOT NULL DEFAULT NOW(),
	"active"            BOOLEAN         NOT NULL DEFAULT TRUE,
	CONSTRAINT job_pkey PRIMARY KEY (id_job)
);



-- migrate down
DROP TABLE cron."job";
DROP SCHEMA IF EXISTS cron;