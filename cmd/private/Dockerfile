FROM golang:1.13 as build
WORKDIR /app
COPY ./ .
RUN go test ./...
RUN go install back-api/cmd/private
RUN chmod +x ./wait-for-it.sh
EXPOSE 8081
ENTRYPOINT ["./wait-for-it.sh", "mongo:27017", "--", "private", "--addr=:8081", "--import=./data/import.csv"]