CREATE TYPE wishlist_status AS ENUM ('active', 'removed');

CREATE TABLE IF NOT EXISTS wish_lists(
	id uuid DEFAULT uuid_generate_v4 () PRIMARY KEY,
	name VARCHAR (50) NOT NULL,
    status wishlist_status NOT NULL DEFAULT 'active',
	user_id uuid NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
);


/*
* Create trigger to call set_timestamp fucntion
*/
CREATE TRIGGER set_timestamp
BEFORE UPDATE ON wish_lists
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp()