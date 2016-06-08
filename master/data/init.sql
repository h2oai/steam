--
-- PostgreSQL database dump
--

-- Dumped from database version 9.5.0
-- Dumped by pg_dump version 9.5.1

-- Started on 2016-06-03 10:06:43 PDT

SET statement_timeout = 0;
SET lock_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

-- DROP DATABASE steam;

--
-- TOC entry 2503 (class 1262 OID 16389)
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
-- TOC entry 6 (class 2615 OID 2200)
-- Name: public; Type: SCHEMA; Schema: -; Owner: steam
--

CREATE SCHEMA public;


ALTER SCHEMA public OWNER TO steam;

--
-- TOC entry 2504 (class 0 OID 0)
-- Dependencies: 6
-- Name: SCHEMA public; Type: COMMENT; Schema: -; Owner: steam
--

COMMENT ON SCHEMA public IS 'standard public schema';


--
-- TOC entry 1 (class 3079 OID 12623)
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- TOC entry 2506 (class 0 OID 0)
-- Dependencies: 1
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET search_path = public, pg_catalog;

--
-- TOC entry 605 (class 1247 OID 16630)
-- Name: principal_type; Type: TYPE; Schema: public; Owner: steam
--

CREATE TYPE principal_type AS ENUM (
    'Group',
    'Role',
    'User'
);


ALTER TYPE principal_type OWNER TO steam;

--
-- TOC entry 602 (class 1247 OID 16622)
-- Name: privilege_type; Type: TYPE; Schema: public; Owner: steam
--

CREATE TYPE privilege_type AS ENUM (
    'Own',
    'Edit',
    'View'
);


ALTER TYPE privilege_type OWNER TO steam;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- TOC entry 198 (class 1259 OID 16700)
-- Name: cluster; Type: TABLE; Schema: public; Owner: steam
--

CREATE TABLE cluster (
    id integer NOT NULL
);


ALTER TABLE cluster OWNER TO steam;

--
-- TOC entry 201 (class 1259 OID 16709)
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
-- TOC entry 2507 (class 0 OID 0)
-- Dependencies: 201
-- Name: cluster_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: steam
--

ALTER SEQUENCE cluster_id_seq OWNED BY cluster.id;


--
-- TOC entry 197 (class 1259 OID 16697)
-- Name: engine; Type: TABLE; Schema: public; Owner: steam
--

CREATE TABLE engine (
    id integer NOT NULL
);


ALTER TABLE engine OWNER TO steam;

--
-- TOC entry 202 (class 1259 OID 16715)
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
-- TOC entry 2508 (class 0 OID 0)
-- Dependencies: 202
-- Name: engine_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: steam
--

ALTER SEQUENCE engine_id_seq OWNED BY engine.id;


--
-- TOC entry 183 (class 1259 OID 16430)
-- Name: group; Type: TABLE; Schema: public; Owner: steam
--

CREATE TABLE "group" (
    id integer NOT NULL,
    name text,
    created timestamp with time zone,
    modified timestamp with time zone
);


ALTER TABLE "group" OWNER TO steam;

--
-- TOC entry 188 (class 1259 OID 16454)
-- Name: group_id_seq; Type: SEQUENCE; Schema: public; Owner: steam
--

CREATE SEQUENCE group_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE group_id_seq OWNER TO steam;

--
-- TOC entry 2509 (class 0 OID 0)
-- Dependencies: 188
-- Name: group_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: steam
--

ALTER SEQUENCE group_id_seq OWNED BY "group".id;


--
-- TOC entry 187 (class 1259 OID 16451)
-- Name: group_role; Type: TABLE; Schema: public; Owner: steam
--

CREATE TABLE group_role (
    group_id integer NOT NULL,
    role_id integer NOT NULL
);


ALTER TABLE group_role OWNER TO steam;

--
-- TOC entry 191 (class 1259 OID 16551)
-- Name: history; Type: TABLE; Schema: public; Owner: steam
--

CREATE TABLE history (
    id integer NOT NULL,
    "time" timestamp without time zone,
    action text,
    user_id integer,
    entity_type_id integer,
    entity_id integer,
    description text
);


ALTER TABLE history OWNER TO steam;

--
-- TOC entry 190 (class 1259 OID 16549)
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
-- TOC entry 2510 (class 0 OID 0)
-- Dependencies: 190
-- Name: history_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: steam
--

ALTER SEQUENCE history_id_seq OWNED BY history.id;


--
-- TOC entry 200 (class 1259 OID 16706)
-- Name: model; Type: TABLE; Schema: public; Owner: steam
--

CREATE TABLE model (
    id integer NOT NULL
);


ALTER TABLE model OWNER TO steam;

--
-- TOC entry 204 (class 1259 OID 16733)
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
-- TOC entry 2511 (class 0 OID 0)
-- Dependencies: 204
-- Name: model_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: steam
--

ALTER SEQUENCE model_id_seq OWNED BY model.id;


--
-- TOC entry 192 (class 1259 OID 16595)
-- Name: permission; Type: TABLE; Schema: public; Owner: steam
--

CREATE TABLE permission (
    id integer NOT NULL,
    name text
);


ALTER TABLE permission OWNER TO steam;

--
-- TOC entry 193 (class 1259 OID 16598)
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
-- TOC entry 2512 (class 0 OID 0)
-- Dependencies: 193
-- Name: permission_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: steam
--

ALTER SEQUENCE permission_id_seq OWNED BY permission.id;


--
-- TOC entry 194 (class 1259 OID 16637)
-- Name: privilege; Type: TABLE; Schema: public; Owner: steam
--

CREATE TABLE privilege (
    privilege_type privilege_type NOT NULL,
    principal_type principal_type NOT NULL,
    principal_id integer NOT NULL,
    entity_type_id integer NOT NULL,
    entity_id integer NOT NULL
);


ALTER TABLE privilege OWNER TO steam;

--
-- TOC entry 182 (class 1259 OID 16427)
-- Name: role; Type: TABLE; Schema: public; Owner: steam
--

CREATE TABLE role (
    id integer NOT NULL,
    name text,
    created timestamp with time zone,
    modified timestamp with time zone
);


ALTER TABLE role OWNER TO steam;

--
-- TOC entry 189 (class 1259 OID 16460)
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
-- TOC entry 2513 (class 0 OID 0)
-- Dependencies: 189
-- Name: role_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: steam
--

ALTER SEQUENCE role_id_seq OWNED BY role.id;


--
-- TOC entry 185 (class 1259 OID 16445)
-- Name: role_permission; Type: TABLE; Schema: public; Owner: steam
--

CREATE TABLE role_permission (
    role_id integer NOT NULL,
    permission_id integer NOT NULL
);


ALTER TABLE role_permission OWNER TO steam;

--
-- TOC entry 199 (class 1259 OID 16703)
-- Name: service; Type: TABLE; Schema: public; Owner: steam
--

CREATE TABLE service (
    id integer NOT NULL
);


ALTER TABLE service OWNER TO steam;

--
-- TOC entry 203 (class 1259 OID 16721)
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
-- TOC entry 2514 (class 0 OID 0)
-- Dependencies: 203
-- Name: service_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: steam
--

ALTER SEQUENCE service_id_seq OWNED BY service.id;


--
-- TOC entry 195 (class 1259 OID 16640)
-- Name: entity_type; Type: TABLE; Schema: public; Owner: steam
--

CREATE TABLE entity_type (
    id integer NOT NULL,
    name text
);


ALTER TABLE entity_type OWNER TO steam;

--
-- TOC entry 196 (class 1259 OID 16643)
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
-- TOC entry 2515 (class 0 OID 0)
-- Dependencies: 196
-- Name: entity_type_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: steam
--

ALTER SEQUENCE entity_type_id_seq OWNED BY entity_type.id;


--
-- TOC entry 181 (class 1259 OID 16421)
-- Name: user; Type: TABLE; Schema: public; Owner: steam
--

CREATE TABLE "user" (
    id integer NOT NULL,
    name text,
    password text,
    is_locked boolean,
    last_login timestamp without time zone,
    created timestamp with time zone,
    modified timestamp with time zone
);


ALTER TABLE "user" OWNER TO steam;

--
-- TOC entry 186 (class 1259 OID 16448)
-- Name: user_group; Type: TABLE; Schema: public; Owner: steam
--

CREATE TABLE user_group (
    user_id integer NOT NULL,
    group_id integer NOT NULL
);


ALTER TABLE user_group OWNER TO steam;

--
-- TOC entry 184 (class 1259 OID 16433)
-- Name: user_id_seq; Type: SEQUENCE; Schema: public; Owner: steam
--

CREATE SEQUENCE user_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE user_id_seq OWNER TO steam;

--
-- TOC entry 2516 (class 0 OID 0)
-- Dependencies: 184
-- Name: user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: steam
--

ALTER SEQUENCE user_id_seq OWNED BY "user".id;


--
-- TOC entry 2336 (class 2604 OID 16711)
-- Name: id; Type: DEFAULT; Schema: public; Owner: steam
--

ALTER TABLE ONLY cluster ALTER COLUMN id SET DEFAULT nextval('cluster_id_seq'::regclass);


--
-- TOC entry 2335 (class 2604 OID 16717)
-- Name: id; Type: DEFAULT; Schema: public; Owner: steam
--

ALTER TABLE ONLY engine ALTER COLUMN id SET DEFAULT nextval('engine_id_seq'::regclass);


--
-- TOC entry 2331 (class 2604 OID 16456)
-- Name: id; Type: DEFAULT; Schema: public; Owner: steam
--

ALTER TABLE ONLY "group" ALTER COLUMN id SET DEFAULT nextval('group_id_seq'::regclass);


--
-- TOC entry 2332 (class 2604 OID 16554)
-- Name: id; Type: DEFAULT; Schema: public; Owner: steam
--

ALTER TABLE ONLY history ALTER COLUMN id SET DEFAULT nextval('history_id_seq'::regclass);


--
-- TOC entry 2338 (class 2604 OID 16735)
-- Name: id; Type: DEFAULT; Schema: public; Owner: steam
--

ALTER TABLE ONLY model ALTER COLUMN id SET DEFAULT nextval('model_id_seq'::regclass);


--
-- TOC entry 2333 (class 2604 OID 16600)
-- Name: id; Type: DEFAULT; Schema: public; Owner: steam
--

ALTER TABLE ONLY permission ALTER COLUMN id SET DEFAULT nextval('permission_id_seq'::regclass);


--
-- TOC entry 2330 (class 2604 OID 16462)
-- Name: id; Type: DEFAULT; Schema: public; Owner: steam
--

ALTER TABLE ONLY role ALTER COLUMN id SET DEFAULT nextval('role_id_seq'::regclass);


--
-- TOC entry 2337 (class 2604 OID 16723)
-- Name: id; Type: DEFAULT; Schema: public; Owner: steam
--

ALTER TABLE ONLY service ALTER COLUMN id SET DEFAULT nextval('service_id_seq'::regclass);


--
-- TOC entry 2334 (class 2604 OID 16645)
-- Name: id; Type: DEFAULT; Schema: public; Owner: steam
--

ALTER TABLE ONLY entity_type ALTER COLUMN id SET DEFAULT nextval('entity_type_id_seq'::regclass);


--
-- TOC entry 2329 (class 2604 OID 16435)
-- Name: id; Type: DEFAULT; Schema: public; Owner: steam
--

ALTER TABLE ONLY "user" ALTER COLUMN id SET DEFAULT nextval('user_id_seq'::regclass);


--
-- TOC entry 2362 (class 2606 OID 16605)
-- Name: pk_permission; Type: CONSTRAINT; Schema: public; Owner: steam
--

ALTER TABLE ONLY permission
    ADD CONSTRAINT pk_permission PRIMARY KEY (id);


--
-- TOC entry 2371 (class 2606 OID 16730)
-- Name: pk_cluster; Type: CONSTRAINT; Schema: public; Owner: steam
--

ALTER TABLE ONLY cluster
    ADD CONSTRAINT pk_cluster PRIMARY KEY (id);


--
-- TOC entry 2369 (class 2606 OID 16732)
-- Name: pk_engine; Type: CONSTRAINT; Schema: public; Owner: steam
--

ALTER TABLE ONLY engine
    ADD CONSTRAINT pk_engine PRIMARY KEY (id);


--
-- TOC entry 2344 (class 2606 OID 16467)
-- Name: pk_group; Type: CONSTRAINT; Schema: public; Owner: steam
--

ALTER TABLE ONLY "group"
    ADD CONSTRAINT pk_group PRIMARY KEY (id);


--
-- TOC entry 2356 (class 2606 OID 16613)
-- Name: pk_group_role; Type: CONSTRAINT; Schema: public; Owner: steam
--

ALTER TABLE ONLY group_role
    ADD CONSTRAINT pk_group_role PRIMARY KEY (group_id, role_id);


--
-- TOC entry 2360 (class 2606 OID 16556)
-- Name: pk_history; Type: CONSTRAINT; Schema: public; Owner: steam
--

ALTER TABLE ONLY history
    ADD CONSTRAINT pk_history PRIMARY KEY (id);


--
-- TOC entry 2375 (class 2606 OID 16740)
-- Name: pk_model; Type: CONSTRAINT; Schema: public; Owner: steam
--

ALTER TABLE ONLY model
    ADD CONSTRAINT pk_model PRIMARY KEY (id);


--
-- TOC entry 2365 (class 2606 OID 16662)
-- Name: pk_privilege; Type: CONSTRAINT; Schema: public; Owner: steam
--

ALTER TABLE ONLY privilege
    ADD CONSTRAINT pk_privilege PRIMARY KEY (entity_id, principal_id, entity_type_id, principal_type, privilege_type);


--
-- TOC entry 2373 (class 2606 OID 16728)
-- Name: pk_service; Type: CONSTRAINT; Schema: public; Owner: steam
--

ALTER TABLE ONLY service
    ADD CONSTRAINT pk_service PRIMARY KEY (id);


--
-- TOC entry 2367 (class 2606 OID 16654)
-- Name: pk_entity_type; Type: CONSTRAINT; Schema: public; Owner: steam
--

ALTER TABLE ONLY entity_type
    ADD CONSTRAINT pk_entity_type PRIMARY KEY (id);


--
-- TOC entry 2348 (class 2606 OID 16615)
-- Name: pk_role_permission; Type: CONSTRAINT; Schema: public; Owner: steam
--

ALTER TABLE ONLY role_permission
    ADD CONSTRAINT pk_role_permission PRIMARY KEY (role_id, permission_id);


--
-- TOC entry 2342 (class 2606 OID 16469)
-- Name: pk_role; Type: CONSTRAINT; Schema: public; Owner: steam
--

ALTER TABLE ONLY role
    ADD CONSTRAINT pk_role PRIMARY KEY (id);


--
-- TOC entry 2352 (class 2606 OID 16617)
-- Name: pk_user_group; Type: CONSTRAINT; Schema: public; Owner: steam
--

ALTER TABLE ONLY user_group
    ADD CONSTRAINT pk_user_group PRIMARY KEY (user_id, group_id);


--
-- TOC entry 2340 (class 2606 OID 16471)
-- Name: pk_user; Type: CONSTRAINT; Schema: public; Owner: steam
--

ALTER TABLE ONLY "user"
    ADD CONSTRAINT pk_user PRIMARY KEY (id);


--
-- TOC entry 2353 (class 1259 OID 16687)
-- Name: fki_group_role__group_id; Type: INDEX; Schema: public; Owner: steam
--

CREATE INDEX fki_group_role__group_id ON group_role USING btree (group_id);


--
-- TOC entry 2354 (class 1259 OID 16673)
-- Name: fki_group_role__role_id; Type: INDEX; Schema: public; Owner: steam
--

CREATE INDEX fki_group_role__role_id ON group_role USING btree (role_id);


--
-- TOC entry 2357 (class 1259 OID 16693)
-- Name: fki_history__entity_type_id; Type: INDEX; Schema: public; Owner: steam
--

CREATE INDEX fki_history__entity_type_id ON history USING btree (entity_type_id);


--
-- TOC entry 2358 (class 1259 OID 16594)
-- Name: fki_history__user_id; Type: INDEX; Schema: public; Owner: steam
--

CREATE INDEX fki_history__user_id ON history USING btree (user_id);


--
-- TOC entry 2363 (class 1259 OID 16660)
-- Name: fki_privilege__entity_type_id; Type: INDEX; Schema: public; Owner: steam
--

CREATE INDEX fki_privilege__entity_type_id ON privilege USING btree (entity_type_id);


--
-- TOC entry 2345 (class 1259 OID 16685)
-- Name: fki_role_permission__permission_id; Type: INDEX; Schema: public; Owner: steam
--

CREATE INDEX fki_role_permission__permission_id ON role_permission USING btree (permission_id);


--
-- TOC entry 2346 (class 1259 OID 16686)
-- Name: fki_role_permission__role_id; Type: INDEX; Schema: public; Owner: steam
--

CREATE INDEX fki_role_permission__role_id ON role_permission USING btree (role_id);


--
-- TOC entry 2349 (class 1259 OID 16485)
-- Name: fki_user_group__group_id; Type: INDEX; Schema: public; Owner: steam
--

CREATE INDEX fki_user_group__group_id ON user_group USING btree (group_id);


--
-- TOC entry 2350 (class 1259 OID 16479)
-- Name: fki_user_group__user_id; Type: INDEX; Schema: public; Owner: steam
--

CREATE INDEX fki_user_group__user_id ON user_group USING btree (user_id);


--
-- TOC entry 2380 (class 2606 OID 16663)
-- Name: fk_group_role__group_id; Type: FK CONSTRAINT; Schema: public; Owner: steam
--

ALTER TABLE ONLY group_role
    ADD CONSTRAINT fk_group_role__group_id FOREIGN KEY (group_id) REFERENCES "group"(id) ON DELETE CASCADE;


--
-- TOC entry 2381 (class 2606 OID 16668)
-- Name: fk_group_role__role_id; Type: FK CONSTRAINT; Schema: public; Owner: steam
--

ALTER TABLE ONLY group_role
    ADD CONSTRAINT fk_group_role__role_id FOREIGN KEY (role_id) REFERENCES role(id) ON DELETE CASCADE;


--
-- TOC entry 2383 (class 2606 OID 16688)
-- Name: fk_history__entity_type_id; Type: FK CONSTRAINT; Schema: public; Owner: steam
--

ALTER TABLE ONLY history
    ADD CONSTRAINT fk_history__entity_type_id FOREIGN KEY (entity_type_id) REFERENCES entity_type(id);


--
-- TOC entry 2382 (class 2606 OID 16589)
-- Name: fk_history__user_id; Type: FK CONSTRAINT; Schema: public; Owner: steam
--

ALTER TABLE ONLY history
    ADD CONSTRAINT fk_history__user_id FOREIGN KEY (user_id) REFERENCES "user"(id);


--
-- TOC entry 2384 (class 2606 OID 16655)
-- Name: fk_privilege__entity_type_id; Type: FK CONSTRAINT; Schema: public; Owner: steam
--

ALTER TABLE ONLY privilege
    ADD CONSTRAINT fk_privilege__entity_type_id FOREIGN KEY (entity_type_id) REFERENCES entity_type(id);


--
-- TOC entry 2377 (class 2606 OID 16680)
-- Name: fk_role_permission__permission_id; Type: FK CONSTRAINT; Schema: public; Owner: steam
--

ALTER TABLE ONLY role_permission
    ADD CONSTRAINT fk_role_permission__permission_id FOREIGN KEY (permission_id) REFERENCES permission(id) ON DELETE CASCADE;


--
-- TOC entry 2376 (class 2606 OID 16675)
-- Name: fk_role_permission__role_id; Type: FK CONSTRAINT; Schema: public; Owner: steam
--

ALTER TABLE ONLY role_permission
    ADD CONSTRAINT fk_role_permission__role_id FOREIGN KEY (role_id) REFERENCES role(id) ON DELETE CASCADE;


--
-- TOC entry 2378 (class 2606 OID 16516)
-- Name: fk_user_group__group_id; Type: FK CONSTRAINT; Schema: public; Owner: steam
--

ALTER TABLE ONLY user_group
    ADD CONSTRAINT fk_user_group__group_id FOREIGN KEY (group_id) REFERENCES "group"(id) ON DELETE CASCADE;


--
-- TOC entry 2379 (class 2606 OID 16521)
-- Name: fk_user_group__user_id; Type: FK CONSTRAINT; Schema: public; Owner: steam
--

ALTER TABLE ONLY user_group
    ADD CONSTRAINT fk_user_group__user_id FOREIGN KEY (user_id) REFERENCES "user"(id) ON DELETE CASCADE;


--
-- TOC entry 2505 (class 0 OID 0)
-- Dependencies: 6
-- Name: public; Type: ACL; Schema: -; Owner: steam
--

REVOKE ALL ON SCHEMA public FROM PUBLIC;
REVOKE ALL ON SCHEMA public FROM steam;
GRANT ALL ON SCHEMA public TO steam;
GRANT ALL ON SCHEMA public TO PUBLIC;


-- Completed on 2016-06-03 10:06:43 PDT

--
-- PostgreSQL database dump complete
--

