FROM golang:alpine3.18
WORKDIR /app
COPY . .
RUN go build
EXPOSE 3000
CMD ["./test-data-api"]