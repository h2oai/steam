--
-- Steam PostgreSQL database init
--

\connect steam

SET statement_timeout = 0;
SET lock_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

-- Meta

INSERT INTO meta (key, value) VALUES ('version', '1');

-- Permissions

INSERT INTO permission (name, description) VALUES ('role.manage', '');
INSERT INTO permission (name, description) VALUES ('role.view', '');
INSERT INTO permission (name, description) VALUES ('workgroup.manage', '');
INSERT INTO permission (name, description) VALUES ('workgroup.view', '');
INSERT INTO permission (name, description) VALUES ('identity.manage', '');
INSERT INTO permission (name, description) VALUES ('identity.view', '');
INSERT INTO permission (name, description) VALUES ('engine.manage', '');
INSERT INTO permission (name, description) VALUES ('engine.view', '');
INSERT INTO permission (name, description) VALUES ('cluster.manage', '');
INSERT INTO permission (name, description) VALUES ('cluster.view', '');
INSERT INTO permission (name, description) VALUES ('model.manage', '');
INSERT INTO permission (name, description) VALUES ('model.view', '');

-- Entity Types

INSERT INTO entity_type (name) VALUES ('role');
INSERT INTO entity_type (name) VALUES ('workgroup');
INSERT INTO entity_type (name) VALUES ('identity');
INSERT INTO entity_type (name) VALUES ('engine');
INSERT INTO entity_type (name) VALUES ('cluster');
INSERT INTO entity_type (name) VALUES ('model');

-- Roles

INSERT INTO role (name, description, created) VALUES ('Superuser', '', now());
INSERT INTO role (name, description, created) VALUES ('Administrator', '', now());

-- Identities

-- FIXME *remove* - should be created by application while priming.
INSERT INTO identity (name, password, is_active, created) VALUES ('Superuser', '', true, now());

-- Workgroups

INSERT INTO workgroup (name, description, created) VALUES ('Everyone', '', now());



