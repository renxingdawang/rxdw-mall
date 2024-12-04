CREATE DATABASE IF NOT EXISTS 'rxdwMall'
    DEFAULT CHARACTER SET ='utf8mb4';

-- user table
CREATE TABLE users (
                       user_id INT AUTO_INCREMENT PRIMARY KEY,
                       email VARCHAR(255) NOT NULL UNIQUE,
                       password VARCHAR(255) NOT NULL,
    -- 可以添加其他用户相关字段，如用户名、注册时间等
                       INDEX(email)
);

-- products table
CREATE TABLE products(
    product_id INT AUTO_INCREMENT PRIMARY KEY ,
    name VARBINARY(255) NOT NULL ,
    description TEXT,
    picture VARCHAR(255),
    price DECIMAL(10,2) NOT NULL ,
    INDEX(name)
);

-- products table --商品类别 按类别分 return list 关联products table
CREATE TABLE product_categories(
    product_id INT,
    category_name VARCHAR(255),
    FOREIGN KEY (product_id) REFERENCES products(product_id)
);

-- carts table  &   carts_items table
CREATE TABLE carts(
    cart_id INT AUTO_INCREMENT PRIMARY KEY ,
    user_id INT NOT NULL ,
    FOREIGN KEY (user_id) REFERENCES users(user_id)
);

CREATE TABLE cart_items(
    cart_item_id INT AUTO_INCREMENT PRIMARY KEY ,
    cart_id INT NOT NULL ,
    product_id INT NOT NULL ,
    quantity INT NOT NULL ,
    FOREIGN KEY (cart_id) REFERENCES carts(cart_id),
    FOREIGN KEY (product_id) REFERENCES products(product_id)
);

-- orders order_items addresses
CREATE TABLE address(
    address_id INT AUTO_INCREMENT PRIMARY KEY ,
    street_address VARCHAR(255) NOT NULL,
    city VARCHAR(255) NOT NULL ,
    state VARCHAR(255) NOT NULL ,
    country VARCHAR(255) NOT NULL ,
    zip_code INT NOT NULL
);

CREATE TABLE orders(
    order_id INT AUTO_INCREMENT KEY,
    user_id INT NOT NULL ,
    user_currency VARCHAR(10) NOT NULL ,
    address_id INT NOT NULL ,
    email VARCHAR(255) NOT NULL ,
    create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(user_id),
    FOREIGN KEY (address_id) REFERENCES address(address_id)
);

CREATE TABLE order_items(
    order_item_id INT AUTO_INCREMENT PRIMARY KEY ,
    order_id INT NOT NULL ,
    product_id INT NOT NULL ,
    quantity INT NOT NULL,
    cost DECIMAL(10,2) NOT NULL ,
    FOREIGN KEY (order_id) REFERENCES orders(order_id),
    FOREIGN KEY (product_id) REFERENCES products(product_id)
);

-- payments table
CREATE TABLE payments(
    payment_id INT AUTO_INCREMENT PRIMARY KEY ,
    amount DECIMAL(10,2) NOT NULL ,
    credit_card_number VARCHAR(255) NOT NULL,
    credit_card_cvv INT NOT NULL ,
    credit_card_expiration_year INT NOT NULL ,
    credit_card_expiration_month INT NOT NULL ,
    order_id INT NOT NULL ,
    user_id INT NOT NULL ,
    transaction_id VARCHAR(255) NOT NULL ,
    FOREIGN KEY (order_id) REFERENCES orders(order_id),
    FOREIGN KEY (user_id) REFERENCES users(user_id)
);

-- tokens table
CREATE TABLE tokens(
    token_id INT AUTO_INCREMENT PRIMARY KEY ,
    user_id INT,
    token VARCHAR(255) NOT NULL ,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    expired_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(user_id)
);