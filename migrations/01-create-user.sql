-- for storing passwords
CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE "user"(
  email TEXT PRIMARY KEY,
  pass TEXT NOT NULL
);