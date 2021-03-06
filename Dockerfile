# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# Set working directory
RUN mkdir -p /go/src/app
WORKDIR /go/src/app

# Copy the local package files to the container's workspace.
COPY src/ /go/src/app
COPY get.sh /tmp

# Install GO DEPS
RUN bash /tmp/get.sh

# Build the command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a
RUN go build -o /go/src/app/api /go/src/app/*

# Run the api command by default when the container starts.
ENTRYPOINT /go/src/app/api

# Document that the service listens on port 8080.
EXPOSE 8080

# Turn your head and cough
HEALTHCHECK CMD curl --fail http://localhost:8080/healthcheck || exit 1
