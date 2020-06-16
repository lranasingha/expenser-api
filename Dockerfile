FROM golang:1.14.4-buster AS builder
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o expense-api .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /build/expense-api /app/
WORKDIR /app/

CMD ["./expense-api"]
EXPOSE 8000