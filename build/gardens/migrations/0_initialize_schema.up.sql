CREATE TABLE IF NOT EXISTS gardens (
	id UUID DEFAULT gen_random_uuid(),
	name TEXT NOT NULL,
	comment TEXT NOT NULL,
	location TEXT NOT NULL,
    grow_type TEXT NOT NULL,
	grow_size TEXT NOT NULL,
	grow_style TEXT NOT NULL,
	container_size TEXT NOT NULL,
    tags TEXT NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP,
    PRIMARY KEY(id)
);
CREATE UNIQUE INDEX ON gardens(id);