CREATE TABLE Users (
  id SERIAL PRIMARY KEY,
  first_name varchar NOT NULL,
  last_name varchar NOT NULL,
  username varchar NOT NULL UNIQUE,
  email varchar NOT NULL UNIQUE,
  user_password varchar NOT NULL,
  user_role int NOT NULL
);

CREATE TABLE Roles (
  id SERIAL PRIMARY KEY,
  role_name varchar NOT NULL
);

CREATE TABLE Incidents (
  id SERIAL PRIMARY KEY,
  incident_name varchar NOT NULL,
  incident_date timestamp,
  country varchar NOT NULL,
  latitude double precision,
  longitude double precision,
  publication_date timestamp NOT NULL,
  comment text,
  incident_status int NOT NULL,
  incident_type int NOT NULL,
  author int NOT NULL
);

CREATE TABLE Statuses (
  id SERIAL PRIMARY KEY,
  status_name varchar NOT NULL
);

CREATE TABLE Types (
  id SERIAL PRIMARY KEY,
  type_name varchar NOT NULL
);

ALTER TABLE Users ADD FOREIGN KEY (user_role) REFERENCES Roles (id);
ALTER TABLE Incidents ADD FOREIGN KEY (incident_type) REFERENCES Types (id);
ALTER TABLE Incidents ADD FOREIGN KEY (incident_status) REFERENCES Statuses (id);
ALTER TABLE Incidents ADD FOREIGN KEY (author) REFERENCES Users (id);

