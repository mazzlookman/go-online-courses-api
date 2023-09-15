FROM golang:1.21-alpine as stage_1
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./
RUN go build -o ./pzn-api main.go

FROM alpine:latest
RUN apk update
RUN apk upgrade
RUN apk add --no-cache ffmpeg

WORKDIR /app
COPY --from=stage_1 /app/pzn-api ./
COPY --from=stage_1 /app/.env ./

EXPOSE 2802

CMD ["./pzn-api"]


