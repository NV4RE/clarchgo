############################
# STEP 1 setup for build project
############################
FROM golang:1.17-alpine AS builder
WORKDIR $GOPATH/src/gitlab.com/sendit-th/drivs/lrd/4pl-dispatching-go/
COPY . .
# Enable go module
ENV GO111MODULE=on
# Fetch dependencies
RUN go mod download
# Build the binary
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" -o /go/bin/app

############################
# STEP 2 build a small image
############################
FROM golang:1.17-stretch
# Copy our static executable.
COPY --from=builder /go/bin/app /go/bin/app
# Run the gogo-blueprint binary
ENTRYPOINT ["/go/bin/app"]
