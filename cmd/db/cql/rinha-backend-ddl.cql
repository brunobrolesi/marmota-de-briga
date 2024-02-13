CREATE TABLE clients (
     id int,
     account_limit int,
     account_balance int,
     PRIMARY KEY (id)
 );

CREATE TABLE transactions (
     client_id int,
     value int,
     type text,
     description text,
     created_at timestamp,
    PRIMARY KEY (client_id, created_at))
    WITH CLUSTERING ORDER BY (created_at DESC);

