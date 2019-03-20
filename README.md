# Go DB Testing

## Running

1- Get a running instance of MySQL

`docker run --name testmysql -e MYSQL_ROOT_PASSWORD=rootpass -e MYSQL_DATABASE=blog -p 3306:3306 -d mysql:latest`

2- Run the migration files sequentially

`cat sql-files/* | mysql -u root -prootpass -h 0.0.0.0 --port=3306`
