#!/usr/bin/env bash

psql -U $USER postgres -a -f drop-database.sql

