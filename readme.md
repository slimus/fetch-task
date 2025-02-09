# Improvements
1. When we work with money, we should always think about different currencies and conversions between different currencies. Maybe we should use [special package](https://github.com/Rhymond/go-money).
1. Replace standard http server with [gin](https://github.com/gin-gonic/gin) \ [Gorilla mux](https://github.com/gorilla/mux) \ etc.
1. Replace standard `log` package with [better logging library](https://github.com/sirupsen/logrus).
1. Cover API functions with unit. I covered only most important part.
1. Add API documentation (Swagger UI, etc).
1. It's a good practice to use .env and ENV together. We could use [godotenv](https://github.com/joho/godotenv) to achieve this.


# How to run and test API
1. `docker compose up -d` - this command builds container with API and run it using port 8383 (can be changed in compose.yaml)
1. open `api.http` file and send requests from there with rest client (I used VSCode rest client)


# How to run unit tests
1. `go test ./...`