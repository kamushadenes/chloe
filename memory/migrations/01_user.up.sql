CREATE TABLE users
(
    id         TEXT PRIMARY KEY,
    first_name TEXT NULL,
    last_name  TEXT NULL,
    username   TEXT NULL
);

CREATE TABLE user_external_ids
(
    user_id     TEXT NOT NULL,
    external_id TEXT NOT NULL,
    interface   TEXT NOT NULL
);

CREATE INDEX user_external_ids_user_id_idx ON user_external_ids (user_id);
CREATE UNIQUE INDEX user_external_ids_external_id_idx ON user_external_ids (external_id, interface);

INSERT INTO users (id, first_name, last_name, username)
VALUES ('1', 'Henrique', 'Goncalves', 'kamushadenes');
