COPY Types(type_name) FROM '/var/lib/postgresql/csvs/types.csv' WITH(FORMAT csv, HEADER);
COPY Roles(role_name) FROM '/var/lib/postgresql/csvs/roles.csv' WITH(FORMAT csv, HEADER);
COPY Statuses(status_name) FROM '/var/lib/postgresql/csvs/statuses.csv' WITH(FORMAT csv, HEADER);
COPY Users(first_name, last_name, username, email, user_password, user_role) FROM '/var/lib/postgresql/csvs/users.csv' WITH(FORMAT csv, HEADER);
COPY Incidents(incident_name, incident_date, country, latitude, longitude, publication_date, comment, incident_status, incident_type, author) FROM '/var/lib/postgresql/csvs/incidents.csv' WITH(FORMAT csv, HEADER);