CREATE ROLE "seller" WITH
    LOGIN
    NOSUPERUSER
    NOCREATEDB
    NOCREATEROLE
    NOREPLICATION
    PASSWORD 'seller'
    CONNECTION LIMIT -1;

GRANT SELECT ON racket, 
                supplier,
                "order",
                order_racket
                TO "seller";
GRANT UPDATE ON racket,
                "order",
                order_racket
                TO "seller";
GRANT INSERT ON racket TO "seller";