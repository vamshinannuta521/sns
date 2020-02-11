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
   account_name varchar(200) REFERENCES Account(name)
);

CREATE TABLE if not exists Action_Type(
   id SERIAL primary key,
   type varchar(50) unique not null
);

CREATE TABLE if not exists Action (
   id SERIAL PRIMARY KEY,
   event_name varchar(50) REFERENCES event(name),
   action_type varchar(50) references Action_type(type),
   action_spec varchar(5000),
   account_name varchar(200) REFERENCES Account(name)
);

CREATE TABLE if not exists Trigger(
	uuid  UUID PRIMARY KEY,
	event_name varchar(50) REFERENCES event(name),
	message varchar(5000) default null,
	account_name varchar(200) REFERENCES Account(name)
);

insert into action_type (type) values ('http'), ('https'), ('smtp');

END_OF_SCRIPT