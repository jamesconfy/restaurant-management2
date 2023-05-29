CREATE TYPE delivery_status AS ENUM ('DELIVERED', 'ONGOING', 'NOT_DELIVERED');    

CREATE TABLE IF NOT EXISTS delivery (
    id uuid DEFAULT uuid_generate_v4(),
    status delivery_status DEFAULT 'ONGOING',
    date_created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    date_updated TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY(id)
);