CREATE TYPE payment_type AS ENUM ('CASH', 'CARD');

CREATE TABLE IF NOT EXISTS payment_method (
    id SERIAL,
    payment_type payment_type DEFAULT 'CASH',
    date_created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    date_updated TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY(id)
);