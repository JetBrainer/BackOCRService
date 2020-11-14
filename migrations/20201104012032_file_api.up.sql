CREATE TABLE account(
    id SERIAL PRIMARY KEY,
    login VARCHAR UNIQUE,
    password VARCHAR,
    organization VARCHAR,
    token VARCHAR
)