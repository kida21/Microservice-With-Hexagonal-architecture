CREATE TABLE IF NOT EXISTS users(
  id bigserial PRIMARY KEY NOT NULL,
  firstname text NOT NULL,
  lastname text NOT NULL,
  email citext UNIQUE NOT NULL,
  activated boolean NOT NULL,
  password bytea NOT NULL,
  created_at timestamp with time zone(0) NOT NULL,
  version integer NOT NULL DEFAULT 1
  );