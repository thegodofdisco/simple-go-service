FROM golang:1.15.6-alpine3.12
WORKDIR /app
COPY . /app
EXPOSE 8080
ENTRYPOINT [ "go", "run" ]
CMD [ "/app/cmd/simple-go-service/main.go" ]