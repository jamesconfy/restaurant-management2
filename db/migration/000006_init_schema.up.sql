CREATE TABLE IF NOT EXISTS menu (
    id uuid DEFAULT uuid_generate_v4(),
    name VARCHAR(250) NOT NULL DEFAULT '',
    category VARCHAR(250) NOT NULL DEFAULT '',
    date_created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    date_updated TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY(id)
);