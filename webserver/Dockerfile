FROM golang:1.16 AS builder
WORKDIR /go/src/
COPY . .
RUN GOOS=linux go build main.go
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
# CMD [ "./main" ]

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/main .
RUN ls
CMD ["./main"]  