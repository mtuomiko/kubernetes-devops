FROM golang:alpine as build

WORKDIR /work
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build

FROM scratch

WORKDIR /app
COPY --from=build /work/ticker-generator /app/ticker-generator
ENTRYPOINT ["/app/ticker-generator"]