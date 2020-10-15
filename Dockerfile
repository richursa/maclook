FROM golang:1.15.3 as builder
WORKDIR /src
ENV CGO_ENABLED=0
COPY . .
RUN apt install git -y
RUN go get "github.com/tidwall/gjson"
RUN go build main.go



FROM alpine:latest
COPY --from=builder /src/main .
ENTRYPOINT [ "./main" ]
