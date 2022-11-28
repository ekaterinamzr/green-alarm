\copy Types(type_name) FROM './data/types.csv' DELIMITER ';' CSV;
\copy Roles(role_name) FROM './data/roles.csv' DELIMITER ';' CSV;
\copy Statuses(status_name) FROM './data/statuses.csv' DELIMITER ';' CSV;
\copy Users(first_name, last_name, username, email, user_password, user_role) FROM './data/users.csv' DELIMITER ';' CSV;
\copy Incidents(incident_name, incident_date, country, latitude, longitude, publication_date, comment, incident_status, incident_type, author) FROM './data/incidents.csv' DELIMITER ';' CSV;