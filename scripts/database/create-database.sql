-- HAND CREATED TO WORK ON aws postgres


DROP DATABASE IF EXISTS steam;

--
-- Name: steam; Type: DATABASE; Schema: -; Owner: steam
--

CREATE DATABASE steam WITH TEMPLATE = template0 ENCODING = 'UTF8' LC_COLLATE = 'en_US.UTF-8' LC_CTYPE = 'en_US.UTF-8';

ALTER DATABASE steam OWNER TO steam;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

DROP EXTENSION IF EXISTS plpgsql;
CREATE EXTENSION plpgsql WITH SCHEMA pg_catalog;

--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';
