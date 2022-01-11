CREATE TYPE category AS ENUM ('caregiver', 'careseeker', 'both');
ALTER TABLE IF EXISTS users
ADD COLUMN user_category category NOT NULL DEFAULT 'careseeker',
ADD COLUMN experience VARCHAR(511) DEFAULT '',
ADD COLUMN bio VARCHAR(127) DEFAULT '';