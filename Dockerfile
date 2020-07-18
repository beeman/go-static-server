FROM golang:1.14.2-alpine as build

WORKDIR /workspace
ADD . /workspace

RUN go get -d -v ./...

RUN go build -o /workspace/app

FROM gcr.io/distroless/base-debian10
COPY --from=build /workspace/app /

CMD ["/app"]