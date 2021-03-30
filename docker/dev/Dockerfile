FROM golang:1.16
WORKDIR /app

# install hot reload
RUN go get github.com/codegangsta/gin

# copy dependencies info
COPY go.mod ./
# install dependencies
RUN go mod download

COPY . .

CMD ["gin", "-p", "4001", "-a", "4000", "-i", "-d", "cmd/mutant-checker/", "run", "main.go"]