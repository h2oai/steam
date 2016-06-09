--
-- PostgreSQL database dump
--

-- Dumped from database version 9.5.0
-- Dumped by pg_dump version 9.5.1

SET statement_timeout = 0;
SET lock_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

DROP DATABASE steam;
--
-- Name: steam; Type: DATABASE; Schema: -; Owner: steam
--

CREATE DATABASE steam WITH TEMPLATE = template0 ENCODING = 'UTF8' LC_COLLATE = 'en_US.UTF-8' LC_CTYPE = 'en_US.UTF-8';


ALTER DATABASE steam OWNER TO steam;

\connect steam

SET statement_timeout = 0;
SET lock_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: public; Type: SCHEMA; Schema: -; Owner: steam
--

CREATE SCHEMA public;


ALTER SCHEMA public OWNER TO steam;

--
-- Name: SCHEMA public; Type: COMMENT; Schema: -; Owner: steam
--

COMMENT ON SCHEMA public IS 'standard public schema';


--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET search_path = public, pg_catalog;

--
-- Name: identity_type; Type: TYPE; Schema: public; Owner: steam
--

CREATE TYPE identity_type AS ENUM (
    'identity',
    'workgroup'
);


ALTER TYPE identity_type OWNER TO steam;

--
-- Name: privilege_type; Type: TYPE; Schema: public; Owner: steam
--

CREATE TYPE privilege_type AS ENUM (
    'own',
    'edit',
    'view'
);


ALTER TYPE privilege_type OWNER TO steam;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: cluster; Type: TABLE; Schema: public; Owner: steam
--

CREATE TABLE cluster (
    id integer NOT NULL
);


ALTER TABLE cluster OWNER TO steam;

--
-- Name: cluster_id_seq; Type: SEQUENCE; Schema: public; Owner: steam
--

CREATE SEQUENCE cluster_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE cluster_id_seq OWNER TO steam;

--
-- Name: cluster_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: steam
--

ALTER SEQUENCE cluster_id_seq OWNED BY cluster.id;


--
-- Name: engine; Type: TABLE; Schema: public; Owner: steam
--

CREATE TABLE engine (
    id integer NOT NULL
);


ALTER TABLE engine OWNER TO steam;

--
-- Name: engine_id_seq; Type: SEQUENCE; Schema: public; Owner: steam
--

CREATE SEQUENCE engine_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE engine_id_seq OWNER TO steam;

--
-- Name: engine_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: steam
--

ALTER SEQUENCE engine_id_seq OWNED BY engine.id;


--
-- Name: entity_type; Type: TABLE; Schema: public; Owner: steam
--

CREATE TABLE entity_type (
    id integer NOT NULL,
    name text NOT NULL
);


ALTER TABLE entity_type OWNER TO steam;

--
-- Name: entity_type_id_seq; Type: SEQUENCE; Schema: public; Owner: steam
--

CREATE SEQUENCE entity_type_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE entity_type_id_seq OWNER TO steam;

--
-- Name: entity_type_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: steam
--

ALTER SEQUENCE entity_type_id_seq OWNED BY entity_type.id;


--
-- Name: history; Type: TABLE; Schema: public; Owner: steam
--

CREATE TABLE history (
    id integer NOT NULL,
    action text NOT NULL,
    identity_id integer NOT NULL,
    entity_type_id integer NOT NULL,
    entity_id integer NOT NULL,
    description text NOT NULL,
    created timestamp with time zone NOT NULL
);


ALTER TABLE history OWNER TO steam;

--
-- Name: history_id_seq; Type: SEQUENCE; Schema: public; Owner: steam
--

CREATE SEQUENCE history_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE history_id_seq OWNER TO steam;

--
-- Name: history_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: steam
--

ALTER SEQUENCE history_id_seq OWNED BY history.id;


--
-- Name: identity; Type: TABLE; Schema: public; Owner: steam
--

CREATE TABLE identity (
    id integer NOT NULL,
    name text NOT NULL,
    password text NOT NULL,
    is_active boolean NOT NULL,
    last_login timestamp with time zone,
    created timestamp with time zone NOT NULL
);


ALTER TABLE identity OWNER TO steam;

--
-- Name: identity_id_seq; Type: SEQUENCE; Schema: public; Owner: steam
--

CREATE SEQUENCE identity_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE identity_id_seq OWNER TO steam;

--
-- Name: identity_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: steam
--

ALTER SEQUENCE identity_id_seq OWNED BY identity.id;


--
-- Name: identity_role; Type: TABLE; Schema: public; Owner: steam
--

CREATE TABLE identity_role (
    identity_id integer NOT NULL,
    role_id integer NOT NULL
);


ALTER TABLE identity_role OWNER TO steam;

--
-- Name: identity_workgroup; Type: TABLE; Schema: public; Owner: steam
--

CREATE TABLE identity_workgroup (
    identity_id integer NOT NULL,
    workgroup_id integer NOT NULL
);


ALTER TABLE identity_workgroup OWNER TO steam;

--
-- Name: meta; Type: TABLE; Schema: public; Owner: steam
--

CREATE TABLE meta (
    id integer NOT NULL,
    key text NOT NULL,
    value text NOT NULL
);


ALTER TABLE meta OWNER TO steam;

--
-- Name: meta_id_seq; Type: SEQUENCE; Schema: public; Owner: steam
--

CREATE SEQUENCE meta_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE meta_id_seq OWNER TO steam;

--
-- Name: meta_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: steam
--

ALTER SEQUENCE meta_id_seq OWNED BY meta.id;


--
-- Name: model; Type: TABLE; Schema: public; Owner: steam
--

CREATE TABLE model (
    id integer NOT NULL
);


ALTER TABLE model OWNER TO steam;

--
-- Name: model_id_seq; Type: SEQUENCE; Schema: public; Owner: steam
--

CREATE SEQUENCE model_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE model_id_seq OWNER TO steam;

--
-- Name: model_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: steam
--

ALTER SEQUENCE model_id_seq OWNED BY model.id;


--
-- Name: permission; Type: TABLE; Schema: public; Owner: steam
--

CREATE TABLE permission (
    id integer NOT NULL,
    name text NOT NULL,
    description text NOT NULL
);


ALTER TABLE permission OWNER TO steam;

--
-- Name: permission_id_seq; Type: SEQUENCE; Schema: public; Owner: steam
--

CREATE SEQUENCE permission_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE permission_id_seq OWNER TO steam;

--
-- Name: permission_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: steam
--

ALTER SEQUENCE permission_id_seq OWNED BY permission.id;


--
-- Name: privilege; Type: TABLE; Schema: public; Owner: steam
--

CREATE TABLE privilege (
    privilege_type privilege_type NOT NULL,
    identity_type identity_type NOT NULL,
    identity_id integer NOT NULL,
    entity_type_id integer NOT NULL,
    entity_id integer NOT NULL
);


ALTER TABLE privilege OWNER TO steam;

--
-- Name: role; Type: TABLE; Schema: public; Owner: steam
--

CREATE TABLE role (
    id integer NOT NULL,
    name text NOT NULL,
    description text NOT NULL,
    created timestamp with time zone NOT NULL
);


ALTER TABLE role OWNER TO steam;

--
-- Name: role_id_seq; Type: SEQUENCE; Schema: public; Owner: steam
--

CREATE SEQUENCE role_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE role_id_seq OWNER TO steam;

--
-- Name: role_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: steam
--

ALTER SEQUENCE role_id_seq OWNED BY role.id;


--
-- Name: role_permission; Type: TABLE; Schema: public; Owner: steam
--

CREATE TABLE role_permission (
    role_id integer NOT NULL,
    permission_id integer NOT NULL
);


ALTER TABLE role_permission OWNER TO steam;

--
-- Name: service; Type: TABLE; Schema: public; Owner: steam
--

CREATE TABLE service (
    id integer NOT NULL
);


ALTER TABLE service OWNER TO steam;

--
-- Name: service_id_seq; Type: SEQUENCE; Schema: public; Owner: steam
--

CREATE SEQUENCE service_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE service_id_seq OWNER TO steam;

--
-- Name: service_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: steam
--

ALTER SEQUENCE service_id_seq OWNED BY service.id;


--
-- Name: workgroup; Type: TABLE; Schema: public; Owner: steam
--

CREATE TABLE workgroup (
    id integer NOT NULL,
    name text NOT NULL,
    description text NOT NULL,
    created timestamp with time zone NOT NULL
);


ALTER TABLE workgroup OWNER TO steam;

--
-- Name: workgroup_id_seq; Type: SEQUENCE; Schema: public; Owner: steam
--

CREATE SEQUENCE workgroup_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE workgroup_id_seq OWNER TO steam;

--
-- Name: workgroup_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: steam
--

ALTER SEQUENCE workgroup_id_seq OWNED BY workgroup.id;


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: steam
--

ALTER TABLE ONLY cluster ALTER COLUMN id SET DEFAULT nextval('cluster_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: steam
--

ALTER TABLE ONLY engine ALTER COLUMN id SET DEFAULT nextval('engine_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: steam
--

ALTER TABLE ONLY entity_type ALTER COLUMN id SET DEFAULT nextval('entity_type_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: steam
--

ALTER TABLE ONLY history ALTER COLUMN id SET DEFAULT nextval('history_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: steam
--

ALTER TABLE ONLY identity ALTER COLUMN id SET DEFAULT nextval('identity_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: steam
--

ALTER TABLE ONLY meta ALTER COLUMN id SET DEFAULT nextval('meta_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: steam
--

ALTER TABLE ONLY model ALTER COLUMN id SET DEFAULT nextval('model_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: steam
--

ALTER TABLE ONLY permission ALTER COLUMN id SET DEFAULT nextval('permission_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: steam
--

ALTER TABLE ONLY role ALTER COLUMN id SET DEFAULT nextval('role_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: steam
--

ALTER TABLE ONLY service ALTER COLUMN id SET DEFAULT nextval('service_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: steam
--

ALTER TABLE ONLY workgroup ALTER COLUMN id SET DEFAULT nextval('workgroup_id_seq'::regclass);


--
-- Name: pk_cluster; Type: CONSTRAINT; Schema: public; Owner: steam
--

ALTER TABLE ONLY cluster
    ADD CONSTRAINT pk_cluster PRIMARY KEY (id);


--
-- Name: pk_engine; Type: CONSTRAINT; Schema: public; Owner: steam
--

ALTER TABLE ONLY engine
    ADD CONSTRAINT pk_engine PRIMARY KEY (id);


--
-- Name: pk_entity_type; Type: CONSTRAINT; Schema: public; Owner: steam
--

ALTER TABLE ONLY entity_type
    ADD CONSTRAINT pk_entity_type PRIMARY KEY (id);


--
-- Name: pk_history; Type: CONSTRAINT; Schema: public; Owner: steam
--

ALTER TABLE ONLY history
    ADD CONSTRAINT pk_history PRIMARY KEY (id);


--
-- Name: pk_identity; Type: CONSTRAINT; Schema: public; Owner: steam
--

ALTER TABLE ONLY identity
    ADD CONSTRAINT pk_identity PRIMARY KEY (id);


--
-- Name: pk_identity_role; Type: CONSTRAINT; Schema: public; Owner: steam
--

ALTER TABLE ONLY identity_role
    ADD CONSTRAINT pk_identity_role PRIMARY KEY (identity_id, role_id);


--
-- Name: pk_identity_workgroup; Type: CONSTRAINT; Schema: public; Owner: steam
--

ALTER TABLE ONLY identity_workgroup
    ADD CONSTRAINT pk_identity_workgroup PRIMARY KEY (identity_id, workgroup_id);


--
-- Name: pk_meta; Type: CONSTRAINT; Schema: public; Owner: steam
--

ALTER TABLE ONLY meta
    ADD CONSTRAINT pk_meta PRIMARY KEY (id);


--
-- Name: pk_model; Type: CONSTRAINT; Schema: public; Owner: steam
--

ALTER TABLE ONLY model
    ADD CONSTRAINT pk_model PRIMARY KEY (id);


--
-- Name: pk_permission; Type: CONSTRAINT; Schema: public; Owner: steam
--

ALTER TABLE ONLY permission
    ADD CONSTRAINT pk_permission PRIMARY KEY (id);


--
-- Name: pk_privilege; Type: CONSTRAINT; Schema: public; Owner: steam
--

ALTER TABLE ONLY privilege
    ADD CONSTRAINT pk_privilege PRIMARY KEY (entity_id, identity_id, entity_type_id, identity_type, privilege_type);


--
-- Name: pk_role; Type: CONSTRAINT; Schema: public; Owner: steam
--

ALTER TABLE ONLY role
    ADD CONSTRAINT pk_role PRIMARY KEY (id);


--
-- Name: pk_role_permission; Type: CONSTRAINT; Schema: public; Owner: steam
--

ALTER TABLE ONLY role_permission
    ADD CONSTRAINT pk_role_permission PRIMARY KEY (role_id, permission_id);


--
-- Name: pk_service; Type: CONSTRAINT; Schema: public; Owner: steam
--

ALTER TABLE ONLY service
    ADD CONSTRAINT pk_service PRIMARY KEY (id);


--
-- Name: pk_workgroup; Type: CONSTRAINT; Schema: public; Owner: steam
--

ALTER TABLE ONLY workgroup
    ADD CONSTRAINT pk_workgroup PRIMARY KEY (id);


--
-- Name: uq_entity_type_name; Type: CONSTRAINT; Schema: public; Owner: steam
--

ALTER TABLE ONLY entity_type
    ADD CONSTRAINT uq_entity_type_name UNIQUE (name);


--
-- Name: uq_identity_name; Type: CONSTRAINT; Schema: public; Owner: steam
--

ALTER TABLE ONLY identity
    ADD CONSTRAINT uq_identity_name UNIQUE (name);


--
-- Name: uq_meta_key; Type: CONSTRAINT; Schema: public; Owner: steam
--

ALTER TABLE ONLY meta
    ADD CONSTRAINT uq_meta_key UNIQUE (key);


--
-- Name: uq_permission_name; Type: CONSTRAINT; Schema: public; Owner: steam
--

ALTER TABLE ONLY permission
    ADD CONSTRAINT uq_permission_name UNIQUE (name);


--
-- Name: uq_role_name; Type: CONSTRAINT; Schema: public; Owner: steam
--

ALTER TABLE ONLY role
    ADD CONSTRAINT uq_role_name UNIQUE (name);


--
-- Name: uq_workgroup_name; Type: CONSTRAINT; Schema: public; Owner: steam
--

ALTER TABLE ONLY workgroup
    ADD CONSTRAINT uq_workgroup_name UNIQUE (name);


--
-- Name: fki_history__entity_type_id; Type: INDEX; Schema: public; Owner: steam
--

CREATE INDEX fki_history__entity_type_id ON history USING btree (entity_type_id);


--
-- Name: fki_history__identity_id; Type: INDEX; Schema: public; Owner: steam
--

CREATE INDEX fki_history__identity_id ON history USING btree (identity_id);


--
-- Name: fki_identity_workgroup__identity_id; Type: INDEX; Schema: public; Owner: steam
--

CREATE INDEX fki_identity_workgroup__identity_id ON identity_workgroup USING btree (identity_id);


--
-- Name: fki_identity_workgroup__workgroup_id; Type: INDEX; Schema: public; Owner: steam
--

CREATE INDEX fki_identity_workgroup__workgroup_id ON identity_workgroup USING btree (workgroup_id);


--
-- Name: fki_privilege__entity_type_id; Type: INDEX; Schema: public; Owner: steam
--

CREATE INDEX fki_privilege__entity_type_id ON privilege USING btree (entity_type_id);


--
-- Name: fki_role_permission__permission_id; Type: INDEX; Schema: public; Owner: steam
--

CREATE INDEX fki_role_permission__permission_id ON role_permission USING btree (permission_id);


--
-- Name: fki_role_permission__role_id; Type: INDEX; Schema: public; Owner: steam
--

CREATE INDEX fki_role_permission__role_id ON role_permission USING btree (role_id);


--
-- Name: fk_history__entity_type_id; Type: FK CONSTRAINT; Schema: public; Owner: steam
--

ALTER TABLE ONLY history
    ADD CONSTRAINT fk_history__entity_type_id FOREIGN KEY (entity_type_id) REFERENCES entity_type(id);


--
-- Name: fk_history__identity_id; Type: FK CONSTRAINT; Schema: public; Owner: steam
--

ALTER TABLE ONLY history
    ADD CONSTRAINT fk_history__identity_id FOREIGN KEY (identity_id) REFERENCES identity(id);


--
-- Name: fk_identity_workgroup__identity_id; Type: FK CONSTRAINT; Schema: public; Owner: steam
--

ALTER TABLE ONLY identity_workgroup
    ADD CONSTRAINT fk_identity_workgroup__identity_id FOREIGN KEY (identity_id) REFERENCES identity(id) ON DELETE CASCADE;


--
-- Name: fk_identity_workgroup__workgroup_id; Type: FK CONSTRAINT; Schema: public; Owner: steam
--

ALTER TABLE ONLY identity_workgroup
    ADD CONSTRAINT fk_identity_workgroup__workgroup_id FOREIGN KEY (workgroup_id) REFERENCES workgroup(id) ON DELETE CASCADE;


--
-- Name: fk_privilege__entity_type_id; Type: FK CONSTRAINT; Schema: public; Owner: steam
--

ALTER TABLE ONLY privilege
    ADD CONSTRAINT fk_privilege__entity_type_id FOREIGN KEY (entity_type_id) REFERENCES entity_type(id);


--
-- Name: fk_role_permission__permission_id; Type: FK CONSTRAINT; Schema: public; Owner: steam
--

ALTER TABLE ONLY role_permission
    ADD CONSTRAINT fk_role_permission__permission_id FOREIGN KEY (permission_id) REFERENCES permission(id) ON DELETE CASCADE;


--
-- Name: fk_role_permission__role_id; Type: FK CONSTRAINT; Schema: public; Owner: steam
--

ALTER TABLE ONLY role_permission
    ADD CONSTRAINT fk_role_permission__role_id FOREIGN KEY (role_id) REFERENCES role(id) ON DELETE CASCADE;


--
-- Name: public; Type: ACL; Schema: -; Owner: steam
--

REVOKE ALL ON SCHEMA public FROM PUBLIC;
GRANT ALL ON SCHEMA public TO steam;
GRANT ALL ON SCHEMA public TO PUBLIC;


--
-- PostgreSQL database dump complete
--

