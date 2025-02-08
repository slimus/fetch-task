# build environment
FROM golang:1.23-alpine3.21 AS build
WORKDIR /app
COPY . ./
RUN go build -o fetch-task cmd/main.go


# production environment
FROM alpine:3.21
COPY --from=build /app/fetch-task /app/fetch-task
CMD ["/app/fetch-task"]