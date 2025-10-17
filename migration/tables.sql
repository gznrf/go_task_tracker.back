CREATE TABLE user_info
(
    id            serial       not null unique,
    name          varchar(255) not null,
    password_hash varchar(255) not null
);

CREATE TABLE project
(
    id serial not null unique,
    Name text not null,
    Description text,
    Owner_id bigint not null
);

CREATE TABLE projects_users
(
    id      serial not null unique,
    from_id bigint not null,
    to_id   bigint not null
);