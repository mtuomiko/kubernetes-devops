FROM golang:alpine as build
WORKDIR /work
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build

FROM scratch
WORKDIR /app
COPY --from=build /work/dogcat-backend .
COPY --from=build /work/assets /app/assets
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENTRYPOINT ["/app/dogcat-backend"]