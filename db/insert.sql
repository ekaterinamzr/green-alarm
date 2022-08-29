\copy Types FROM './data/types.csv' DELIMITER ';' CSV;
\copy Roles FROM '/home/ekaterina/goprojects/src/green-alarm/db/data/roles.csv' DELIMITER ';' CSV;
\copy Statuses FROM '/home/ekaterina/goprojects/src/green-alarm/db/data/statuses.csv' DELIMITER ';' CSV;
\copy Users(first_name, last_name, username, email, user_password, user_role) FROM '/home/ekaterina/goprojects/src/green-alarm/db/data/users.csv' DELIMITER ';' CSV;
\copy Incidents FROM '/home/ekaterina/goprojects/src/green-alarm/db/data/incidents.csv' DELIMITER ';' CSV;
\copy Images FROM '/home/ekaterina/goprojects/src/green-alarm/db/data/images.csv' DELIMITER ';' CSV;
\copy Editing_Requests FROM '/home/ekaterina/goprojects/src/green-alarm/db/data/requests.csv' DELIMITER ';' CSV;
