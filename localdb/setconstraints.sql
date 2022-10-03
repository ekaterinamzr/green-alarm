ALTER TABLE Incidents
    ADD CONSTRAINT latitude_check CHECK (latitude >= -90 AND latitude <= 90),
    ADD CONSTRAINT longitude_check CHECK (longitude >= -180 AND longitude < 180)