#!/bin/bash

psql -f createdb.sql
sudo -u greenalarm psql -f createtables.sql
sudo -u greenalarm psql -f createroles.sql
sudo -u greenalarm psql -f setconstraints.sql
sudo -u greenalarm psql -f createtrigger.sql
sudo -u greenalarm psql -f insert.sql


