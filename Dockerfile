FROM golang:1.13-alpine as builder
WORKDIR /app
COPY go.mod .
COPY go.sum .

#dependencies
RUN go mod download
COPY . .

RUN go build -o ./dist/main

FROM alpine:3.5
RUN apk add --update ca-certificates
RUN apk add --no-cache tzdata && \
  cp -f /usr/share/zoneinfo/Asia/Ho_Chi_Minh /etc/localtime && \
  apk del tzdata

WORKDIR /app
COPY ./config/config.yaml ./config/
COPY --from=builder /app/dist/main .
ENV MONGO__HOST mongo-docker
EXPOSE 9090
ENTRYPOINT ["./main"]
