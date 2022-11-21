FROM golang:1.19 as build
WORKDIR /go/src/app
ADD go.mod go.sum ./
RUN go mod download
ADD . .
RUN CGO_ENABLED=0 go build -o /go/bin/gardensApp ./cmd/gardens

FROM gcr.io/distroless/static-debian11
# FROM alpine
COPY --from=build /go/bin/gardensApp /
COPY build/gardens/migrations /migrations
COPY build/gardens/default.yml .
ENTRYPOINT ["/gardensApp"]
CMD ["-config", "default.yml"]
