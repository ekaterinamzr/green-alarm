DROP ROLE IF EXISTS admin, moderator, authorised_user, guest;

-- Администратор
CREATE USER admin 
    WITH PASSWORD 'password';

GRANT SELECT, INSERT, UPDATE, DELETE
    ON ALL TABLES IN SCHEMA public 
    TO admin;

-- Модератор
CREATE USER moderator 
    WITH PASSWORD 'password';

GRANT SELECT, INSERT, UPDATE, DELETE
    ON TABLE incidents, users
    TO moderator;

-- Авторизированный пользователь
CREATE USER authorised_user
    WITH PASSWORD 'password';

GRANT SELECT, INSERT
    ON TABLE incidents
    TO authorised_user;

-- Гость
CREATE USER guest
    WITH PASSWORD 'password';

GRANT SELECT
    ON TABLE incidents 
    TO guest;

GRANT INSERT
    ON TABLE users 
    TO guest;
