--
-- Drops the Steam database (for administration/development use only)
--

\set ON_ERROR_STOP on

SET statement_timeout = 0;
SET lock_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;

DROP DATABASE IF EXISTS steam;

