# FROM golang:1.15-alpine as builder

# WORKDIR /app
# COPY ./src .
# RUN GOOS=linux go build main.go

# FROM scratch
# WORKDIR /app
# COPY --from=builder /app .
# ENTRYPOINT [ "./app" ]

FROM golang:1.15-alpine
WORKDIR /app
ENV GO111MODULE=on
ENV CGO_ENABLED=1
COPY ./src /app
RUN GOOS=linux go build main.go
CMD ["./main"]