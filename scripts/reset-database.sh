#!/usr/bin/env bash

psql -U $USER postgres -a -f database/drop-database.sql
psql -U $USER postgres -a -f database/create-database.sql

