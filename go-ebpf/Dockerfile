# Stage 1: Build the Go Fibonacci binary
FROM golang:1.22-bullseye AS builder
WORKDIR /src
COPY fibonacci.go .
RUN go build -gcflags="all=-N -l" -o fibonacci fibonacci.go

# Stage 2: Runtime with bpftrace
FROM debian:bullseye

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
