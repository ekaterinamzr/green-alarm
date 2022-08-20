COPY Incidents FROM '/home/ekaterina/goprojects/src/green-alarm/db/data/incidents.csv' DELIMITER ';' CSV;
COPY Users FROM '/home/ekaterina/goprojects/src/green-alarm/db/data/users.csv' DELIMITER ';' CSV;
COPY Roles FROM '/home/ekaterina/goprojects/src/green-alarm/db/data/roles.csv' DELIMITER ';' CSV;
COPY Statuses FROM '/home/ekaterina/goprojects/src/green-alarm/db/data/statuses.csv' DELIMITER ';' CSV;
COPY Images FROM '/home/ekaterina/goprojects/src/green-alarm/db/data/images.csv' DELIMITER ';' CSV;
COPY Types FROM '/home/ekaterina/goprojects/src/green-alarm/db/data/types.csv' DELIMITER ';' CSV;
COPY Editing_Requests FROM '/home/ekaterina/goprojects/src/green-alarm/db/data/requests.csv' DELIMITER ';' CSV;