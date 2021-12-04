FROM golang:1.15

WORKDIR /backend

COPY .. .

RUN make build

CMD ["./bin/digihey"]
