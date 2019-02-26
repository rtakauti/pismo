FROM mysql
MAINTAINER Rubens Takauti<rtakauti@hotmail.com>

ENV MYSQL_ROOT_PASSWORD=$PASSWORD
COPY ./mysql/pismo.sql /docker-entrypoint-initdb.d
EXPOSE 3306