FROM postgres:12.3
COPY ./bootstrap.sh /docker-entrypoint-initdb.d

ENV POSTGRES_USER admin
ENV POSTGRES_DB expense_db

EXPOSE 5432
CMD ["postgres"]
