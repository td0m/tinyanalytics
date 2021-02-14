CREATE TABLE site(
  domain TEXT PRIMARY KEY,
  owner TEXT,

  FOREIGN KEY (owner) REFERENCES "user"(email)
);
