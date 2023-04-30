CREATE TABLE IF NOT EXISTS nodes (
	mac_address TEXT NOT NULL,
	name TEXT NOT NULL,
	owned_by UUID,
	custom_labels JSONB,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP,
	adopted_at TIMESTAMP,
    PRIMARY KEY(mac_address)
);
CREATE UNIQUE INDEX ON nodes(owned_by);

CREATE TABLE IF NOT EXISTS sensors (
	id UUID DEFAULT gen_random_uuid(),
	name TEXT NOT NULL,
	node_id TEXT NOT NULL,
	model TEXT,
	category INT
    PRIMARY KEY(id),
	CONSTRAINT fk_sensor_nodes
      FOREIGN KEY(node_id) 
	  REFERENCES nodes(mac_address)
	  ON DELETE CASCADE
);
CREATE UNIQUE INDEX ON sensors(node_id);
