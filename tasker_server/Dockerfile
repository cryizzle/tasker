FROM golang:1.23.4-alpine AS build

WORKDIR /app

COPY go.* .

RUN go mod tidy
RUN go mod verify
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /tasker_server

FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build /tasker_server /tasker_server

USER nonroot:nonroot

ENTRYPOINT ["/tasker_server"]
