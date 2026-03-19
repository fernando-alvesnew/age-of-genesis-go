CREATE TABLE IF NOT EXISTS users (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(20) NOT NULL,
    login VARCHAR(15) NOT NULL UNIQUE,
    email VARCHAR(40) NOT NULL UNIQUE,
    user_type ENUM('player', 'tutor', 'gm', 'admin') NOT NULL DEFAULT 'player',
    password VARCHAR(255) NOT NULL,
    last_ip VARCHAR(45) NULL,
    is_banned TINYINT(1) NOT NULL DEFAULT 0,
    created_at DATETIME NULL,
    updated_at DATETIME NULL
);

CREATE TABLE IF NOT EXISTS pagseguro_credit_card (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    store_carts_id BIGINT UNSIGNED NOT NULL,
    users_id BIGINT UNSIGNED NOT NULL,
    payment_id VARCHAR(255) NOT NULL,
    reference_id VARCHAR(255) NOT NULL UNIQUE,
    amount BIGINT NOT NULL,
    status VARCHAR(50) NOT NULL,
    description TEXT NULL,
    created_at DATETIME NULL,
    updated_at DATETIME NULL,
    deleted_at DATETIME NULL,
    KEY idx_pgc_users_id (users_id),
    KEY idx_pgc_store_carts_id (store_carts_id),
    CONSTRAINT fk_pgc_user FOREIGN KEY (users_id) REFERENCES users(id)
);
