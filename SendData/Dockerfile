FROM golang:1.19.4-alpine as build
WORKDIR /build
COPY kafka-recorder/ /build/
RUN go mod tidy
RUN go get
RUN go build -o kafka-recorder

FROM alpine
COPY --from=build /build/kafka-recorder /app/

ENTRYPOINT ["/app/kafka-recorder", "replay"]
