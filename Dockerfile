FROM golang:1.15.5

WORKDIR /usr/app
COPY . .

RUN make build

EXPOSE 8080
CMD ["./target/api"]
