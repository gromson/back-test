FROM golang:1.13 as build
WORKDIR /app
COPY ./ .
RUN go test ./...
RUN go install back-api/cmd/public
RUN chmod +x ./wait-for-it.sh
EXPOSE 8080
ENTRYPOINT ["./wait-for-it.sh", "mongo:27017", "--", "public", "--addr=:8080"]