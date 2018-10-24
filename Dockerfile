FROM golang:latest
# as builder

WORKDIR /app

COPY main.go .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o webserver


#FROM scratch

#COPY --from=builder /app/webserver /webserver

CMD ["/app/webserver"]
