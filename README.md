# Expenser API Service

### How to setup database?
This app uses Postgre 12.3 to store data. The boostrap scripts are in `db/bootstrap`. The Dockerfile creates a database
and users. 
* Go to `db/bootstrap` dir and run `docker build .`
* Run the container using the image created using

` docker run -it -d -e POSTGRES_PASSWORD=<admin password> -e EXPUSER_PW=<user password> <image id>`

* Then use `db/migrate/0001_create_expense_table.sql` to create tables.
