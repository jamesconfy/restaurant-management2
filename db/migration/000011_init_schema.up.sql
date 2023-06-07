CREATE TABLE IF NOT EXISTS orders (
    id uuid DEFAULT uuid_generate_v4(),
    delivery_id integer NOT NULL DEFAULT 1,
    payment_id integer NOT NULL DEFAULT 1,
    table_id uuid NOT NULL,
    date_created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    date_updated TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY(id),
    FOREIGN KEY(delivery_id) REFERENCES delivery(id) ON DELETE CASCADE,
    FOREIGN KEY(table_id) REFERENCES tables(id) ON DELETE CASCADE,
    FOREIGN KEY(payment_id) REFERENCES payments(id) ON DELETE CASCADE
);