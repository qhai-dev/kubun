CREATE TABLE users (
  id   BIGSERIAL PRIMARY KEY,
  name text      NOT NULL,
  nick_name  text,
  email text,
  phone text,
  avatar text,
  hash_password text,
  enter_prise_list text[],
  department_list text[],
  created_at timestamptz NOT NULL DEFAULT NOW(),
  updated_at timestamptz NOT NULL DEFAULT NOW()
);