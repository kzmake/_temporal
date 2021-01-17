FROM golang:1.15 as builder

WORKDIR /go/src/github.com/kzmake/_temporal/booking_saga

ENV GO111MODULE=on

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/trip-api github.com/kzmake/_temporal/booking_saga/cmd/trip-api
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/hotel-api github.com/kzmake/_temporal/booking_saga/cmd/hotel-api
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/car-api github.com/kzmake/_temporal/booking_saga/cmd/car-api
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/flight-api github.com/kzmake/_temporal/booking_saga/cmd/flight-api
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/worker github.com/kzmake/_temporal/booking_saga/cmd/worker

FROM alpine

COPY --from=builder /go/bin /usr/local/bin

EXPOSE 3000

CMD ["worker"]
