FROM golang:1.21

WORKDIR /appdir

COPY go.mod .
COPY go.sum .
COPY main.go .
COPY lists ./lists

RUN go build -o bin .

ENTRYPOINT [ "/appdir/bin" ]