FROM postgres:11

ENV POSTGRES_USER postgres
ENV POSTGRES_PASSWORD password
ENV POSTGRES_DB goskeleton

COPY . /docker-entrypoint-initdb.d/

