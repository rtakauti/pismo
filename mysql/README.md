#Manual

>Instruções para testar

```bash
docker run -it --rm --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123 rtakauti/pismo_mysql
mysql -h 127.0.0.1 -u root -p
```