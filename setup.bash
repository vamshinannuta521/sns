#!/bin/bash
 
###################################################
# Bash Shell script to create psql tables
###################################################
 
sudo su - postgres
#Set the value of db name
database="sns"
account = "CREATE TABLE ACCOUNT(
   ID INT PRIMARY KEY     NOT NULL,
   EventAction_list TEXT    NULL,
   NAME TEXT NOT NULL,
)"

subscription = "CREATE TABLE SUBSCRIBED_EVENT_ACTION(
   ID INT PRIMARY KEY     NOT NULL,
   Event_Id INT REFERENCES Event(id) NOT NULL,
   Action_type INT NOT NULL,
   Action_spec TEXT not null
)"

event = "CREATE TABLE EVENT(
   ID INT PRIMARY KEY NOT NULL,
   Name TEXT NOT NULL,
   CreatedBy INT REFERENCES Account(id) NOT NULL
)"

createdb $database
psql -d $database -c $account
psql -d $database -c $event
psql -d $database -c $subscription