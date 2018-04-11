# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# Copy the local package files to the container's workspace.
COPY . /go/src/learn/docker

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go install learn/docker && \
	rm -rf src/learn

# Run the command by default when the container starts.
ENTRYPOINT /go/bin/docker

# For service that listens on port 8080.
#EXPOSE 8080

# Execute:
# docker build -t arief/mygo:v1 .
# docker run -t arief/mygo:v1