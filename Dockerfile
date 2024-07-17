# Build stage
FROM golang:1.21-alpine3.19 AS builder

WORKDIR /app
COPY . .
RUN go build -o main main.go

# Run stage
FROM alpine:3.19
WORKDIR /app
# The directory where all application data will reside
# Our Volume will be mounted at this directory inside the container.
# This means all data being written inside this directory will actually be
# written to the external volume, so it is preserved even after the container
# exits.
RUN mkdir /data
COPY --from=builder /app/main .

EXPOSE 8080
CMD [ "/app/main" ]

