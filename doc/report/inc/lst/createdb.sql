DROP DATABASE IF EXISTS greenalarm;
DROP ROLE IF EXISTS greenalarm;

CREATE DATABASE greenalarm;
CREATE USER greenalarm WITH 
    CREATEROLE
    PASSWORD 'greenalarm';
GRANT ALL PRIVILEGES ON DATABASE greenalarm TO greenalarm;
