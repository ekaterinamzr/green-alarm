CREATE FUNCTION count_moderators() RETURNS integer AS $$
    SELECT
        COUNT (*)
    FROM
        users
    WHERE user_role = 2;
$$ LANGUAGE SQL;

CREATE FUNCTION check_moderators() RETURNS TRIGGER AS
$$
BEGIN
	IF (OLD.user_role = 2 and count_moderators() = 1)
	THEN
		RETURN NULL;
	ELSE
		RETURN OLD;
	END IF;
END;
$$ LANGUAGE  plpgsql;

CREATE TRIGGER check_moderators BEFORE DELETE OR UPDATE ON users
    FOR EACH ROW EXECUTE FUNCTION check_moderators();
