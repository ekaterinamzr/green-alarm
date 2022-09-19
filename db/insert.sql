\copy Types(type_name) FROM './data/types.csv' DELIMITER ';' CSV;
\copy Roles(role_name) FROM '/home/ekaterina/goprojects/src/green-alarm/db/data/roles.csv' DELIMITER ';' CSV;
\copy Statuses(status_name) FROM '/home/ekaterina/goprojects/src/green-alarm/db/data/statuses.csv' DELIMITER ';' CSV;
\copy Users(first_name, last_name, username, email, user_password, user_role) FROM '/home/ekaterina/goprojects/src/green-alarm/db/data/users.csv' DELIMITER ';' CSV;
\copy Incidents(incident_name, incident_date, country, latitude, longitude, publication_date, comment, incident_status, incident_type, author) FROM '/home/ekaterina/goprojects/src/green-alarm/db/data/incidents.csv' DELIMITER ';' CSV;
\copy Images(image_path, incident) FROM '/home/ekaterina/goprojects/src/green-alarm/db/data/images.csv' DELIMITER ';' CSV;
\copy Editing_Requests(old, new) FROM '/home/ekaterina/goprojects/src/green-alarm/db/data/requests.csv' DELIMITER ';' CSV;
