FROM golang:1.20.5

RUN go version
ENV GOPATH=/

COPY ./ ./


RUN go mod download
RUN go build -o LinkShortner ./cmd/main.go

CMD ./LinkShortner