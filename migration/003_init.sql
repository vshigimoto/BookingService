CREATE TABLE IF NOT EXISTS "user_token"
(
    "user_id"                serial PRIMARY KEY,
    "token"                  VARCHAR(500)        NOT NULL,
    "refresh_token"          VARCHAR(500)        NOT NULL
);

CREATE TABLE IF NOT EXISTS "user_role"
(
    "id"                       serial              PRIMARY KEY,
    "user_id"                  int                 NOT NULL,
    "role"                     VARCHAR(255)        NOT NULL
);

insert into user_role (id, user_id, role) values (1, 1, 'user');
insert into user_role (id, user_id, role) values (2, 2, 'user');
insert into user_role (id, user_id, role) values (3, 3, 'user');
insert into user_role (id, user_id, role) values (4, 4, 'user');
insert into user_role (id, user_id, role) values (5, 5, 'user');
insert into user_role (id, user_id, role) values (6, 6, 'admin');
insert into user_role (id, user_id, role) values (7, 7, 'user');
insert into user_role (id, user_id, role) values (8, 8, 'user');
insert into user_role (id, user_id, role) values (9, 9, 'user');
insert into user_role (id, user_id, role) values (10, 10, 'admin');
insert into user_role (id, user_id, role) values (11, 11, 'admin');
insert into user_role (id, user_id, role) values (12, 12, 'admin');
insert into user_role (id, user_id, role) values (13, 13, 'user');
insert into user_role (id, user_id, role) values (14, 14, 'user');
insert into user_role (id, user_id, role) values (15, 15, 'user');
insert into user_role (id, user_id, role) values (16, 16, 'user');
insert into user_role (id, user_id, role) values (17, 17, 'user');
insert into user_role (id, user_id, role) values (18, 18, 'user');
insert into user_role (id, user_id, role) values (19, 19, 'user');
insert into user_role (id, user_id, role) values (20, 20, 'admin');
