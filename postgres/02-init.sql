INSERT INTO clients (name) VALUES ('Нина');
INSERT INTO clients (name) VALUES ('Вера');
INSERT INTO clients (name) VALUES ('Наталья');
INSERT INTO clients (name) VALUES ('Марина');

INSERT INTO services (name, cost) VALUES ('Массаж', 1500);
INSERT INTO services (name, cost) VALUES ('Солярий', 1000);
INSERT INTO services (name, cost) VALUES ('Сауна', 2500);

INSERT INTO orders ( service_id, client_id, order_time ) VALUES (1, 1, '2021-05-23 12:00:00');
INSERT INTO orders ( service_id, client_id, order_time ) VALUES (2, 2, '2021-05-23 13:00:00');
INSERT INTO orders ( service_id, client_id, order_time ) VALUES (3, 3, '2021-05-23 14:00:00');
INSERT INTO orders ( service_id, client_id, order_time ) VALUES (1, 4, '2021-05-23 15:00:00');
INSERT INTO orders ( service_id, client_id, order_time ) VALUES (2, 4, '2021-05-23 16:00:00');
INSERT INTO orders ( service_id, client_id, order_time ) VALUES (3, 4, '2021-05-23 17:00:00');