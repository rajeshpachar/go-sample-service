FROM golang:alpine as builder
RUN apk update && apk add git && apk add ca-certificates

RUN mkdir /build

ADD . /build/
WORKDIR /build
RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o healthcheck "github.com/rajeshpachar/go-health-check" 

FROM scratch
COPY --from=builder /build/main /app/
COPY --from=builder /build/healthcheck /app/
ENV PORT=3333
EXPOSE $PORT

WORKDIR /app

HEALTHCHECK --interval=5s --timeout=5s --start-period=2s --retries=3 CMD [ "./healthcheck" ]

# COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
CMD ["./main"]

