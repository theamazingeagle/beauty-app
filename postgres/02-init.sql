INSERT INTO clients (name, surname) VALUES ('Нина', 'Петрова');
INSERT INTO clients (name, surname) VALUES ('Вера', 'Калинина');
INSERT INTO clients (name, surname) VALUES ('Наталья', 'Круглова');
INSERT INTO clients (name, surname) VALUES ('Марина', 'Беглова');

INSERT INTO services (name, cost) VALUES ('Массаж', 1500);
INSERT INTO services (name, cost) VALUES ('Солярий', 1000);
INSERT INTO services (name, cost) VALUES ('Сауна', 2500);

INSERT INTO orders (service_id, client_id, order_time ) VALUES (1, 1, '2021-05-23 12:00:00');
INSERT INTO orders (service_id, client_id, order_time ) VALUES (2, 2, '2021-05-23 13:00:00');
INSERT INTO orders (service_id, client_id, order_time ) VALUES (3, 3, '2021-05-23 14:00:00');
INSERT INTO orders (service_id, client_id, order_time ) VALUES (4, 1, '2021-05-23 15:00:00');
INSERT INTO orders (service_id, client_id, order_time ) VALUES (4, 2, '2021-05-23 16:00:00');
INSERT INTO orders (service_id, client_id, order_time ) VALUES (4, 3, '2021-05-23 17:00:00');