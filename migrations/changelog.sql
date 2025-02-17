--changeset bonefabric:create_users_table
CREATE TABLE IF NOT EXISTS users
(
    id         SERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT (NOW()),
    updated_at TIMESTAMP NOT NULL DEFAULT (NOW())
        CONSTRAINT updated_at_valid CHECK ( updated_at >= created_at )
)