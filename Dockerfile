FROM golang:1.20

WORKDIR /api

COPY .. .

RUN make build

CMD ["./bin/vehicle"]