CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE IF NOT EXISTS users (
	id uuid DEFAULT uuid_generate_v4 (),
	email VARCHAR(255) NOT NULL,
	first_name VARCHAR(255) NOT NULL,
	last_name VARCHAR(255) NOT NULL,
	active BOOLEAN NOT NULL,
	salt VARCHAR(255) NOT NULL,
	hash VARCHAR(255) NOT NULL,
	created_at TIMESTAMP WITH TIME ZONE NOT NULL,
	updated_at TIMESTAMP WITH TIME ZONE,
	deleted_at TIMESTAMP WITH TIME ZONE
);
CREATE UNIQUE INDEX user_id_idx ON users(id);
CREATE UNIQUE INDEX user_email_idx ON users(email);
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
CREATE INDEX user_email_idx ON users(email);