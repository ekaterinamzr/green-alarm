#!/bin/bash

mongoimport --type csv -d greenalarm -c incidents --headerline --drop /data/scvs/incidents.csv
mongoimport --type csv -d greenalarm -c users --headerline --drop /data/scvs/users.csv
mongoimport --type csv -d greenalarm -c counters --headerline --drop /data/scvs/mgcounters.csv