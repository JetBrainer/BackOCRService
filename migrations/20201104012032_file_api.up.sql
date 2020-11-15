CREATE TABLE acc(
    id              SERIAL PRIMARY KEY,
    email           VARCHAR UNIQUE,
    encpassword     VARCHAR,
    organization    VARCHAR,
    token           VARCHAR
)