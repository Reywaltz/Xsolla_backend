CREATE TABLE IF NOT EXISTS item (
    SKU text PRIMARY KEY,
    name text NOT NULL,
    type text NOT NULL,
    cost numeric
)