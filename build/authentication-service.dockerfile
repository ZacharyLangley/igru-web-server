FROM golang:1.19 as build
WORKDIR /go/src/app
ADD go.mod go.sum ./
RUN go mod download
ADD . .
RUN CGO_ENABLED=0 go build -o /go/bin/authApp ./cmd/auth

FROM gcr.io/distroless/static-debian11
# FROM alpine
COPY --from=build /go/bin/authApp /
COPY build/authentication/migrations /migrations
COPY build/authentication/default.yml .
ENTRYPOINT ["/authApp"]
CMD ["-config", "default.yml"]
