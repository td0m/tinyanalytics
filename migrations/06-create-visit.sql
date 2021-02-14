-- for time series data
CREATE EXTENSION IF NOT EXISTS timescaledb;
-- for location
CREATE EXTENSION IF NOT EXISTS postgis;

CREATE TABLE visit(
  time TIMESTAMPTZ NOT NULL,
  ip INET NOT NULL,
  domain TEXT NOT NULL,
  path TEXT NOT NULL,

  geo GEOGRAPHY,
  browser browser NOT NULL,
  platform platform NOT NULL,

  FOREIGN KEY (domain, path) REFERENCES page(domain, path)
);

-- create time series hyper table
SELECT create_hypertable('visit', 'time');

-- create geography index for faster geography-based queries
CREATE INDEX ON visit USING gist (geo);