CREATE TABLE IF NOT EXISTS food (
    id uuid DEFAULT uuid_generate_v4(),
    name VARCHAR(250) NOT NULL DEFAULT '',
    price FLOAT NOT NULL DEFAULT 0.0,
    image VARCHAR(250),
    menu_id uuid NOT NULL,
    date_created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    date_updated TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY(id),
    FOREIGN KEY(menu_id) REFERENCES "menu"(id) ON DELETE CASCADE,
    CHECK(price > 0.0)
);