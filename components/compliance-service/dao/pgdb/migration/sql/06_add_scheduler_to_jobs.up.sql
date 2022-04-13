ALTER TABLE IF EXISTS jobs
  ADD COLUMN IF NOT EXISTS scheduled_time timestamp DEFAULT '0001-01-01T00:00:00Z00:00',
  ADD COLUMN IF NOT EXISTS recurrence TEXT NOT NULL DEFAULT '',
  ADD COLUMN IF NOT EXISTS parent_id TEXT NOT NULL DEFAULT '',
  ADD COLUMN IF NOT EXISTS job_count int NOT NULL DEFAULT 0;