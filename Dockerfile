## Builder
FROM golang:latest as builder
WORKDIR /go/src/github.com/go-project
COPY . .
RUN go get .
RUN go build -o main .

## Start from the latest golang base image
FROM golang:latest
WORKDIR /app
ARG LOG_DIR=/app/logs
RUN mkdir -p ${LOG_DIR}
ENV LOG_FILE_LOCATION=${LOG_DIR}/app.log

EXPOSE 8083

# Add from source to /app
RUN mkdir /app/config
COPY --from=builder /go/src/github.com/go-project/config/config.json /app/config/
COPY --from=builder /go/src/github.com/go-project/core/ /app/core/
COPY --from=builder /go/src/github.com/go-project/main /app

# Declare volumes to mount
VOLUME [${LOG_DIR}]

# Run the binary program produced by `go install`
CMD /app/main