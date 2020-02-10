#!/bin/bash
 
###################################################
# Bash Shell script to create psql tables
###################################################



psql -h localhost -U postgres -p ${PGPORT:=5432} --echo-all << END_OF_SCRIPT

drop database if exists sns;
create database sns;

\c sns


CREATE TABLE if not exists ACCOUNT  (
   id SERIAL PRIMARY KEY,
   name varchar(200) NOT NULL UNIQUE
);

CREATE TABLE if not exists EVENT(
   id SERIAL PRIMARY KEY,
   name varchar(50) NOT NULL UNIQUE,
   account_id integer REFERENCES Account(id) NOT NULL
);

CREATE TABLE if not exists Action (
   id SERIAL PRIMARY KEY,
   event_id integer REFERENCES event(id) NOT NULL,
   action_type varchar(50) NOT NULL,
   action_spec varchar(5000) not null,
   account_id integer REFERENCES account (id) not null
);

CREATE TABLE if not exists Trigger(
	id SERIAL PRIMARY KEY,
	event_id  integer REFERENCES Event(id) NOT NULL,
	message varchar(5000) default null,
	account_id integer REFERENCES Account(id) NOT NULL
);

END_OF_SCRIPT