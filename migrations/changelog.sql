--changeset bonefabric:create_users_table
CREATE TABLE IF NOT EXISTS users
(
    id         SERIAL PRIMARY KEY,
    telegram_user_id INT UNIQUE NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT (NOW()),
    updated_at TIMESTAMP NOT NULL DEFAULT (NOW())
        CONSTRAINT updated_at_valid CHECK ( updated_at >= created_at )
);
--rollback DROP TABLE IF EXISTS users;
