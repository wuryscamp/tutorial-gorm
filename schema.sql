CREATE TABLE products (
    id character varying(10) UNIQUE PRIMARY KEY NOT NULL, 
    name character varying(50) NOT NULL, 
    quantity integer NOT NULL NOT NULL, 
    created_at timestamp with time zone DEFAULT now() NOT NULL, 
    updated_at timestamp with time zone, 
    version integer DEFAULT 1 NOT NULL
);