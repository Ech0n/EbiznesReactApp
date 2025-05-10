FROM golang:1.20-alpine

RUN addgroup -S nonroot \
    && adduser -S nonroot -G nonroot

ENV CGO_ENABLED=1

WORKDIR /app

RUN apk update && apk add --no-cache gcc musl-dev sqlite

COPY ./app /app

RUN go mod tidy

EXPOSE 8080

RUN chown -R nonroot:nonroot /go/pkg /app

USER nonroot

CMD ["go", "run", "main.go"]