CREATE TYPE category AS ENUM ('caregiver', 'careseeker', 'both');
ALTER TABLE IF EXISTS users
ADD COLUMN user_category category NOT NULL DEFAULT 'careseeker',
ADD COLUMN experience VARCHAR(511) DEFAULT '',
ADD COLUMN bio VARCHAR(127) DEFAULT '',
ADD COLUMN preferences VARCHAR(511) DEFAULT '',
ADD COLUMN country VARCHAR(63) DEFAULT 'United States of America',
ADD COLUMN state VARCHAR(63) DEFAULT 'Georgia',
ADD COLUMN city VARCHAR(127) DEFAULT '',
ADD COLUMN zipcode VARCHAR(20) DEFAULT '',
ADD COLUMN address VARCHAR(255) DEFAULT '';
