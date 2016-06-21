#!/usr/bin/env bash

psql -U $USER postgres -a -f create-database.sql
