CREATE TABLE IF NOT EXISTS groups (
	id UUID NOT NULL DEFAULT gen_random_uuid(),
	name TEXT NOT NULL,
	user_group BOOLEAN NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP,
    PRIMARY KEY(id)
);
CREATE UNIQUE INDEX ON groups(id);

CREATE TABLE IF NOT EXISTS users (
	id UUID DEFAULT gen_random_uuid(),
	email TEXT NOT NULL,
	group_id UUID NOT NULL,
	full_name TEXT,
	active BOOL DEFAULT TRUE,
	salt TEXT NOT NULL,
	hash TEXT NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP,
    PRIMARY KEY(id),
	CONSTRAINT fk_user_group
      FOREIGN KEY(group_id) 
	  REFERENCES groups(id)
	  ON DELETE CASCADE
);
CREATE UNIQUE INDEX ON users(id);
CREATE UNIQUE INDEX ON users(email);

CREATE TABLE IF NOT EXISTS group_members (
	user_id UUID NOT NULL,
	group_id UUID NOT NULL,
	role INTEGER NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP,
    PRIMARY KEY(user_id, group_id),
	CONSTRAINT fk_users
      FOREIGN KEY(user_id) 
	  REFERENCES users(id)
	  ON DELETE CASCADE,
	CONSTRAINT fk_groups
      FOREIGN KEY(group_id) 
	  REFERENCES groups(id)
	  ON DELETE CASCADE
);
CREATE INDEX ON group_members(user_id);
CREATE INDEX ON group_members(group_id);
CREATE TABLE IF NOT EXISTS sessions (
	id UUID NOT NULL DEFAULT gen_random_uuid(),
	user_id UUID NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	expired_at TIMESTAMP NOT NULL,
	PRIMARY KEY(id),
	CONSTRAINT fk_session
		FOREIGN KEY(user_id) 
		REFERENCES users(id)
		ON DELETE CASCADE
);