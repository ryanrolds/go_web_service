# Build the binary
FROM golang:1.17-alpine AS build

RUN apk update && apk --no-cache --update add build-base

# make alternative and linter
RUN go install github.com/go-task/task/v3/cmd/task@latest
RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.44.2

WORKDIR /service
COPY . .

RUN task build
RUN task lint
RUN task test

# Create new image that contains just the backend binary
FROM alpine
WORKDIR /service
COPY --from=build /service/bin/backend .

CMD ["./backend"]