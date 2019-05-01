FROM golang:1.11 as builder

WORKDIR /box
COPY go.sum .
COPY go.mod .
RUN go mod download

COPY main.go .
COPY controllers ./controllers
COPY core ./core
COPY routers ./routers

RUN CGO_ENABLED="0" go build

FROM scratch

COPY --from=builder /box/quote .
COPY conf conf

EXPOSE 8099

ENTRYPOINT [ "./quote" ]