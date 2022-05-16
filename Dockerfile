FROM golang:1.17-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY . ./

RUN go test ./...
RUN go build -o /mars-exploration

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /mars-exploration /mars-exploration

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/mars-exploration"]