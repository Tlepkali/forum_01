FROM golang:1.18.3 AS builder

LABEL maintainers="MAT TEAM"

WORKDIR /src

COPY . .
COPY config.json /src/config.json
COPY ui /src/ui
COPY initdb.sh /src/initdb.sh
COPY migrations /src/migrations

RUN go mod download

RUN GOOS=linux go build -o main.exe ./cmd/

FROM ubuntu:20.04

RUN apt-get update && apt-get install -y sqlite3

WORKDIR /src

COPY --from=builder /src/main.exe .
COPY --from=builder /src/config.json /src/config.json
COPY --from=builder /src/ui /src/ui
COPY --from=builder /src/Makefile /src/Makefile
COPY --from=builder /src/initdb.sh /src/initdb.sh
COPY --from=builder /src/migrations /src/migrations

EXPOSE 8080

CMD ["./main.exe"]