# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:alpine

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/schambers/go-samples/go-docker/webapp

# Build the webapp command inside the container.
RUN go install github.com/schambers/go-samples/go-docker/webapp

# Run the webapp command by default when the container starts.
ENTRYPOINT /go/bin/webapp

# Document that the service listens on port 8080.
EXPOSE 8080
