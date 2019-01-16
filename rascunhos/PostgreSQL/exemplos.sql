CREATE TABLE users
(
    id serial NOT NULL,
    name character varying(240) NOT NULL,
    password character varying(240) NOT NULL,
    full_name character varying(240),
    salary numeric NOT NULL,
    birthday  time without time zone,
    create_date time without time zone NOT NULL DEFAULT Now(),
    PRIMARY KEY (id)
);
