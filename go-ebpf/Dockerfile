# Stage 1: Build the Go Fibonacci binary
FROM golang:1.24-bullseye@sha256:637f45ef9f8fb4228406268d544df3f1251703cda025f706902c3627fa621c54 AS builder
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
