FROM golang:1.17-alpine

WORKDIR /src/
COPY main.go go.* /src/
RUN go mod download
RUN CGO_ENABLED=0 go build -o /christmas-tree

CMD ["/christmas-tree"]
