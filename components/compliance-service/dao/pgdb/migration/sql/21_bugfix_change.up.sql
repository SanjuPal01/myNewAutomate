ALTER TABLE IF EXISTS nodes
  ADD COLUMN IF NOT EXISTS connection_error text NOT NULL DEFAULT '';

ALTER TABLE IF EXISTS nodes
  ADD COLUMN IF NOT EXISTS statechange_timestamp timestamp NOT NULL DEFAULT '0001-01-01T00:00:00Z00:00';