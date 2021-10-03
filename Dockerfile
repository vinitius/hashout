FROM golang:1.16-alpine

ARG ENV
ENV env=$ENV

WORKDIR /app

COPY . ./
RUN go mod download

WORKDIR /app/cmd

RUN go build -o /viniti.us/hashout
EXPOSE 8181

RUN apk --no-cache add curl
HEALTHCHECK --start-period=2s --interval=2s --timeout=5s \
    CMD curl --fail http://localhost:8181/ping || exit 1

CMD "/viniti.us/hashout"