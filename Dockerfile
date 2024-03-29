FROM golang:alpine as build-env
COPY . /src
WORKDIR /src
RUN go build -o gowon-tenor

FROM alpine:3.15.0
WORKDIR /app
COPY --from=build-env /src/gowon-tenor /app/
ENTRYPOINT ["./gowon-tenor"]
