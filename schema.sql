CREATE TABLE IF NOT EXISTS item (
    SKU text UNIQUE PRIMARY KEY,
    name text NOT NULL,
    type text NOT NULL,
    cost numeric
)