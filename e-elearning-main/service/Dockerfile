FROM golang:1.22-alpine

WORKDIR /app

COPY . .
COPY cmd/merge-blob cmd/merge-blob

# RUN apk update && apk add --no-cache ffmpeg

RUN go mod tidy
RUN go build -o app cmd/merge-blob/main.go

CMD ["./app"]
