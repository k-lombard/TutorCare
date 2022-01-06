CREATE TYPE category AS ENUM ('caregiver', 'careseeker', 'both');
ALTER TABLE IF EXISTS users
ADD COLUMN user_category category,
ADD COLUMN experience VARCHAR(511),
ADD COLUMN bio VARCHAR(127);