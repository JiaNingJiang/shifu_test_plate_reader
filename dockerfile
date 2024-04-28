# syntax=docker/dockerfile:1

FROM golang:1.17-alpine
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY *.go ./
RUN go build -o /plate-detector
EXPOSE 11111
CMD [ "/plate-detector" ] 