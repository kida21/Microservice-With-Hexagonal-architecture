CREATE EXTENSION IF NOT EXISTS citext;
CREATE TABLE IF NOT EXISTS users(
  id bigserial PRIMARY KEY NOT NULL,
  firstname text NOT NULL,
  lastname text NOT NULL,
  email citext UNIQUE NOT NULL,
  activated boolean NOT NULL DEFAULT false,
  password_hash bytea NOT NULL,
  created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
  version integer NOT NULL DEFAULT 1
  );