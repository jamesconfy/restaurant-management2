CREATE TYPE payment_status AS ENUM ('PAID', 'NOT PAID');

CREATE TABLE IF NOT EXISTS payment (
    id uuid DEFAULT uuid_generate_v4(),
    status payment_status,
    date_created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    date_updated TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY(id)
);