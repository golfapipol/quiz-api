FROM golang:1.11 as builder
WORKDIR /module
COPY ./quiz/ /module
RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/app

FROM alpine
WORKDIR /root/
COPY --from=builder /module/bin .
EXPOSE 3000

ENV GIN_MODE release
CMD ["./app"]
