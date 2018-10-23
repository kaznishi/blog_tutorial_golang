FROM mysql:5.6
MAINTAINER kaznishi<kaz.nishimu@gmail.com>

COPY docker/mysql/my.cnf /etc/my.cnf

COPY docker/mysql/initdb.d/*.sql /docker-entrypoint-initdb.d/

EXPOSE 3306