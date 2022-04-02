SET TIMEZONE='US/Eastern';
CREATE TABLE IF NOT EXISTS reviews(
    user_id uuid NOT NULL, 
    review_id INT GENERATED ALWAYS AS IDENTITY, 
    reviewer_id uuid NOT NULL,
    post_id INT NOT Null,
    rating INT NOT NULL DEFAULT 0, 
    comment VARCHAR(2047) NOT NULL DEFAULT '',
    review_date TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    show BOOLEAN NOT NULL DEFAULT TRUE,
    PRIMARY KEY (review_id),  
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(user_id),
    CONSTRAINT fk_reviewer FOREIGN KEY (reviewer_id) REFERENCES users(user_id),
    CONSTRAINT fk_post FOREIGN KEY (post_id) REFERENCES posts(post_id)
);