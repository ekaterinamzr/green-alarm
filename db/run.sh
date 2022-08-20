#!/bin/bash

psql -f create.sql
psql -d labs -f setconstraints.sql
psql -d labs -f insert.sql
