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
  incident_date date,
  country varchar NOT NULL,
  latitude double precision,
  longitude double precision,
  -- coordinate geography NOT NULL,
  publication_date date NOT NULL,
  comment text,
  incident_status int NOT NULL,
  incident_type int NOT NULL,
  author int NOT NULL
);

CREATE TABLE Statuses (
  id SERIAL PRIMARY KEY,
  status_name varchar NOT NULL
);

CREATE TABLE Images (
  id SERIAL PRIMARY KEY,
  image_path varchar NOT NULL,
  incident int NOT NULL
);

CREATE TABLE Types (
  id SERIAL PRIMARY KEY,
  type_name varchar NOT NULL
);

CREATE TABLE Editing_Requests (
  id SERIAL PRIMARY KEY,
  old int NOT NULL,
  new int NOT NULL
);

ALTER TABLE Users ADD FOREIGN KEY (user_role) REFERENCES Roles (id);
ALTER TABLE Incidents ADD FOREIGN KEY (incident_type) REFERENCES Types (id);
ALTER TABLE Incidents ADD FOREIGN KEY (incident_status) REFERENCES Statuses (id);
ALTER TABLE Incidents ADD FOREIGN KEY (author) REFERENCES Users (id);
ALTER TABLE Images ADD FOREIGN KEY (incident) REFERENCES Incidents (id);
ALTER TABLE Editing_Requests ADD FOREIGN KEY (old) REFERENCES Incidents (id);
ALTER TABLE Editing_Requests ADD FOREIGN KEY (new) REFERENCES Incidents (id);

