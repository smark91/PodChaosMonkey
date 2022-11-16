FROM golang:1.19.3 AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /pod-chaos-monkey

## Deploy
FROM gcr.io/distroless/base-debian11

WORKDIR /

COPY --from=build /pod-chaos-monkey /app/pod-chaos-monkey

USER nonroot:nonroot

ENTRYPOINT ["/app/pod-chaos-monkey"]
