SET TIMEZONE='US/Eastern';
CREATE TABLE IF NOT EXISTS geolocation(user_id uuid NOT NULL, location_id INT GENERATED ALWAYS AS IDENTITY, accuracy float(32) NOT NULL, latitude float(32) NOT NULL, longitude float(32) NOT NULL, timestamp TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP, PRIMARY KEY (location_id), CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(user_id));
