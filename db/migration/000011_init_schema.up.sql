CREATE TABLE IF NOT EXISTS order (
    id uuid DEFAULT uuid_generate_v4(),
    delivery_id uuid,
    payment_id uuid,
    table_id uuid,
    date_created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    date_updated TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY(id),
    FOREIGN KEY(delivery_id) REFERENCES delivery(id) ON CASCADE SET NULL,
    FOREIGN KEY(table_id) REFERENCES table(id) ON CASCADE SET NULL,
    FOREIGN KEY(payment_id) REFERENCES payment(id) ON CASCADE SET NULL
);