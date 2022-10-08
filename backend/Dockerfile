FROM golang

COPY . .
EXPOSE 8080

ENV GOPATH=/

RUN go build cmd/app/main.go

CMD ["./main"]