FROM golang:1.10-alpine AS build
WORKDIR /go/src/github.com/brancz/printserver
COPY . .
RUN go install

FROM alpine:3.7
COPY --from=build /go/bin/printserver .
ENTRYPOINT ["./printserver"]
EXPOSE 8080
