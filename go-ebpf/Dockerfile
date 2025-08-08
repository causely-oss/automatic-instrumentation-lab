# Stage 1: Build the Go Fibonacci binary
FROM golang:1.22-bullseye@sha256:c89d51959c480f13e7830a33fa0174480227c8a027d0e18066e79a03cde174c4 AS builder
WORKDIR /src
COPY fibonacci.go .
RUN go build -gcflags="all=-N -l" -o fibonacci fibonacci.go

# Stage 2: Runtime with bpftrace
FROM debian:bullseye@sha256:25c0cab214b810db1b3c8adef5a12a92596979abddf86bb364e8d9c9d111df9f

# Install bpftrace + tools
RUN apt-get update && \
    apt-get install -y bpftrace linux-perf clang llvm ca-certificates && \
    apt-get clean

# Copy binary and bpftrace script
WORKDIR /app
COPY --from=builder /src/fibonacci /app/fibonacci
COPY trace.bt /app/trace.bt

# Default command: run bpftrace
CMD ["bpftrace", "/app/trace.bt"]
