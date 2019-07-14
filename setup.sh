#!/bin/bash

#Initialize environment variables
source .env

#Initialize Database

if !( mysql --user=$MYSQL_USER --password=$MYSQL_PASSWORD -e 'use logistics;')
then
    echo "setting up databse"
    mysql --user=$MYSQL_USER --password=$MYSQL_PASSWORD < files/database/*.sql
else
    echo "db already setup"
fi

#Add Depenedencies
dep ensure -v