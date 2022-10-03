#!/bin/bash

mongoimport --type csv -d greenalarm -c incidents --headerline --drop /data/scvs/incidents.csv