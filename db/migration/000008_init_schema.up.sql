CREATE TABLE IF NOT EXISTS tables (
    id uuid DEFAULT uuid_generate_v4(),
    seats INTEGER NOT NULL DEFAULT 1,
    number SERIAL,
    date_created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    date_updated TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY(id),
    CHECK (seats > 0)
);