FROM registry.access.redhat.com/ubi8/go-toolset:1.17.7

# Setup go environment variables
ENV GOPATH /go

# Change working directory
WORKDIR $GOPATH/src/goginapp/

# Install dependencies
ENV GO111MODULE=on
COPY . ./
USER root
RUN chmod -R 777 /go
USER default
RUN if test -e "go.mod"; then go build ./...; fi

ENV PORT 8080
ENV GIN_MODE release
EXPOSE 8080

RUN go build -o app
CMD ["./app"]
