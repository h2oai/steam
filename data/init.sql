--
-- PostgreSQL database dump
--

-- Dumped from database version 9.5.0
-- Dumped by pg_dump version 9.5.1

-- Started on 2016-06-02 17:41:05 PDT

SET statement_timeout = 0;
SET lock_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- TOC entry 1 (class 3079 OID 12623)
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- TOC entry 2470 (class 0 OID 0)
-- Dependencies: 1
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: -
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET search_path = public, pg_catalog;

--
-- TOC entry 597 (class 1247 OID 16630)
-- Name: principal_type; Type: TYPE; Schema: public; Owner: -
--

CREATE TYPE principal_type AS ENUM (
    'Group',
    'Role',
    'User'
);


--
-- TOC entry 594 (class 1247 OID 16622)
-- Name: privilege_type; Type: TYPE; Schema: public; Owner: -
--

CREATE TYPE privilege_type AS ENUM (
    'Own',
    'Edit',
    'View'
);


SET default_tablespace = '';

SET default_with_oids = false;

--
-- TOC entry 195 (class 1259 OID 16640)
-- Name: subject; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE subject (
    id integer NOT NULL,
    name text
);


--
-- TOC entry 196 (class 1259 OID 16643)
-- Name: entity_type_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE entity_type_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- TOC entry 2471 (class 0 OID 0)
-- Dependencies: 196
-- Name: entity_type_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE entity_type_id_seq OWNED BY subject.id;


--
-- TOC entry 183 (class 1259 OID 16430)
-- Name: group; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE "group" (
    id integer NOT NULL,
    name text,
    created timestamp with time zone,
    modified timestamp with time zone
);


--
-- TOC entry 188 (class 1259 OID 16454)
-- Name: group_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE group_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- TOC entry 2472 (class 0 OID 0)
-- Dependencies: 188
-- Name: group_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE group_id_seq OWNED BY "group".id;


--
-- TOC entry 187 (class 1259 OID 16451)
-- Name: group_role; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE group_role (
    group_id integer NOT NULL,
    role_id integer NOT NULL
);


--
-- TOC entry 191 (class 1259 OID 16551)
-- Name: history; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE history (
    id integer NOT NULL,
    "time" timestamp without time zone,
    action text,
    user_id integer,
    entity_id integer,
    description text,
    subject_id integer
);


--
-- TOC entry 190 (class 1259 OID 16549)
-- Name: history_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE history_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- TOC entry 2473 (class 0 OID 0)
-- Dependencies: 190
-- Name: history_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE history_id_seq OWNED BY history.id;


--
-- TOC entry 192 (class 1259 OID 16595)
-- Name: permission; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE permission (
    id integer NOT NULL,
    name text
);


--
-- TOC entry 193 (class 1259 OID 16598)
-- Name: permission_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE permission_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- TOC entry 2474 (class 0 OID 0)
-- Dependencies: 193
-- Name: permission_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE permission_id_seq OWNED BY permission.id;


--
-- TOC entry 194 (class 1259 OID 16637)
-- Name: privilege; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE privilege (
    privilege_type privilege_type NOT NULL,
    principal_type principal_type NOT NULL,
    principal_id integer NOT NULL,
    subject_id integer NOT NULL,
    entity_id integer NOT NULL
);


--
-- TOC entry 182 (class 1259 OID 16427)
-- Name: role; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE role (
    id integer NOT NULL,
    name text,
    created timestamp with time zone,
    modified timestamp with time zone
);


--
-- TOC entry 189 (class 1259 OID 16460)
-- Name: role_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE role_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- TOC entry 2475 (class 0 OID 0)
-- Dependencies: 189
-- Name: role_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE role_id_seq OWNED BY role.id;


--
-- TOC entry 185 (class 1259 OID 16445)
-- Name: role_permission; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE role_permission (
    role_id integer NOT NULL,
    permission_id integer NOT NULL
);


--
-- TOC entry 181 (class 1259 OID 16421)
-- Name: user; Type: TABLE; Schema: public; Owner: -
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


--
-- TOC entry 186 (class 1259 OID 16448)
-- Name: user_group; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE user_group (
    user_id integer NOT NULL,
    group_id integer NOT NULL
);


--
-- TOC entry 184 (class 1259 OID 16433)
-- Name: user_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE user_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- TOC entry 2476 (class 0 OID 0)
-- Dependencies: 184
-- Name: user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE user_id_seq OWNED BY "user".id;


--
-- TOC entry 2307 (class 2604 OID 16456)
-- Name: id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY "group" ALTER COLUMN id SET DEFAULT nextval('group_id_seq'::regclass);


--
-- TOC entry 2308 (class 2604 OID 16554)
-- Name: id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY history ALTER COLUMN id SET DEFAULT nextval('history_id_seq'::regclass);


--
-- TOC entry 2309 (class 2604 OID 16600)
-- Name: id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY permission ALTER COLUMN id SET DEFAULT nextval('permission_id_seq'::regclass);


--
-- TOC entry 2306 (class 2604 OID 16462)
-- Name: id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY role ALTER COLUMN id SET DEFAULT nextval('role_id_seq'::regclass);


--
-- TOC entry 2310 (class 2604 OID 16645)
-- Name: id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY subject ALTER COLUMN id SET DEFAULT nextval('entity_type_id_seq'::regclass);


--
-- TOC entry 2305 (class 2604 OID 16435)
-- Name: id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY "user" ALTER COLUMN id SET DEFAULT nextval('user_id_seq'::regclass);


--
-- TOC entry 2334 (class 2606 OID 16605)
-- Name: permission_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY permission
    ADD CONSTRAINT permission_pkey PRIMARY KEY (id);


--
-- TOC entry 2316 (class 2606 OID 16467)
-- Name: pk_group; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY "group"
    ADD CONSTRAINT pk_group PRIMARY KEY (id);


--
-- TOC entry 2328 (class 2606 OID 16613)
-- Name: pk_group_role; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY group_role
    ADD CONSTRAINT pk_group_role PRIMARY KEY (group_id, role_id);


--
-- TOC entry 2332 (class 2606 OID 16556)
-- Name: pk_history; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY history
    ADD CONSTRAINT pk_history PRIMARY KEY (id);


--
-- TOC entry 2337 (class 2606 OID 16662)
-- Name: pk_privilege; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY privilege
    ADD CONSTRAINT pk_privilege PRIMARY KEY (entity_id, principal_id, subject_id, principal_type, privilege_type);


--
-- TOC entry 2339 (class 2606 OID 16654)
-- Name: pk_subject; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY subject
    ADD CONSTRAINT pk_subject PRIMARY KEY (id);


--
-- TOC entry 2320 (class 2606 OID 16615)
-- Name: role_permission_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY role_permission
    ADD CONSTRAINT role_permission_pkey PRIMARY KEY (role_id, permission_id);


--
-- TOC entry 2314 (class 2606 OID 16469)
-- Name: role_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY role
    ADD CONSTRAINT role_pkey PRIMARY KEY (id);


--
-- TOC entry 2324 (class 2606 OID 16617)
-- Name: user_group_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY user_group
    ADD CONSTRAINT user_group_pkey PRIMARY KEY (user_id, group_id);


--
-- TOC entry 2312 (class 2606 OID 16471)
-- Name: user_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY "user"
    ADD CONSTRAINT user_pkey PRIMARY KEY (id);


--
-- TOC entry 2325 (class 1259 OID 16687)
-- Name: fki_group_role__group_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX fki_group_role__group_id ON group_role USING btree (group_id);


--
-- TOC entry 2326 (class 1259 OID 16673)
-- Name: fki_group_role__role_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX fki_group_role__role_id ON group_role USING btree (role_id);


--
-- TOC entry 2329 (class 1259 OID 16693)
-- Name: fki_history__subject_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX fki_history__subject_id ON history USING btree (subject_id);


--
-- TOC entry 2330 (class 1259 OID 16594)
-- Name: fki_history__user_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX fki_history__user_id ON history USING btree (user_id);


--
-- TOC entry 2335 (class 1259 OID 16660)
-- Name: fki_privilege_entity_type_id_fkey; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX fki_privilege_entity_type_id_fkey ON privilege USING btree (subject_id);


--
-- TOC entry 2317 (class 1259 OID 16685)
-- Name: fki_role_permission__permission_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX fki_role_permission__permission_id ON role_permission USING btree (permission_id);


--
-- TOC entry 2318 (class 1259 OID 16686)
-- Name: fki_role_permission__role_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX fki_role_permission__role_id ON role_permission USING btree (role_id);


--
-- TOC entry 2321 (class 1259 OID 16485)
-- Name: fki_user_group__group_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX fki_user_group__group_id ON user_group USING btree (group_id);


--
-- TOC entry 2322 (class 1259 OID 16479)
-- Name: fki_user_group__user_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX fki_user_group__user_id ON user_group USING btree (user_id);


--
-- TOC entry 2344 (class 2606 OID 16663)
-- Name: fk_group_role__group_id; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY group_role
    ADD CONSTRAINT fk_group_role__group_id FOREIGN KEY (group_id) REFERENCES "group"(id);


--
-- TOC entry 2345 (class 2606 OID 16668)
-- Name: fk_group_role__role_id; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY group_role
    ADD CONSTRAINT fk_group_role__role_id FOREIGN KEY (role_id) REFERENCES role(id);


--
-- TOC entry 2347 (class 2606 OID 16688)
-- Name: fk_history__subject_id; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY history
    ADD CONSTRAINT fk_history__subject_id FOREIGN KEY (subject_id) REFERENCES subject(id);


--
-- TOC entry 2346 (class 2606 OID 16589)
-- Name: fk_history__user_id; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY history
    ADD CONSTRAINT fk_history__user_id FOREIGN KEY (user_id) REFERENCES "user"(id);


--
-- TOC entry 2348 (class 2606 OID 16655)
-- Name: fk_privilege__entity_type_id; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY privilege
    ADD CONSTRAINT fk_privilege__entity_type_id FOREIGN KEY (subject_id) REFERENCES subject(id);


--
-- TOC entry 2341 (class 2606 OID 16680)
-- Name: fk_role_permission__permission_id; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY role_permission
    ADD CONSTRAINT fk_role_permission__permission_id FOREIGN KEY (permission_id) REFERENCES permission(id);


--
-- TOC entry 2340 (class 2606 OID 16675)
-- Name: fk_role_permission__role_id; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY role_permission
    ADD CONSTRAINT fk_role_permission__role_id FOREIGN KEY (role_id) REFERENCES role(id);


--
-- TOC entry 2342 (class 2606 OID 16516)
-- Name: fk_user_group__group_id; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY user_group
    ADD CONSTRAINT fk_user_group__group_id FOREIGN KEY (group_id) REFERENCES "group"(id) ON DELETE CASCADE;


--
-- TOC entry 2343 (class 2606 OID 16521)
-- Name: fk_user_group__user_id; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY user_group
    ADD CONSTRAINT fk_user_group__user_id FOREIGN KEY (user_id) REFERENCES "user"(id) ON DELETE CASCADE;


--
-- TOC entry 2469 (class 0 OID 0)
-- Dependencies: 6
-- Name: public; Type: ACL; Schema: -; Owner: -
--

REVOKE ALL ON SCHEMA public FROM PUBLIC;
REVOKE ALL ON SCHEMA public FROM prithvi;
GRANT ALL ON SCHEMA public TO prithvi;
GRANT ALL ON SCHEMA public TO PUBLIC;


-- Completed on 2016-06-02 17:41:05 PDT

--
-- PostgreSQL database dump complete
--

