create schema config;

CREATE TABLE config.Roles(
    roleId int not null primary key,
    roleName varchar not null unique,
    roleDescription varchar null,
);

INSERT INTO config.Roles (roleId, roleName) VALUES (0, 'Final Boss');
INSERT INTO config.Roles (roleId, roleName) VALUES (2, ''Managing Director'');
INSERT INTO config.Roles (roleId, roleName) VALUES (3, ''HR Administrator'');
INSERT INTO config.Roles (roleId, roleName) VALUES (5, ''Programmer'');
INSERT INTO config.Roles (roleId, roleName) VALUES (6, ''System Administrator'');
