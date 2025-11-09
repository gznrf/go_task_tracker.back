CREATE TABLE IF NOT EXISTS user_info
(
    id            serial not null unique,

    name          text   not null,
    email         text   not null unique,
    password_hash text   not null,

    created_at    timestamp DEFAULT CURRENT_TIMESTAMP,
    deleted_at    timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at    timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS project
(
    id          serial                           not null unique,
    owner_id    bigint REFERENCES user_info (id) not null,

    name        text                             not null,
    description text,

    created_at  timestamp DEFAULT CURRENT_TIMESTAMP,
    deleted_at  timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at  timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS projects_users
(
    id         serial                           not null unique,

    from_id    bigint REFERENCES project (id)   not null,
    to_id      bigint REFERENCES user_info (id) not null,

    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE  IF NOT EXISTS board
(
    id         serial                         not null unique,
    project_id bigint REFERENCES project (id) not null,

    name       text                           not null,

    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE  IF NOT EXISTS column_info
(
    id         serial                       not null unique,
    board_id   bigint REFERENCES board (id) not null,

    name       text                         not null,

    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE  IF NOT EXISTS task
(
    id          serial                             not null unique,
    column_id   bigint REFERENCES column_info (id) not null,
    creator_id  bigint REFERENCES user_info (id)   not null,
    executor_id bigint REFERENCES user_info (id)   not null,

    number      serial                             not null unique,
    name        text                               not null,
    description text,

    created_at  timestamp DEFAULT CURRENT_TIMESTAMP,
    deleted_at  timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at  timestamp DEFAULT CURRENT_TIMESTAMP
);

