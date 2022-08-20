DROP DATABASE IF EXISTS greenalarm;
CREATE DATABASE greenalarm;
\connect greenalarm

CREATE TABLE Users (
  id SERIAL PRIMARY KEY,
  name varchar NOT NULL,
  surname varchar NOT NULL,
  username varchar NOT NULL UNIQUE,
  email varchar NOT NULL UNIQUE,
  password varchar NOT NULL
  role int NOT NULL
);

CREATE TABLE Roles (
  id SERIAL PRIMARY KEY,
  name varchar NOT NULL
);

CREATE TABLE Incidents (
  id SERIAL PRIMARY KEY,
  name varchar NOT NULL,
  incident_date date,
  country varchar NOT NULL,
  latitude double precision,
  longitude double precision,
  -- coordinate geography NOT NULL,
  publication_date date NOT NULL,
  comment text,
  status int NOT NULL,
  type int NOT NULL,
  user int NOT NULL
);

CREATE TABLE Statuses (
  id SERIAL PRIMARY KEY,
  name varchar NOT NULL
);

CREATE TABLE Images (
  id SERIAL PRIMARY KEY,
  path varchar NOT NULL,
  incident int NOT NULL
);

CREATE TABLE Types (
  id SERIAL PRIMARY KEY,
  name varchar NOT NULL
);

CREATE TABLE Editing_Requests (
  id SERIAL PRIMARY KEY,
  old int NOT NULL,
  new int NOT NULL
)

ALTER TABLE Users ADD FOREIGN KEY (role) REFERENCES Roles (id);
ALTER TABLE Incidents ADD FOREIGN KEY (type) REFERENCES Types (id);
ALTER TABLE Incidents ADD FOREIGN KEY (status) REFERENCES Statuses (id);
ALTER TABLE Incidents ADD FOREIGN KEY (user) REFERENCES Users (id);
ALTER TABLE Images ADD FOREIGN KEY (incident) REFERENCES Incidents (id);
ALTER TABLE Editing_Requests ADD FOREIGN KEY (old) REFERENCES Incidents (id);
ALTER TABLE Editing_Requests ADD FOREIGN KEY (new) REFERENCES Incidents (id);

