CREATE DATABASE brimoveapp;

\c brimoveapp;

CREATE TABLE "accounts" (
                            "id" INTEGER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
                            "username" varchar(256) NOT NULL,
                            "password" varchar(256) NOT NULL
);