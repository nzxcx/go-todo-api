FROM golang:1.23-alpine

WORKDIR /app

# Install required tools
RUN apk add --no-cache netcat-openbsd
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main .

# Make the init script executable
COPY scripts/init.sh /init.sh
RUN chmod +x /init.sh

EXPOSE 8080

ENTRYPOINT ["/init.sh"]
CMD ["./main"] 
