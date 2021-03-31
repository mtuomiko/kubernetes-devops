FROM golang:alpine as build
WORKDIR /work
COPY go.mod .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build

FROM scratch
WORKDIR /app
COPY --from=build /work/todo-generator .
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENTRYPOINT ["/app/todo-generator"]