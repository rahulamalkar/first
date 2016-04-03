## PostgreSQL - database management

PostgreSQL is an object-relational database management system available for many platforms including Linux, FreeBSD, Solaris, Microsoft Windows and Mac OS X. It is released under an MIT-style license, and is thus free and open source software. It's larger than MySQL because it's designed for enterprise usage like Oracle. Postgresql is good choice for enterprise type projects.


## Links

https://godoc.org/github.com/lib/pq

https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/05.4.html

## Code Review

[Create database connectivity documents](example1/example1.go)

## Database

### PostgreSQL
#### qwery to create Database

```
create database mit_workshop
```

#### qwery to enter Database

```
psql -U postgres mit_workshop
```

#### create table users

```
CREATE TABLE users (
id SERIAL,
firstname varchar(100),
lastname varchar(100),
email varchar(320),
password varchar(100),
password_confirmation varchar(100),
PRIMARY KEY(id));
```
