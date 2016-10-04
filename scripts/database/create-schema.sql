-- -- MODIFIED BY HAND TO WORK ON ubuntu precise with postgres 9.1

-- --
-- -- PostgreSQL database dump
-- --

-- \set ON_ERROR_STOP on

-- -- Dumped from database version 9.5.0
-- -- Dumped by pg_dump version 9.5.1

-- SET statement_timeout = 0;
-- SET client_encoding = 'UTF8';
-- SET standard_conforming_strings = on;
-- SET check_function_bodies = false;
-- SET client_min_messages = warning;

-- \connect steam

-- SET statement_timeout = 0;
-- SET client_encoding = 'UTF8';
-- SET standard_conforming_strings = on;
-- SET check_function_bodies = false;
-- SET client_min_messages = warning;

-- --
-- -- Name: public; Type: SCHEMA; Schema: -; Owner: steam
-- --

-- DROP SCHEMA IF EXISTS public;
-- CREATE SCHEMA public;


-- ALTER SCHEMA public OWNER TO steam;

-- --
-- -- Name: SCHEMA public; Type: COMMENT; Schema: -; Owner: steam
-- --

-- COMMENT ON SCHEMA public IS 'standard public schema';

-- SET search_path = public, pg_catalog;

-- --
-- -- Name: job_state; Type: TYPE; Schema: public; Owner: steam
-- --

-- CREATE TYPE job_state AS ENUM (
--     'idle',
--     'starting',
--     'started',
--     'suspending',
--     'suspended',
--     'stopping',
--     'stopped',
--     'blocked',
--     'disconnected',
--     'failed',
--     'completed'
-- );


-- ALTER TYPE job_state OWNER TO steam;

-- --
-- -- Name: privilege_type; Type: TYPE; Schema: public; Owner: steam
-- --

-- CREATE TYPE privilege_type AS ENUM (
--     'own',
--     'edit',
--     'view'
-- );


-- ALTER TYPE privilege_type OWNER TO steam;

-- --
-- -- Name: workgroup_type; Type: TYPE; Schema: public; Owner: steam
-- --

-- CREATE TYPE workgroup_type AS ENUM (
--     'identity',
--     'workgroup'
-- );


-- ALTER TYPE workgroup_type OWNER TO steam;

-- SET default_tablespace = '';

-- SET default_with_oids = false;

--
-- Name: binomial_model; Type: TABLE; Schema: public; Owner: steam
--

CREATE TABLE binomial_model (
    id integer PRIMARY KEY AUTOINCREMENT,
    model_id integer NOT NULL,
    mse double precision,
    r_squared double precision,
    logloss double precision,
    auc double precision,
    gini double precision, 

    FOREIGN KEY (model_id) REFERENCES model(id) ON DELETE CASCADE
);


-- ALTER TABLE binomial_model OWNER TO steam;

--
-- Name: binomial_model_id_seq; Type: SEQUENCE; Schema: public; Owner: steam
--

-- CREATE SEQUENCE binomial_model_id_seq
--     START WITH 1
--     INCREMENT BY 1
--     NO MINVALUE
--     NO MAXVALUE
--     CACHE 1;


-- ALTER TABLE binomial_model_id_seq OWNER TO steam;

--
-- Name: binomial_model_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: steam
--

-- ALTER SEQUENCE binomial_model_id_seq OWNED BY binomial_model.id;


--
-- Name: cluster; Type: TABLE; Schema: public; Owner: steam
--

CREATE TABLE cluster (
    id integer PRIMARY KEY AUTOINCREMENT,
    name text NOT NULL,
    type_id integer NOT NULL,
    detail_id integer NOT NULL,
    address text NOT NULL,
    state job_state NOT NULL,
    created datetime NOT NULL,

    FOREIGN KEY (type_id) REFERENCES cluster_type(id)
);


-- ALTER TABLE cluster OWNER TO steam;

--
-- Name: cluster_id_seq; Type: SEQUENCE; Schema: public; Owner: steam
--

-- CREATE SEQUENCE cluster_id_seq
--     START WITH 1
--     INCREMENT BY 1
--     NO MINVALUE
--     NO MAXVALUE
--     CACHE 1;


-- ALTER TABLE cluster_id_seq OWNER TO steam;

--
-- Name: cluster_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: steam
--

-- ALTER SEQUENCE cluster_id_seq OWNED BY cluster.id;


--
-- Name: cluster_type; Type: TABLE; Schema: public; Owner: steam
--

CREATE TABLE cluster_type (
    id integer PRIMARY KEY AUTOINCREMENT,
    name text NOT NULL UNIQUE 
);


-- ALTER TABLE cluster_type OWNER TO steam;

--
-- Name: cluster_type_id_seq; Type: SEQUENCE; Schema: public; Owner: steam
--

-- CREATE SEQUENCE cluster_type_id_seq
--     START WITH 1
--     INCREMENT BY 1
--     NO MINVALUE
--     NO MAXVALUE
--     CACHE 1;


-- ALTER TABLE cluster_type_id_seq OWNER TO steam;

--
-- Name: cluster_type_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: steam
--

-- ALTER SEQUENCE cluster_type_id_seq OWNED BY cluster_type.id;


--
-- Name: cluster_yarn; Type: TABLE; Schema: public; Owner: steam
--

CREATE TABLE cluster_yarn (
    id integer PRIMARY KEY AUTOINCREMENT,
    engine_id integer NOT NULL,
    size integer NOT NULL,
    application_id text NOT NULL,
    memory text NOT NULL,
    username text NOT NULL,
    output_dir text NOT NULL,

    FOREIGN KEY (engine_id) REFERENCES engine(id)
);


-- ALTER TABLE cluster_yarn OWNER TO steam;

--
-- Name: TABLE cluster_yarn; Type: COMMENT; Schema: public; Owner: steam
--

-- COMMENT ON TABLE cluster_yarn IS 'Launch parameters for YARN clusters.';


--
-- Name: cluster_yarn_id_seq; Type: SEQUENCE; Schema: public; Owner: steam
--

-- CREATE SEQUENCE cluster_yarn_id_seq
--     START WITH 1
--     INCREMENT BY 1
--     NO MINVALUE
--     NO MAXVALUE
--     CACHE 1;


-- ALTER TABLE cluster_yarn_id_seq OWNER TO steam;

--
-- Name: cluster_yarn_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: steam
--

-- ALTER SEQUENCE cluster_yarn_id_seq OWNED BY cluster_yarn.id;


--
-- Name: dataset; Type: TABLE; Schema: public; Owner: steam
--

CREATE TABLE dataset (
    id integer PRIMARY KEY AUTOINCREMENT,
    datasource_id integer NOT NULL,
    name text NOT NULL,
    description text NOT NULL,
    frame_name text NOT NULL,
    response_column_name text NOT NULL,
    properties text NOT NULL,
    properties_version text NOT NULL,
    created datetime NOT NULL,

    FOREIGN KEY (datasource_id) REFERENCES datasource(id) ON DELETE CASCADE
);


-- ALTER TABLE dataset OWNER TO steam;

--
-- Name: dataset_id_seq; Type: SEQUENCE; Schema: public; Owner: steam
--

-- CREATE SEQUENCE dataset_id_seq
--     START WITH 1
--     INCREMENT BY 1
--     NO MINVALUE
--     NO MAXVALUE
--     CACHE 1;


-- ALTER TABLE dataset_id_seq OWNER TO steam;

--
-- Name: dataset_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: steam
--

-- ALTER SEQUENCE dataset_id_seq OWNED BY dataset.id;


--
-- Name: datasource; Type: TABLE; Schema: public; Owner: steam
--

CREATE TABLE datasource (
    id integer PRIMARY KEY AUTOINCREMENT,
    project_id integer NOT NULL,
    name text NOT NULL,
    description text NOT NULL,
    kind text NOT NULL,
    configuration text NOT NULL,
    created datetime NOT NULL,

    FOREIGN KEY (project_id) REFERENCES project(id) ON DELETE CASCADE
);


-- ALTER TABLE datasource OWNER TO steam;

--
-- Name: datasource_id_seq; Type: SEQUENCE; Schema: public; Owner: steam
--

-- CREATE SEQUENCE datasource_id_seq
--     START WITH 1
--     INCREMENT BY 1
--     NO MINVALUE
--     NO MAXVALUE
--     CACHE 1;


-- ALTER TABLE datasource_id_seq OWNER TO steam;

--
-- Name: datasource_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: steam
--

-- ALTER SEQUENCE datasource_id_seq OWNED BY datasource.id;


--
-- Name: engine; Type: TABLE; Schema: public; Owner: steam
--

CREATE TABLE engine (
    id integer PRIMARY KEY AUTOINCREMENT,
    name text NOT NULL,
    location text NOT NULL,
    created datetime NOT NULL
);


-- ALTER TABLE engine OWNER TO steam;

--
-- Name: engine_id_seq; Type: SEQUENCE; Schema: public; Owner: steam
--

-- CREATE SEQUENCE engine_id_seq
--     START WITH 1
--     INCREMENT BY 1
--     NO MINVALUE
--     NO MAXVALUE
--     CACHE 1;


-- ALTER TABLE engine_id_seq OWNER TO steam;

--
-- Name: engine_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: steam
--

-- ALTER SEQUENCE engine_id_seq OWNED BY engine.id;


--
-- Name: entity_type; Type: TABLE; Schema: public; Owner: steam
--

CREATE TABLE entity_type (
    id integer PRIMARY KEY AUTOINCREMENT,
    name text NOT NULL UNIQUE
);


-- ALTER TABLE entity_type OWNER TO steam;

--
-- Name: entity_type_id_seq; Type: SEQUENCE; Schema: public; Owner: steam
--

-- CREATE SEQUENCE entity_type_id_seq
--     START WITH 1
--     INCREMENT BY 1
--     NO MINVALUE
--     NO MAXVALUE
--     CACHE 1;


-- ALTER TABLE entity_type_id_seq OWNER TO steam;

--
-- Name: entity_type_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: steam
--

-- ALTER SEQUENCE entity_type_id_seq OWNED BY entity_type.id;


--
-- Name: history; Type: TABLE; Schema: public; Owner: steam
--

CREATE TABLE history (
    id integer PRIMARY KEY AUTOINCREMENT,
    action text NOT NULL,
    identity_id integer NOT NULL,
    entity_type_id integer NOT NULL,
    entity_id integer NOT NULL,
    description text NOT NULL,
    created datetime NOT NULL,

    FOREIGN KEY (entity_type_id) REFERENCES entity_type(id),
    FOREIGN KEY (identity_id) REFERENCES identity(id)
);


-- ALTER TABLE history OWNER TO steam;

--
-- Name: history_id_seq; Type: SEQUENCE; Schema: public; Owner: steam
--

-- CREATE SEQUENCE history_id_seq
--     START WITH 1
--     INCREMENT BY 1
--     NO MINVALUE
--     NO MAXVALUE
--     CACHE 1;


-- ALTER TABLE history_id_seq OWNER TO steam;

--
-- Name: history_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: steam
--

-- ALTER SEQUENCE history_id_seq OWNED BY history.id;


--
-- Name: identity; Type: TABLE; Schema: public; Owner: steam
--

CREATE TABLE identity (
    id integer PRIMARY KEY AUTOINCREMENT,
    name text NOT NULL UNIQUE,
    password text NOT NULL,
    workgroup_id integer NOT NULL,
    is_active boolean NOT NULL,
    last_login integer with time zone,
    created datetime NOT NULL
);


-- ALTER TABLE identity OWNER TO steam;

--
-- Name: identity_id_seq; Type: SEQUENCE; Schema: public; Owner: steam
--

-- CREATE SEQUENCE identity_id_seq
--     START WITH 1
--     INCREMENT BY 1
--     NO MINVALUE
--     NO MAXVALUE
--     CACHE 1;


-- ALTER TABLE identity_id_seq OWNER TO steam;

--
-- Name: identity_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: steam
--

-- ALTER SEQUENCE identity_id_seq OWNED BY identity.id;


--
-- Name: identity_role; Type: TABLE; Schema: public; Owner: steam
--

CREATE TABLE identity_role (
    identity_id integer NOT NULL,
    role_id integer NOT NULL,

    PRIMARY KEY (identity_id, role_id)
);


-- ALTER TABLE identity_role OWNER TO steam;

--
-- Name: identity_workgroup; Type: TABLE; Schema: public; Owner: steam
--

CREATE TABLE identity_workgroup (
    identity_id integer NOT NULL,
    workgroup_id integer NOT NULL,

    PRIMARY KEY (identity_id, workgroup_id),
    FOREIGN KEY (identity_id) REFERENCES identity(id) ON DELETE CASCADE,
    FOREIGN KEY (workgroup_id) REFERENCES workgroup(id) ON DELETE CASCADE
);


-- ALTER TABLE identity_workgroup OWNER TO steam;

--
-- Name: label; Type: TABLE; Schema: public; Owner: steam
--

CREATE TABLE label (
    id integer PRIMARY KEY AUTOINCREMENT,
    project_id integer NOT NULL,
    model_id integer,
    name text NOT NULL,
    description text NOT NULL,
    created datetime NOT NULL,

    FOREIGN KEY (model_id) REFERENCES model(id) ON DELETE SET NULL,
    FOREIGN KEY (project_id) REFERENCES project(id) ON DELETE CASCADE
);


-- ALTER TABLE label OWNER TO steam;

--
-- Name: label_id_seq; Type: SEQUENCE; Schema: public; Owner: steam
--

-- CREATE SEQUENCE label_id_seq
--     START WITH 1
--     INCREMENT BY 1
--     NO MINVALUE
--     NO MAXVALUE
--     CACHE 1;


-- ALTER TABLE label_id_seq OWNER TO steam;

--
-- Name: label_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: steam
--

-- ALTER SEQUENCE label_id_seq OWNED BY label.id;


--
-- Name: meta; Type: TABLE; Schema: public; Owner: steam
--

CREATE TABLE meta (
    id integer NOT NULL,
    key text NOT NULL UNIQUE,
    value text NOT NULL,

    PRIMARY KEY (id)
);


-- ALTER TABLE meta OWNER TO steam;

--
-- Name: meta_id_seq; Type: SEQUENCE; Schema: public; Owner: steam
--

-- CREATE SEQUENCE meta_id_seq
--     START WITH 1
--     INCREMENT BY 1
--     NO MINVALUE
--     NO MAXVALUE
--     CACHE 1;


-- ALTER TABLE meta_id_seq OWNER TO steam;

--
-- Name: meta_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: steam
--

-- ALTER SEQUENCE meta_id_seq OWNED BY meta.id;


--
-- Name: model; Type: TABLE; Schema: public; Owner: steam
--

CREATE TABLE model (
    id integer PRIMARY KEY AUTOINCREMENT,
    project_id integer NOT NULL,
    training_dataset_id integer NOT NULL,
    validation_dataset_id integer NOT NULL,
    name text NOT NULL,
    cluster_name text NOT NULL,
    model_key text NOT NULL,
    algorithm text NOT NULL,
    model_category text NOT NULL,
    dataset_name text NOT NULL,
    response_column_name text NOT NULL,
    logical_name text NOT NULL,
    location text NOT NULL,
    max_run_time integer,
    metrics text NOT NULL,
    metrics_version text NOT NULL,
    created datetime NOT NULL,

    FOREIGN KEY (project_id) REFERENCES project(id),
    FOREIGN KEY (training_dataset_id) REFERENCES dataset(id),
    FOREIGN KEY (validation_dataset_id) REFERENCES dataset(id)
);


-- ALTER TABLE model OWNER TO steam;

--
-- Name: COLUMN model.name; Type: COMMENT; Schema: public; Owner: steam
--

-- COMMENT ON COLUMN model.name IS 'The physical name of this model as stored on disk.';


--
-- Name: COLUMN model.cluster_name; Type: COMMENT; Schema: public; Owner: steam
--

-- COMMENT ON COLUMN model.cluster_name IS 'The name of the cluster this model was sourced from.';


--
-- Name: COLUMN model.logical_name; Type: COMMENT; Schema: public; Owner: steam
--

-- COMMENT ON COLUMN model.logical_name IS 'The logical name of the model (typically the Java language class name).';


--
-- Name: COLUMN model.location; Type: COMMENT; Schema: public; Owner: steam
--

-- COMMENT ON COLUMN model.location IS 'The location of this model''s saved assets (e.g. /var/master/model).';


--
-- Name: COLUMN model.metrics; Type: COMMENT; Schema: public; Owner: steam
--

-- COMMENT ON COLUMN model.metrics IS 'Raw model metrics JSON obtained from H2O.';


--
-- Name: COLUMN model.metrics_version; Type: COMMENT; Schema: public; Owner: steam
--

-- COMMENT ON COLUMN model.metrics_version IS 'Version of the deserializer to use for unpacking metrics';


--
-- Name: model_id_seq; Type: SEQUENCE; Schema: public; Owner: steam
--

-- CREATE SEQUENCE model_id_seq
--     START WITH 1
--     INCREMENT BY 1
--     NO MINVALUE
--     NO MAXVALUE
--     CACHE 1;


-- ALTER TABLE model_id_seq OWNER TO steam;

--
-- Name: model_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: steam
--

-- ALTER SEQUENCE model_id_seq OWNED BY model.id;


--
-- Name: multinomial_model; Type: TABLE; Schema: public; Owner: steam
--

CREATE TABLE multinomial_model (
    id integer PRIMARY KEY AUTOINCREMENT,
    model_id integer NOT NULL,
    mse double precision,
    r_squared double precision,
    logloss double precision,

    FOREIGN KEY (model_id) REFERENCES model(id) ON DELETE CASCADE
);


-- ALTER TABLE multinomial_model OWNER TO steam;

--
-- Name: multinomial_model_id_seq; Type: SEQUENCE; Schema: public; Owner: steam
--

-- CREATE SEQUENCE multinomial_model_id_seq
--     START WITH 1
--     INCREMENT BY 1
--     NO MINVALUE
--     NO MAXVALUE
--     CACHE 1;


-- ALTER TABLE multinomial_model_id_seq OWNER TO steam;

--
-- Name: multinomial_model_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: steam
--

-- ALTER SEQUENCE multinomial_model_id_seq OWNED BY multinomial_model.id;


--
-- Name: permission; Type: TABLE; Schema: public; Owner: steam
--

CREATE TABLE permission (
    id integer PRIMARY KEY AUTOINCREMENT,
    code text NOT NULL UNIQUE,
    description text NOT NULL
);


-- ALTER TABLE permission OWNER TO steam;

--
-- Name: permission_id_seq; Type: SEQUENCE; Schema: public; Owner: steam
--

-- CREATE SEQUENCE permission_id_seq
--     START WITH 1
--     INCREMENT BY 1
--     NO MINVALUE
--     NO MAXVALUE
--     CACHE 1;


-- ALTER TABLE permission_id_seq OWNER TO steam;

--
-- Name: permission_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: steam
--

-- ALTER SEQUENCE permission_id_seq OWNED BY permission.id;


--
-- Name: privilege; Type: TABLE; Schema: public; Owner: steam
--

CREATE TABLE privilege (
    privilege_type text NOT NULL,
    workgroup_id integer NOT NULL,
    entity_type_id integer NOT NULL,
    entity_id integer NOT NULL,

    PRIMARY KEY (privilege_type, workgroup_id, entity_type_id, entity_id),
    FOREIGN KEY (entity_type_id) REFERENCES entity_type(id),
    FOREIGN KEY (workgroup_id) REFERENCES workgroup(id)
);


-- ALTER TABLE privilege OWNER TO steam;

--
-- Name: project; Type: TABLE; Schema: public; Owner: steam
--

CREATE TABLE project (
    id integer PRIMARY KEY AUTOINCREMENT,
    name text NOT NULL,
    description text NOT NULL,
    model_category text NOT NULL,
    created datetime NOT NULL  
);


-- ALTER TABLE project OWNER TO steam;

--
-- Name: project_id_seq; Type: SEQUENCE; Schema: public; Owner: steam
--

-- CREATE SEQUENCE project_id_seq
--     START WITH 1
--     INCREMENT BY 1
--     NO MINVALUE
--     NO MAXVALUE
--     CACHE 1;


-- ALTER TABLE project_id_seq OWNER TO steam;

--
-- Name: project_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: steam
--

-- ALTER SEQUENCE project_id_seq OWNED BY project.id;


--
-- Name: regression_model; Type: TABLE; Schema: public; Owner: steam
--

CREATE TABLE regression_model (
    id integer PRIMARY KEY AUTOINCREMENT,
    model_id integer NOT NULL,
    mse double precision,
    r_squared double precision,
    mean_residual_deviance double precision,

    FOREIGN KEY (model_id) REFERENCES model(id) ON DELETE CASCADE
);


-- ALTER TABLE regression_model OWNER TO steam;

--
-- Name: regression_model_id_seq; Type: SEQUENCE; Schema: public; Owner: steam
--

-- CREATE SEQUENCE regression_model_id_seq
--     START WITH 1
--     INCREMENT BY 1
--     NO MINVALUE
--     NO MAXVALUE
--     CACHE 1;


-- ALTER TABLE regression_model_id_seq OWNER TO steam;

--
-- Name: regression_model_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: steam
--

-- ALTER SEQUENCE regression_model_id_seq OWNED BY regression_model.id;


--
-- Name: role; Type: TABLE; Schema: public; Owner: steam
--

CREATE TABLE role (
    id integer PRIMARY KEY AUTOINCREMENT,
    name text NOT NULL UNIQUE,
    description text NOT NULL,
    created datetime NOT NULL  
);


-- ALTER TABLE role OWNER TO steam;

--
-- Name: role_id_seq; Type: SEQUENCE; Schema: public; Owner: steam
--

-- CREATE SEQUENCE role_id_seq
--     START WITH 1
--     INCREMENT BY 1
--     NO MINVALUE
--     NO MAXVALUE
--     CACHE 1;


-- ALTER TABLE role_id_seq OWNER TO steam;

--
-- Name: role_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: steam
--

-- ALTER SEQUENCE role_id_seq OWNED BY role.id;


--
-- Name: role_permission; Type: TABLE; Schema: public; Owner: steam
--

CREATE TABLE role_permission (
    role_id integer NOT NULL,
    permission_id integer NOT NULL,

    PRIMARY KEY (role_id, permission_id),
    FOREIGN KEY (permission_id) REFERENCES permission(id) ON DELETE CASCADE,
    FOREIGN KEY (role_id) REFERENCES role(id) ON DELETE CASCADE
);


-- ALTER TABLE role_permission OWNER TO steam;

--
-- Name: service; Type: TABLE; Schema: public; Owner: steam
--

CREATE TABLE service (
    id integer PRIMARY KEY AUTOINCREMENT,
    project_id integer NOT NULL,
    model_id integer NOT NULL,
    name text NOT NULL,
    address text NOT NULL,
    port integer NOT NULL,
    process_id integer NOT NULL,
    state job_state NOT NULL,
    created datetime NOT NULL,

    FOREIGN KEY (model_id) REFERENCES model(id)
);


-- ALTER TABLE service OWNER TO steam;

--
-- Name: service_id_seq; Type: SEQUENCE; Schema: public; Owner: steam
--

-- CREATE SEQUENCE service_id_seq
--     START WITH 1
--     INCREMENT BY 1
--     NO MINVALUE
--     NO MAXVALUE
--     CACHE 1;


-- ALTER TABLE service_id_seq OWNER TO steam;

--
-- Name: service_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: steam
--

-- ALTER SEQUENCE service_id_seq OWNED BY service.id;


--
-- Name: workgroup; Type: TABLE; Schema: public; Owner: steam
--

CREATE TABLE workgroup (
    id integer PRIMARY KEY AUTOINCREMENT,
    type workgroup_type NOT NULL,
    name text NOT NULL UNIQUE,
    description text NOT NULL,
    created datetime NOT NULL 
);


-- ALTER TABLE workgroup OWNER TO steam;

--
-- Name: workgroup_id_seq; Type: SEQUENCE; Schema: public; Owner: steam
--

-- CREATE SEQUENCE workgroup_id_seq
--     START WITH 1
--     INCREMENT BY 1
--     NO MINVALUE
--     NO MAXVALUE
--     CACHE 1;


-- ALTER TABLE workgroup_id_seq OWNER TO steam;

--
-- Name: workgroup_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: steam
--

-- ALTER SEQUENCE workgroup_id_seq OWNED BY workgroup.id;


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: steam
--

-- ALTER TABLE ONLY binomial_model ALTER COLUMN id SET DEFAULT nextval('binomial_model_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: steam
--

-- ALTER TABLE ONLY cluster ALTER COLUMN id SET DEFAULT nextval('cluster_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: steam
--

-- ALTER TABLE ONLY cluster_type ALTER COLUMN id SET DEFAULT nextval('cluster_type_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: steam
--

-- ALTER TABLE ONLY cluster_yarn ALTER COLUMN id SET DEFAULT nextval('cluster_yarn_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: steam
--

-- ALTER TABLE ONLY dataset ALTER COLUMN id SET DEFAULT nextval('dataset_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: steam
--

-- ALTER TABLE ONLY datasource ALTER COLUMN id SET DEFAULT nextval('datasource_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: steam
--

-- ALTER TABLE ONLY engine ALTER COLUMN id SET DEFAULT nextval('engine_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: steam
--

-- ALTER TABLE ONLY entity_type ALTER COLUMN id SET DEFAULT nextval('entity_type_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: steam
--

-- ALTER TABLE ONLY history ALTER COLUMN id SET DEFAULT nextval('history_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: steam
--

-- ALTER TABLE ONLY identity ALTER COLUMN id SET DEFAULT nextval('identity_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: steam
--

-- ALTER TABLE ONLY label ALTER COLUMN id SET DEFAULT nextval('label_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: steam
--

-- ALTER TABLE ONLY meta ALTER COLUMN id SET DEFAULT nextval('meta_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: steam
--

-- ALTER TABLE ONLY model ALTER COLUMN id SET DEFAULT nextval('model_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: steam
--

-- ALTER TABLE ONLY multinomial_model ALTER COLUMN id SET DEFAULT nextval('multinomial_model_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: steam
--

-- ALTER TABLE ONLY permission ALTER COLUMN id SET DEFAULT nextval('permission_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: steam
--

-- ALTER TABLE ONLY project ALTER COLUMN id SET DEFAULT nextval('project_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: steam
--

-- ALTER TABLE ONLY regression_model ALTER COLUMN id SET DEFAULT nextval('regression_model_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: steam
--

-- ALTER TABLE ONLY role ALTER COLUMN id SET DEFAULT nextval('role_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: steam
--

-- ALTER TABLE ONLY service ALTER COLUMN id SET DEFAULT nextval('service_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: steam
--

-- ALTER TABLE ONLY workgroup ALTER COLUMN id SET DEFAULT nextval('workgroup_id_seq'::regclass);


--
-- Name: pk_binomial_model; Type: CONSTRAINT; Schema: public; Owner: steam
--

-- ALTER TABLE ONLY binomial_model
--     ADD CONSTRAINT pk_binomial_model PRIMARY KEY (id);


--
-- Name: pk_cluster; Type: CONSTRAINT; Schema: public; Owner: steam
--

-- ALTER TABLE ONLY cluster
--     ADD CONSTRAINT pk_cluster PRIMARY KEY (id);


--
-- Name: pk_cluster_type; Type: CONSTRAINT; Schema: public; Owner: steam
--

-- ALTER TABLE ONLY cluster_type
--     ADD CONSTRAINT pk_cluster_type PRIMARY KEY (id);


-- --
-- -- Name: pk_cluster_yarn; Type: CONSTRAINT; Schema: public; Owner: steam
-- --

-- ALTER TABLE ONLY cluster_yarn
--     ADD CONSTRAINT pk_cluster_yarn PRIMARY KEY (id);


-- --
-- -- Name: pk_dataset; Type: CONSTRAINT; Schema: public; Owner: steam
-- --

-- ALTER TABLE ONLY dataset
--     ADD CONSTRAINT pk_dataset PRIMARY KEY (id);


-- --
-- -- Name: pk_datasource; Type: CONSTRAINT; Schema: public; Owner: steam
-- --

-- ALTER TABLE ONLY datasource
--     ADD CONSTRAINT pk_datasource PRIMARY KEY (id);


-- --
-- -- Name: pk_engine; Type: CONSTRAINT; Schema: public; Owner: steam
-- --

-- ALTER TABLE ONLY engine
--     ADD CONSTRAINT pk_engine PRIMARY KEY (id);


-- --
-- -- Name: pk_entity_type; Type: CONSTRAINT; Schema: public; Owner: steam
-- --

-- ALTER TABLE ONLY entity_type
--     ADD CONSTRAINT pk_entity_type PRIMARY KEY (id);


-- --
-- -- Name: pk_history; Type: CONSTRAINT; Schema: public; Owner: steam
-- --

-- ALTER TABLE ONLY history
--     ADD CONSTRAINT pk_history PRIMARY KEY (id);


-- --
-- -- Name: pk_identity; Type: CONSTRAINT; Schema: public; Owner: steam
-- --

-- ALTER TABLE ONLY identity
--     ADD CONSTRAINT pk_identity PRIMARY KEY (id);


-- --
-- -- Name: pk_identity_role; Type: CONSTRAINT; Schema: public; Owner: steam
-- --

-- ALTER TABLE ONLY identity_role
--     ADD CONSTRAINT pk_identity_role PRIMARY KEY (identity_id, role_id);


-- --
-- -- Name: pk_identity_workgroup; Type: CONSTRAINT; Schema: public; Owner: steam
-- --

-- ALTER TABLE ONLY identity_workgroup
--     ADD CONSTRAINT pk_identity_workgroup PRIMARY KEY (identity_id, workgroup_id);


-- --
-- -- Name: pk_label; Type: CONSTRAINT; Schema: public; Owner: steam
-- --

-- ALTER TABLE ONLY label
--     ADD CONSTRAINT pk_label PRIMARY KEY (id);


-- --
-- -- Name: pk_meta; Type: CONSTRAINT; Schema: public; Owner: steam
-- --

-- ALTER TABLE ONLY meta
--     ADD CONSTRAINT pk_meta PRIMARY KEY (id);


-- --
-- -- Name: pk_model; Type: CONSTRAINT; Schema: public; Owner: steam
-- --

-- ALTER TABLE ONLY model
--     ADD CONSTRAINT pk_model PRIMARY KEY (id);


-- --
-- -- Name: pk_multinomial_model; Type: CONSTRAINT; Schema: public; Owner: steam
-- --

-- ALTER TABLE ONLY multinomial_model
--     ADD CONSTRAINT pk_multinomial_model PRIMARY KEY (id);


-- --
-- -- Name: pk_permission; Type: CONSTRAINT; Schema: public; Owner: steam
-- --

-- ALTER TABLE ONLY permission
--     ADD CONSTRAINT pk_permission PRIMARY KEY (id);


-- --
-- -- Name: pk_privilege; Type: CONSTRAINT; Schema: public; Owner: steam
-- --

-- ALTER TABLE ONLY privilege
--     ADD CONSTRAINT pk_privilege PRIMARY KEY (privilege_type, workgroup_id, entity_type_id, entity_id);


-- --
-- -- Name: pk_project; Type: CONSTRAINT; Schema: public; Owner: steam
-- --

-- ALTER TABLE ONLY project
--     ADD CONSTRAINT pk_project PRIMARY KEY (id);


-- --
-- -- Name: pk_regression_model; Type: CONSTRAINT; Schema: public; Owner: steam
-- --

-- ALTER TABLE ONLY regression_model
--     ADD CONSTRAINT pk_regression_model PRIMARY KEY (id);


-- --
-- -- Name: pk_role; Type: CONSTRAINT; Schema: public; Owner: steam
-- --

-- ALTER TABLE ONLY role
--     ADD CONSTRAINT pk_role PRIMARY KEY (id);


-- --
-- -- Name: pk_role_permission; Type: CONSTRAINT; Schema: public; Owner: steam
-- --

-- ALTER TABLE ONLY role_permission
--     ADD CONSTRAINT pk_role_permission PRIMARY KEY (role_id, permission_id);


-- --
-- -- Name: pk_service; Type: CONSTRAINT; Schema: public; Owner: steam
-- --

-- ALTER TABLE ONLY service
--     ADD CONSTRAINT pk_service PRIMARY KEY (id);


-- --
-- -- Name: pk_workgroup; Type: CONSTRAINT; Schema: public; Owner: steam
-- --

-- ALTER TABLE ONLY workgroup
--     ADD CONSTRAINT pk_workgroup PRIMARY KEY (id);


--
-- Name: uq_cluster_type_name; Type: CONSTRAINT; Schema: public; Owner: steam
--

-- ALTER TABLE ONLY cluster_type
--     ADD CONSTRAINT uq_cluster_type_name UNIQUE (name);


-- --
-- -- Name: uq_entity_type_name; Type: CONSTRAINT; Schema: public; Owner: steam
-- --

-- ALTER TABLE ONLY entity_type
--     ADD CONSTRAINT uq_entity_type_name UNIQUE (name);


-- --
-- -- Name: uq_identity_name; Type: CONSTRAINT; Schema: public; Owner: steam
-- --

-- ALTER TABLE ONLY identity
--     ADD CONSTRAINT uq_identity_name UNIQUE (name);


-- --
-- -- Name: uq_meta_key; Type: CONSTRAINT; Schema: public; Owner: steam
-- --

-- ALTER TABLE ONLY meta
--     ADD CONSTRAINT uq_meta_key UNIQUE (key);


-- --
-- -- Name: uq_permission_name; Type: CONSTRAINT; Schema: public; Owner: steam
-- --

-- ALTER TABLE ONLY permission
--     ADD CONSTRAINT uq_permission_name UNIQUE (code);


-- --
-- -- Name: uq_role_name; Type: CONSTRAINT; Schema: public; Owner: steam
-- --

-- ALTER TABLE ONLY role
--     ADD CONSTRAINT uq_role_name UNIQUE (name);


-- --
-- -- Name: uq_workgroup_name; Type: CONSTRAINT; Schema: public; Owner: steam
-- --

-- ALTER TABLE ONLY workgroup
--     ADD CONSTRAINT uq_workgroup_name UNIQUE (name);


--
-- Name: fki_binomial_model__model_id; Type: INDEX; Schema: public; Owner: steam
--

CREATE INDEX fki_binomial_model__model_id ON binomial_model (model_id);


--
-- Name: fki_cluster__cluster_type_id; Type: INDEX; Schema: public; Owner: steam
--

CREATE INDEX fki_cluster__cluster_type_id ON cluster (type_id);


--
-- Name: fki_cluster_yarn__engine_id; Type: INDEX; Schema: public; Owner: steam
--

CREATE INDEX fki_cluster_yarn__engine_id ON cluster_yarn (engine_id);


--
-- Name: fki_dataset__datasource_id; Type: INDEX; Schema: public; Owner: steam
--

CREATE INDEX fki_dataset__datasource_id ON dataset (datasource_id);


--
-- Name: fki_datasource__project_id; Type: INDEX; Schema: public; Owner: steam
--

CREATE INDEX fki_datasource__project_id ON datasource (project_id);


--
-- Name: fki_history__entity_type_id; Type: INDEX; Schema: public; Owner: steam
--

CREATE INDEX fki_history__entity_type_id ON history (entity_type_id);


--
-- Name: fki_history__identity_id; Type: INDEX; Schema: public; Owner: steam
--

CREATE INDEX fki_history__identity_id ON history (identity_id);


--
-- Name: fki_identity_workgroup__identity_id; Type: INDEX; Schema: public; Owner: steam
--

CREATE INDEX fki_identity_workgroup__identity_id ON identity_workgroup (identity_id);


--
-- Name: fki_identity_workgroup__workgroup_id; Type: INDEX; Schema: public; Owner: steam
--

CREATE INDEX fki_identity_workgroup__workgroup_id ON identity_workgroup (workgroup_id);


--
-- Name: fki_label__model_id; Type: INDEX; Schema: public; Owner: steam
--

CREATE INDEX fki_label__model_id ON label (model_id);


--
-- Name: fki_label__project_id; Type: INDEX; Schema: public; Owner: steam
--

CREATE INDEX fki_label__project_id ON label (project_id);


--
-- Name: fki_model__project_id; Type: INDEX; Schema: public; Owner: steam
--

CREATE INDEX fki_model__project_id ON model (project_id);


--
-- Name: fki_model_id; Type: INDEX; Schema: public; Owner: steam
--

CREATE INDEX fki_model_id ON service (model_id);


--
-- Name: fki_model_training__dataset_id; Type: INDEX; Schema: public; Owner: steam
--

CREATE INDEX fki_model_training__dataset_id ON model (training_dataset_id);


--
-- Name: fki_model_validation__dataset_id; Type: INDEX; Schema: public; Owner: steam
--

CREATE INDEX fki_model_validation__dataset_id ON model (validation_dataset_id);


--
-- Name: fki_multinomial_model__model_id; Type: INDEX; Schema: public; Owner: steam
--

CREATE INDEX fki_multinomial_model__model_id ON multinomial_model (model_id);


--
-- Name: fki_privilege__entity_type_id; Type: INDEX; Schema: public; Owner: steam
--

CREATE INDEX fki_privilege__entity_type_id ON privilege (entity_type_id);


--
-- Name: fki_privilege__workgroup_id; Type: INDEX; Schema: public; Owner: steam
--

CREATE INDEX fki_privilege__workgroup_id ON privilege (workgroup_id);


--
-- Name: fki_regression_model__model_id; Type: INDEX; Schema: public; Owner: steam
--

CREATE INDEX fki_regression_model__model_id ON regression_model (model_id);


--
-- Name: fki_role_permission__permission_id; Type: INDEX; Schema: public; Owner: steam
--

CREATE INDEX fki_role_permission__permission_id ON role_permission (permission_id);


--
-- Name: fki_role_permission__role_id; Type: INDEX; Schema: public; Owner: steam
--

CREATE INDEX fki_role_permission__role_id ON role_permission (role_id);


--
-- Name: fki_workgroup_id; Type: INDEX; Schema: public; Owner: steam
--

CREATE INDEX fki_workgroup_id ON identity (workgroup_id);


-- --
-- -- Name: fk_binomial_model__model_id; Type: FK CONSTRAINT; Schema: public; Owner: steam
-- --

-- ALTER TABLE ONLY binomial_model
--     ADD CONSTRAINT fk_binomial_model__model_id FOREIGN KEY (model_id) REFERENCES model(id) ON DELETE CASCADE;


-- --
-- -- Name: fk_cluster__cluster_type_id; Type: FK CONSTRAINT; Schema: public; Owner: steam
-- --

-- ALTER TABLE ONLY cluster
--     ADD CONSTRAINT fk_cluster__cluster_type_id FOREIGN KEY (type_id) REFERENCES cluster_type(id);


-- --
-- -- Name: fk_cluster_yarn__engine_id; Type: FK CONSTRAINT; Schema: public; Owner: steam
-- --

-- ALTER TABLE ONLY cluster_yarn
--     ADD CONSTRAINT fk_cluster_yarn__engine_id FOREIGN KEY (engine_id) REFERENCES engine(id);


-- --
-- -- Name: fk_dataset__datasource_id; Type: FK CONSTRAINT; Schema: public; Owner: steam
-- --

-- ALTER TABLE ONLY dataset
--     ADD CONSTRAINT fk_dataset__datasource_id FOREIGN KEY (datasource_id) REFERENCES datasource(id) ON DELETE CASCADE;


-- --
-- -- Name: fk_datasource__project_id; Type: FK CONSTRAINT; Schema: public; Owner: steam
-- --

-- ALTER TABLE ONLY datasource
--     ADD CONSTRAINT fk_datasource__project_id FOREIGN KEY (project_id) REFERENCES project(id) ON DELETE CASCADE;


-- --
-- -- Name: fk_history__entity_type_id; Type: FK CONSTRAINT; Schema: public; Owner: steam
-- --

-- ALTER TABLE ONLY history
--     ADD CONSTRAINT fk_history__entity_type_id FOREIGN KEY (entity_type_id) REFERENCES entity_type(id);


-- --
-- -- Name: fk_history__identity_id; Type: FK CONSTRAINT; Schema: public; Owner: steam
-- --

-- ALTER TABLE ONLY history
--     ADD CONSTRAINT fk_history__identity_id FOREIGN KEY (identity_id) REFERENCES identity(id);


-- --
-- -- Name: fk_identity_workgroup__identity_id; Type: FK CONSTRAINT; Schema: public; Owner: steam
-- --

-- ALTER TABLE ONLY identity_workgroup
--     ADD CONSTRAINT fk_identity_workgroup__identity_id FOREIGN KEY (identity_id) REFERENCES identity(id) ON DELETE CASCADE;


-- --
-- -- Name: fk_identity_workgroup__workgroup_id; Type: FK CONSTRAINT; Schema: public; Owner: steam
-- --

-- ALTER TABLE ONLY identity_workgroup
--     ADD CONSTRAINT fk_identity_workgroup__workgroup_id FOREIGN KEY (workgroup_id) REFERENCES workgroup(id) ON DELETE CASCADE;


-- --
-- -- Name: fk_label__model_id; Type: FK CONSTRAINT; Schema: public; Owner: steam
-- --

-- ALTER TABLE ONLY label
--     ADD CONSTRAINT fk_label__model_id FOREIGN KEY (model_id) REFERENCES model(id);


-- --
-- -- Name: fk_label__project_id; Type: FK CONSTRAINT; Schema: public; Owner: steam
-- --

-- ALTER TABLE ONLY label
--     ADD CONSTRAINT fk_label__project_id FOREIGN KEY (project_id) REFERENCES project(id);


-- --
-- -- Name: fk_model__project_id; Type: FK CONSTRAINT; Schema: public; Owner: steam
-- --

-- ALTER TABLE ONLY model
--     ADD CONSTRAINT fk_model__project_id FOREIGN KEY (project_id) REFERENCES project(id);


-- --
-- -- Name: fk_model_training__dataset_id; Type: FK CONSTRAINT; Schema: public; Owner: steam
-- --

-- ALTER TABLE ONLY model
--     ADD CONSTRAINT fk_model_training__dataset_id FOREIGN KEY (training_dataset_id) REFERENCES dataset(id);


-- --
-- -- Name: fk_model_validation__dataset_id; Type: FK CONSTRAINT; Schema: public; Owner: steam
-- --

-- ALTER TABLE ONLY model
--     ADD CONSTRAINT fk_model_validation__dataset_id FOREIGN KEY (validation_dataset_id) REFERENCES dataset(id);


-- --
-- -- Name: fk_multinomial_model__model_id; Type: FK CONSTRAINT; Schema: public; Owner: steam
-- --

-- ALTER TABLE ONLY multinomial_model
--     ADD CONSTRAINT fk_multinomial_model__model_id FOREIGN KEY (model_id) REFERENCES model(id) ON DELETE CASCADE;


-- --
-- -- Name: fk_privilege__entity_type_id; Type: FK CONSTRAINT; Schema: public; Owner: steam
-- --

-- ALTER TABLE ONLY privilege
--     ADD CONSTRAINT fk_privilege__entity_type_id FOREIGN KEY (entity_type_id) REFERENCES entity_type(id);


-- --
-- -- Name: fk_privilege__workgroup_id; Type: FK CONSTRAINT; Schema: public; Owner: steam
-- --

-- ALTER TABLE ONLY privilege
--     ADD CONSTRAINT fk_privilege__workgroup_id FOREIGN KEY (workgroup_id) REFERENCES workgroup(id);


-- --
-- -- Name: fk_regression_model__model_id; Type: FK CONSTRAINT; Schema: public; Owner: steam
-- --

-- ALTER TABLE ONLY regression_model
--     ADD CONSTRAINT fk_regression_model__model_id FOREIGN KEY (model_id) REFERENCES model(id) ON DELETE CASCADE;


-- --
-- -- Name: fk_role_permission__permission_id; Type: FK CONSTRAINT; Schema: public; Owner: steam
-- --

-- ALTER TABLE ONLY role_permission
--     ADD CONSTRAINT fk_role_permission__permission_id FOREIGN KEY (permission_id) REFERENCES permission(id) ON DELETE CASCADE;


-- --
-- -- Name: fk_role_permission__role_id; Type: FK CONSTRAINT; Schema: public; Owner: steam
-- --

-- ALTER TABLE ONLY role_permission
--     ADD CONSTRAINT fk_role_permission__role_id FOREIGN KEY (role_id) REFERENCES role(id) ON DELETE CASCADE;


-- --
-- -- Name: fk_service__model_id; Type: FK CONSTRAINT; Schema: public; Owner: steam
-- --

-- ALTER TABLE ONLY service
--     ADD CONSTRAINT fk_service__model_id FOREIGN KEY (model_id) REFERENCES model(id);


-- --
-- -- Name: fk_workgroup_id; Type: FK CONSTRAINT; Schema: public; Owner: steam
-- --

-- ALTER TABLE ONLY identity
--     ADD CONSTRAINT fk_workgroup_id FOREIGN KEY (workgroup_id) REFERENCES workgroup(id);


--
-- Name: public; Type: ACL; Schema: -; Owner: steam
--

-- REVOKE ALL ON SCHEMA public FROM PUBLIC;
-- REVOKE ALL ON SCHEMA public FROM steam;
-- GRANT ALL ON SCHEMA public TO steam;
-- GRANT ALL ON SCHEMA public TO PUBLIC;


--
-- PostgreSQL database dump complete
--

