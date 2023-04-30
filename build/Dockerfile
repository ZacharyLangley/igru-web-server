FROM node:lts as js-builder
WORKDIR /web
COPY web/package*.json web/tsconfig*.json ./
ENV PATH /web/node_modules/.bin:$PATH
ADD web/public/ public/
ADD web/src/ src/
RUN npm install
RUN npm run build

FROM golang:1.20 as build
WORKDIR /go/src/app
ADD go.mod go.sum ./
RUN go mod download
ADD . .
COPY --from=js-builder /web/build/ /go/src/app/cmd/ingress/public/
RUN CGO_ENABLED=0 go build -o /go/bin/app .

FROM gcr.io/distroless/static-debian11
COPY --from=build /go/bin/app /usr/bin/igru
ENTRYPOINT ["igru"]
CMD ["ingress", "serve", "--config", "default.yml"]
