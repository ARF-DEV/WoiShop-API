FROM golang:1.20.1-alpine AS build

WORKDIR /app

COPY . .
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /woishop

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /woishop /woishop

EXPOSE 8000

USER nonroot:nonroot

ENTRYPOINT ["/woishop"]