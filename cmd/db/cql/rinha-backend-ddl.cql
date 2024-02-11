CREATE TABLE clients (
     id int,
     account_limit int,
     account_balance int,
     PRIMARY KEY (id)
 );

CREATE TABLE transactions (
     id int,
     client_id int,
     value int,
     type text,
     description text,
     created_at timestamp,
     PRIMARY KEY (id)
 );

CREATE INDEX clientTransactionIndex ON transactions (client_id);
