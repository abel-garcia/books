CREATE TABLE IF NOT EXISTS wish_list_books(
	id uuid DEFAULT uuid_generate_v4 () PRIMARY KEY,
	user_id uuid NOT NULL,
    book_id uuid NOT NULL,
    wish_list_id uuid NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY(wish_list_id) REFERENCES wish_lists(id) ON DELETE CASCADE,
    FOREIGN KEY(book_id) REFERENCES books(id) ON DELETE CASCADE
);


/*
* Create trigger to call set_timestamp fucntion
*/
CREATE TRIGGER set_timestamp
BEFORE UPDATE ON wish_list_books
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp()