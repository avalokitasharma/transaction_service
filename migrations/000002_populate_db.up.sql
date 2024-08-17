-- Basic transactions without parent
INSERT INTO transactions (id, amount, type) VALUES
(1, 1000, 'cars'),
(2, 2000, 'shopping'),
(3, 3000, 'food');

-- Transactions with parents
INSERT INTO transactions (id, amount, type, parent_id) VALUES
(4, 500, 'cars', 1),
(5, 750, 'cars', 1),
(6, 1500, 'shopping', 2),
(7, 250, 'food', 3);

-- Multi-level linked transactions
INSERT INTO transactions (id, amount, type, parent_id) VALUES
(8, 100, 'cars', 4),
(9, 200, 'cars', 4),
(10, 300, 'shopping', 6);

-- Transactions of the same type
INSERT INTO transactions (id, amount, type) VALUES
(11, 5000, 'electronics'),
(12, 7500, 'electronics'),
(13, 10000, 'electronics');

-- Transaction with a large sum of child transactions
INSERT INTO transactions (id, amount, type) VALUES
(14, 10000, 'investment');

INSERT INTO transactions (id, amount, type, parent_id) VALUES
(15, 1000, 'stocks', 14),
(16, 2000, 'bonds', 14),
(17, 3000, 'real_estate', 14);

INSERT INTO transactions (id, amount, type, parent_id) VALUES
(18, 500, 'stocks', 15),
(19, 1500, 'stocks', 15),
(20, 1000, 'bonds', 16);