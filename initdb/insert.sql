COPY Types FROM '/var/lib/postgresql/csvs/types.csv' DELIMITER ';' CSV;
COPY Roles FROM '/var/lib/postgresql/csvs/roles.csv' DELIMITER ';' CSV;
COPY Statuses FROM '/var/lib/postgresql/csvs/statuses.csv' DELIMITER ';' CSV;
COPY Users FROM '/var/lib/postgresql/csvs/users.csv' DELIMITER ';' CSV;
COPY Incidents FROM '/var/lib/postgresql/csvs/incidents.csv' DELIMITER ';' CSV;
COPY Images FROM '/var/lib/postgresql/csvs/images.csv' DELIMITER ';' CSV;
COPY Editing_Requests FROM '/var/lib/postgresql/csvs/requests.csv' DELIMITER ';' CSV;
