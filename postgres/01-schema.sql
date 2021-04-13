CREATE TABLE clients (
    id SERIAL NOT NULL PRIMARY KEY,
    name VARCHAR(256),
    surname VARCHAR(256)
);

CREATE TABLE services (
    id SERIAL NOT NULL PRIMARY KEY,
    name TEXT,
    description TEXT,
    cost BIGINT
);

CREATE TABLE orders (
    id SERIAL NOT NULL PRIMARY KEY,
    service_id BIGINT,
    client_id BIGINT,
    creation_time TIMESTAMP,
    order_time TIMESTAMP
);

ALTER TABLE orders ADD FOREIGN KEY (service_id) REFERENCES services(id);
ALTER TABLE orders ADD FOREIGN KEY (client_id) REFERENCES clients(id);
ALTER TABLE orders ALTER COLUMN creation_time SET DEFAULT Now();