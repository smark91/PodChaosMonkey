FROM golang:1.19.3 AS base

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

FROM base as test 
RUN go test -v

FROM base as build 
RUN go build -o /pod-chaos-monkey

## Deploy
FROM gcr.io/distroless/base-debian11

WORKDIR /

COPY --from=build /pod-chaos-monkey /app/pod-chaos-monkey

USER nonroot:nonroot

ENTRYPOINT ["/app/pod-chaos-monkey"]
