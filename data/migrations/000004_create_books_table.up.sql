CREATE TYPE book_status AS ENUM ('active', 'removed');

CREATE TABLE IF NOT EXISTS books(
	id uuid DEFAULT uuid_generate_v4 () PRIMARY KEY,
    google_id VARCHAR (50) NOT NULL,
	author VARCHAR (50) NOT NULL,
    title VARCHAR (50) NOT NULL,
    status book_status DEFAULT 'active',
    publisher VARCHAR (50) NOT NULL,
	user_id uuid NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
);


/*
* Create trigger to call set_timestamp fucntion
*/
CREATE TRIGGER set_timestamp
BEFORE UPDATE ON books
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp()