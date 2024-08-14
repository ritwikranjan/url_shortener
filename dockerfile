FROM golang:1.22.5

WORKDIR /app

COPY . .

RUN go build -o bin .

ENTRYPOINT [ "/app/bin" ]

EXPOSE 8080
