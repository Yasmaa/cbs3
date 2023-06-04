FROM golang:alpine

LABEL maintainer="asma faouzi"

RUN apk update && apk add --no-cache git && apk add --no-cach bash && apk add build-base

RUN mkdir /app
WORKDIR /app

COPY . .
COPY .env .

RUN go mod download

RUN go build -o /build

EXPOSE 8080

# Run the executable
CMD [ "/build" ]