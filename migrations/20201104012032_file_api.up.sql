CREATE TABLE acc(
    id              SERIAL PRIMARY KEY,
    email           VARCHAR UNIQUE,
    password        VARCHAR,
    organization    VARCHAR,
    token           VARCHAR
)