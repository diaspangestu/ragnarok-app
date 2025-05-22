FROM --platform=linux/amd64 golang:1.22-alpine AS build

WORKDIR /app
COPY . .

RUN go mod download
#RUN go build -o /app/ragnarok

RUN GOOS=linux GOARCH=amd64 go build -o /app/ragnarok

FROM --platform=linux/amd64 alpine:latest

RUN apk add --no-cache tzdata

# Set the timezone
ENV TZ Asia/Jakarta

WORKDIR /app
COPY --from=build /app/ragnarok .

RUN chmod +x ragnarok

CMD ["./ragnarok"]