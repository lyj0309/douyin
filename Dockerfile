FROM golang as builder

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
# CGO_ENABLED alpine禁用cgo

WORKDIR /app
ADD go.mod .
ADD go.sum .
RUN go mod download


COPY . .
#RUN GOOS=linux GOARCH=amd64 go build -o app ./${type}
RUN go build -o app ./

RUN mkdir publish && cp app publish

FROM alpine
RUN apk add gcompat
WORKDIR /app
COPY --from=builder /app/publish .
ENV GIN_MODE=release \
    PORT=8081


EXPOSE 8081
ENTRYPOINT ["./app"]