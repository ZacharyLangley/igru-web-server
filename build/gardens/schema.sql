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

CREATE TABLE IF NOT EXISTS plants (
	id UUID DEFAULT gen_random_uuid(),
	name TEXT NOT NULL,
	comment TEXT NOT NULL,
	notes TEXT NOT NULL,
    grow_cycle_length TEXT NOT NULL,
	parentage UUID NOT NULL,
	origin TEXT NOT NULL,
	gender TEXT NOT NULL,
    days_flowering FLOAT NOT NULL,
    days_cured FLOAT NOT NULL,
    harvested_weight TEXT NOT NULL,
    bud_density FLOAT NOT NULL,
    bud_pistils BOOLEAN NOT NULL DEFAULT FALSE,
    tags TEXT NOT NULL,
    acquired_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP,
    PRIMARY KEY(id)
);
CREATE UNIQUE INDEX ON plants(id);

CREATE TABLE IF NOT EXISTS strains (
	id UUID DEFAULT gen_random_uuid(),
	name TEXT NOT NULL,
	comment TEXT NOT NULL,
	notes TEXT NOT NULL,
    type TEXT NOT NULL,
	price FLOAT NOT NULL,
	thc_percent FLOAT NOT NULL,
    cbd_percent FLOAT NOT NULL,
	parentage UUID NOT NULL,
	aroma TEXT NOT NULL,
    taste TEXT NOT NULL,
    tags TEXT NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP,
    PRIMARY KEY(id)
);
CREATE UNIQUE INDEX ON strains(id);