# --- Base ----
FROM golang:1.20.0-buster AS base
WORKDIR $GOPATH/teste_go

# ---- Dependencies ----
FROM base AS dependencies
ENV GO111MODULE=on
COPY go.mod go.sum ./
RUN go mod download

# ---- Build ----
FROM dependencies AS build
COPY . .
RUN go build -o main .
EXPOSE 8080
CMD ["./main"]