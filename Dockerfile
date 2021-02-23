FROM golang:1.16.0 as build
WORKDIR /build/src
COPY . /build/src/
RUN CGO_ENABLED=0 go build -o /build/bin/server main.go

FROM scratch as dist
COPY --from=build /build/bin/server /server
EXPOSE 8090
ENTRYPOINT ["/server"]
