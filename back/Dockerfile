# STEP 1 build executable binary
FROM alpine:latest as builder
RUN apk add go git
WORKDIR /bbs-task/
COPY ./ ./
#RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o "main" ./cmd/bbs-task

# STEP 2 build a small image
FROM scratch
COPY --from=builder "/bbs-task/main" "/"
COPY "./conf.json" "/"
EXPOSE 8080
ENTRYPOINT ["/main"]