SET TIMEZONE='US/Eastern';
CREATE TYPE care_category AS ENUM ('tutoring', 'baby-sitting', 'other');
CREATE TABLE IF NOT EXISTS posts(user_id uuid NOT NULL, caregiver_id uuid DEFAULT NULL, post_id INT GENERATED ALWAYS AS IDENTITY, title VARCHAR(50) NOT NULL DEFAULT '', tags VARCHAR(255) NOT NULL DEFAULT '', care_description VARCHAR(1023) NOT NULL DEFAULT '', care_type care_category NOT NULL, completed boolean NOT NULL DEFAULT FALSE, start_date DATE NOT NULL, start_time TIME with time zone NOT NULL, end_date DATE NOT NULL, end_time TIME with time zone NOT NULL, date_posted TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP, PRIMARY KEY (post_id), CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(user_id));
