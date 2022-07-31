FROM golang:1.18.4-alpine

WORKDIR /app

# download the required Go dependencies
COPY go.mod ./
COPY go.sum ./
RUN go mod download
#COPY *.go ./
COPY . ./

RUN go build -o /go-crud

EXPOSE 3000

CMD ["/go-crud"]

# docker build -t go-crud .

# docker run -p 8080:8080 -d go-crud