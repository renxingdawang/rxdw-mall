CREATE TABLE tokens(
                       token_id INT AUTO_INCREMENT PRIMARY KEY ,
                       user_id INT,
                       token VARCHAR(255) NOT NULL ,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       expired_at TIMESTAMP,
                       FOREIGN KEY (user_id) REFERENCES users(user_id)
);