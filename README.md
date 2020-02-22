# Barmen

* Copy **.env-example** and write your params
* Run containers
```
docker-compose up
```
* Run migrations
```
goose postgres "database connection string" up

//database connection string example: "host=host port=port user=user password=password dbname=dbname sslmode=sslmode"

//other goose commands:
goose postgres "database connection string" down
goose postgres "database connection string" create new_miration sql
```

You can use **BloomRPC** for checking gRPC API