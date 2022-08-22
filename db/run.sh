#!/bin/bash

psql -f createdb.sql
sudo -u greenalarm psql -f createtables.sql
sudo -u greenalarm psql -f setconstraints.sql
sudo -u greenalarm psql -f insert.sql
# PGPASSWORD=greenalarm psql -U greenalarm -f createtables.sql
# PGPASSWORD=greenalarm psql -U greenalarm -f setconstraints.sql
# PGPASSWORD=greenalarm psql -U greenalarm -f insert.sql
