FROM postgres:alpine

COPY /pkg/db/scripts/initdb.sql /docker-entrypoint-initdb.d/initdb.sql

EXPOSE 5432