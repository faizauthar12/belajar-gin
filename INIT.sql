CREATE DATABASE IF NOT EXIST "db-book-sql"
    WITH
    OWNER = faizauthar
    ENCODING = 'UTF8'
    LC_COLLATE = 'en_US.UTF-8'
    LC_CTYPE = 'en_US.UTF-8'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1
    IS_TEMPLATE = False;

CREATE TABLE IF NOT EXIST books(
    book_id SERIAL PRIMARY KEY,
    title varchar(100) NOT NULL,
    author varchar(100) NOT NULL,
    description varchar(100) NOT NULL
)