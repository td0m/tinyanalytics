CREATE TABLE page(
  domain TEXT NOT NULL,
  path TEXT NOT NULL,

  PRIMARY KEY (domain, path),
  FOREIGN KEY (domain) REFERENCES site(domain)
);