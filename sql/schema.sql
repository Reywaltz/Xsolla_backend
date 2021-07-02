CREATE TABLE IF NOT EXISTS item (
    SKU text UNIQUE PRIMARY KEY,
    name text NOT NULL,
    type text NOT NULL,
    cost numeric
);

INSERT INTO item (SKU, name, type, cost) VALUES ('DOT-SUB', 'Dota Plus', 'Subscription', '9.99');
INSERT INTO item (SKU, name, type, cost) VALUES ('SON-GAM', 'Sonic rangers', 'Game', '59.99');
INSERT INTO item (SKU, name, type, cost) VALUES ('DMC-GAM', 'DMC:Devil may cry', 'Game', '69.99');
INSERT INTO item (SKU, name, type, cost) VALUES ('OVE-GAM', 'Overwatch', 'Game', '39.99');
INSERT INTO item (SKU, name, type, cost) VALUES ('HOT-GAM', 'Hotline Miami', 'Game', '49.99');
INSERT INTO item (SKU, name, type, cost) VALUES ('XBO-SUB', 'Xbox Gamepass', 'Subsсription', '5.99');
INSERT INTO item (SKU, name, type, cost) VALUES ('PLA-SUB', 'Playstation Plus', 'Subsсription', '6.99');