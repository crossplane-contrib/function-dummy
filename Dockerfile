FROM golang:1.20 as build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY input/ ./input
COPY *.go ./

RUN CGO_ENABLED=0 go build -o /function-dummy .

FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /function-dummy /function-dummy

EXPOSE 9443

USER nonroot:nonroot

ENTRYPOINT ["/function-dummy"]