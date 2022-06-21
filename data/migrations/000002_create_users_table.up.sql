CREATE TABLE IF NOT EXISTS users(
	id uuid DEFAULT uuid_generate_v4 () PRIMARY KEY,
	user_name VARCHAR (50) UNIQUE NOT NULL,
	password TEXT NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);


/*
* Create trigger to call set_timestamp fucntion
*/
CREATE TRIGGER set_timestamp
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp()