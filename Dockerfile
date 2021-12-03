FROM golang:1.16-buster AS build

WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY ./ ./
RUN cd example && go build -o /example

FROM gcr.io/distroless/base-debian10

WORKDIR /
COPY --from=build /example /example
ENTRYPOINT ["/example"]