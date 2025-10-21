CREATE TABLE user_info
(
    id            serial       not null unique,
    name          varchar(255) not null,
    password_hash varchar(255) not null
);

CREATE TABLE project
(
    id serial not null unique,
    name text not null,
    description text,
    owner_id bigint not null
);

CREATE TABLE projects_users
(
    id      serial not null unique,
    from_id bigint not null,
    to_id   bigint not null
);

CREATE TABLE board(
    id serial not null unique,
    name text not null,
    project_id bigint not null
);

