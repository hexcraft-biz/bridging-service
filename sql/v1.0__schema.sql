CREATE TABLE IF NOT EXISTS endpoints(
    `id` BINARY(16) NOT NULL,
    `path` VARCHAR(127) NOT NULL,
    `ctime` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `mtime` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY(`id`),
    UNIQUE(`path`)
) ENGINE InnoDB COLLATE 'utf8mb4_unicode_ci' CHARACTER SET 'utf8mb4';

CREATE TABLE IF NOT EXISTS topics(
    `id` BINARY(16) NOT NULL,
    `name` VARCHAR(255) NOT NULL,
    `ctime` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `mtime` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY(`id`),
    UNIQUE(`name`)
) ENGINE InnoDB COLLATE 'utf8mb4_unicode_ci' CHARACTER SET 'utf8mb4';

CREATE TABLE IF NOT EXISTS endpoint_topic_rels(
    `id` BINARY(16) NOT NULL,
	`endpoint_id` BINARY(16) NOT NULL,
	`topic_id` BINARY(16) NOT NULL,
    `ctime` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `mtime` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY(`id`),
    CONSTRAINT `unique_idx` UNIQUE (`endpoint_id`,`topic_id`),
    FOREIGN KEY(endpoint_id) REFERENCES endpoints(id) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY(topic_id) REFERENCES topics(id) ON UPDATE CASCADE ON DELETE CASCADE
) ENGINE InnoDB COLLATE 'utf8mb4_unicode_ci' CHARACTER SET 'utf8mb4';

CREATE OR REPLACE VIEW view_endpoint_topic_rels AS SELECT
	etr.id,
	etr.endpoint_id,
    e.path AS endpoint_path,
	etr.topic_id,
	t.name AS topic_name,
	etr.ctime,
	etr.mtime
FROM
	endpoint_topic_rels etr
LEFT JOIN endpoints AS e ON etr.endpoint_id = e.id
LEFT JOIN topics AS t ON etr.topic_id = t.id;
