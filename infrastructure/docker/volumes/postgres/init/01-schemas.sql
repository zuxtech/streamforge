-- Schemas
CREATE SCHEMA IF NOT EXISTS streamforge;
CREATE SCHEMA IF NOT EXISTS kratos;


-- Runtime application role
GRANT USAGE
ON SCHEMA streamforge
TO streamforge_app;


-- Existing tables
GRANT SELECT, INSERT, UPDATE, DELETE
ON ALL TABLES IN SCHEMA streamforge
TO streamforge_app;


-- Future tables created by streamforge_migrator
ALTER DEFAULT PRIVILEGES
FOR ROLE streamforge_migrator
IN SCHEMA streamforge
GRANT SELECT, INSERT, UPDATE, DELETE
ON TABLES
TO streamforge_app;


-- Existing sequences
GRANT USAGE, SELECT
ON ALL SEQUENCES IN SCHEMA streamforge
TO streamforge_app;


-- Future sequences created by streamforge_migrator
ALTER DEFAULT PRIVILEGES
FOR ROLE streamforge_migrator
IN SCHEMA streamforge
GRANT USAGE, SELECT
ON SEQUENCES
TO streamforge_app;