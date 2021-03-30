FROM golang:1.16 as builder
WORKDIR /app

# copy dependencies info
COPY go.mod ./
# install dependencies
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main /app/cmd/mutant-checker/main.go

# # now move to scratch
FROM scratch
WORKDIR /app

# copy the binary from builder
COPY --from=builder /app .
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# run the binary
CMD ["./main"]