CREATE TABLE IF NOT EXISTS sessions (
	id uuid DEFAULT uuid_generate_v4 (),
	user_id uuid NOT NULL,
	created_at TIMESTAMP WITH TIME ZONE NOT NULL,
	expired_at TIMESTAMP WITH TIME ZONE NOT NULL,
	PRIMARY KEY(id),
	CONSTRAINT fk_session
		FOREIGN KEY(user_id) 
		REFERENCES users(id)
		ON DELETE CASCADE
);