FROM postgres:16-alpine

COPY /secrets/postgres-pwd /run/secrets/postgres-pwd
COPY init.sql /docker-entrypoint-initdb.d/

EXPOSE 5432