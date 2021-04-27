FROM golang:alpine as build

WORKDIR /work
COPY go.mod .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build

FROM scratch

WORKDIR /app
COPY --from=build /work/ticker-server /app/ticker-server
ENTRYPOINT ["/app/ticker-server"]