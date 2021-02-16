CREATE TABLE referral(
  from_domain TEXT NOT NULL,
  from_path TEXT NOT NULL,
  to_domain TEXT NOT NULL,
  to_path TEXT NOT NULL,
  count INT NOT NULL DEFAULT 0,

  FOREIGN KEY (to_domain, to_path) REFERENCES page(domain, path),
  PRIMARY KEY (from_domain, from_path, to_domain, to_path)
);