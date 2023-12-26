CREATE TABLE  IF NOT EXISTS timtables (
    "id"  SERIAL PRIMARY KEY,
    "value"json NOT NULL DEFAULT '{}'::json,
    );
