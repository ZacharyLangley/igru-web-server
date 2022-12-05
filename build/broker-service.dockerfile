FROM golang:1.19 as build
WORKDIR /go/src/app
ADD go.mod go.sum ./
RUN go mod download
ADD . .
RUN CGO_ENABLED=0 go build -o /go/bin/app .

FROM gcr.io/distroless/static-debian11
COPY --from=build /go/bin/app /
COPY build/broker/default.yml .
ENTRYPOINT ["/app", "broker", "serve"]
CMD ["--config", "default.yml"]
