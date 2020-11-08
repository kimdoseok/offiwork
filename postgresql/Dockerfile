FROM postgres

LABEL MAINTAINER Doseok Kim <doseok@madang.com>

ENV POSTGRES_PASSWORD k!md0$25K

COPY init-user-db.sh /docker-entrypoint-initdb.d/

EXPOSE  5432