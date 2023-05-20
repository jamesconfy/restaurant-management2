CREATE TABLE IF NOT EXISTS order_item (
    id uuid DEFAULT uuid_generate_v4(),
    quantity INTEGER NOT NULL DEFAULT 1,
    order_id uuid NOT NULL,
    food_id uuid NOT NULL,
    date_created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    date_updated TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY(id),
    FOREIGN KEY(order_id) REFERENCES order(id) ON CASCADE DELETE,
    FOREIGN KEY(food_id) REFERENCES food(id) ON CASCADE DELETE
);