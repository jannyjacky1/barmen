#!/usr/bin/env bash

set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" <<-EOSQL
    CREATE DATABASE $DB_NAME;

    \c $DB_NAME;
    CREATE TABLE IF NOT EXISTS tbl_files (
      id SERIAL PRIMARY KEY,
      filepath varchar(255) not null
    );

    CREATE TABLE IF NOT EXISTS tbl_complication_levels (
      id SERIAL PRIMARY KEY,
      name varchar(60) not null,
      time varchar(60) not null
    );

    CREATE TABLE IF NOT EXISTS tbl_fortress_levels (
      id SERIAL PRIMARY KEY,
      name varchar(60) not null,
      fortress_from integer not null,
      fortress_to integer not null
    );

    CREATE TABLE IF NOT EXISTS tbl_areas (
      id SERIAL PRIMARY KEY,
      name varchar(60) not null,
      parent_id integer REFERENCES tbl_areas(id)
    );

    CREATE TABLE IF NOT EXISTS tbl_cocktails (
      id SERIAL PRIMARY KEY,
      name varchar(60) not null,
      description text not null,
      complication_id integer REFERENCES tbl_complication_levels(id),
      fortress integer,
      area_id integer REFERENCES tbl_areas(id),
      recipe text not null,
      mark float,
      preview_id integer REFERENCES tbl_files(id)
    );

    CREATE TABLE IF NOT EXISTS tbl_ingredients (
      id SERIAL PRIMARY KEY,
      name varchar(60) not null,
      description text,
      img_id integer REFERENCES tbl_files(id)
    );

    CREATE TABLE IF NOT EXISTS tbl_instruments (
      id SERIAL PRIMARY KEY,
      name varchar(60) not null,
      description text,
      img_id integer REFERENCES tbl_files(id)
    );

    CREATE TABLE IF NOT EXISTS tbl_cocktails_to_tbl_files (
      img_id integer not null REFERENCES tbl_files(id),
      cocktail_id integer not null REFERENCES tbl_cocktails(id),
      PRIMARY KEY (img_id, cocktail_id)
    );

    CREATE TABLE IF NOT EXISTS tbl_cocktails_to_tbl_ingredients (
      ingredient_id integer not null REFERENCES tbl_ingredients(id),
      cocktail_id integer not null REFERENCES tbl_cocktails(id),
      volume integer not null,
      PRIMARY KEY (ingredient_id, cocktail_id)
    );

    CREATE TABLE IF NOT EXISTS tbl_cocktails_to_tbl_instruments (
      instrument_id integer not null REFERENCES tbl_instruments(id),
      cocktail_id integer not null REFERENCES tbl_cocktails(id),
      PRIMARY KEY (instrument_id, cocktail_id)
    );

    CREATE TABLE IF NOT EXISTS tbl_users (
      id SERIAL PRIMARY KEY,
      device_id varchar(255) not null unique
    );

    CREATE TABLE IF NOT EXISTS tbl_favorites (
      user_id integer not null REFERENCES tbl_users(id),
      cocktail_id integer not null REFERENCES tbl_cocktails(id),
      PRIMARY KEY (user_id, cocktail_id)
    );

    CREATE TABLE IF NOT EXISTS tbl_tries (
      user_id integer not null REFERENCES tbl_users(id),
      cocktail_id integer not null REFERENCES tbl_cocktails(id),
      PRIMARY KEY (user_id, cocktail_id)
    );

    CREATE TABLE IF NOT EXISTS tbl_admins (
      id SERIAL PRIMARY KEY,
      email varchar not null unique,
      username varchar not null,
      password_hash varchar not null,
      token varchar,
      token_expire timestamp
    );
EOSQL
