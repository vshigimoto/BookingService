### Golang Booking microservices app, with Prometheus, Grafana monitoring, Kafka and Swagger

#### List of what i used in project:
* [gin](https://github.com/gin-gonic/gin) - Gin Web Framework
* [sql](https://pkg.go.dev/database/sql) - Extensions to database/sql.
* [pgx](https://github.com/jackc/pgx) - PostgreSQL driver and toolkit for Go
* [viper](https://github.com/spf13/viper) - Go configuration with fangs
* [go-redis](https://github.com/go-redis/redis) - Redis client for Golang
* [goose](https://github.com/pressly/goose) - Database migrations
* [Docker](https://www.docker.com/) - Docker
* [Prometheus](https://prometheus.io/) - Prometheus
* [Grafana](https://grafana.com/) - Grafana
* [Gin-swagger](https://github.com/swaggo/gin-swagger) - swagger for gin
* [Kafka](https://kafka.apache.org/) - kafka message broker

#### Recommendation for local development most comfortable usage:
    make local // run all containers
    make migrate_up // migrate data to DB
    make runUser // run the User service
    make runAuth // run Auth service
    make runBooking // run main Booking service

#### Migrations commands:
    make migrate_up - migrate up data to DB
    make migrate_down - migrate down data from DB

### Kafka UI:

http://localhost:9000

### Prometheus UI:

http://localhost:9090

### Grafana UI:

http://localhost:3000

### Swagger UI:

http://localhost:9234/docs/index.html
