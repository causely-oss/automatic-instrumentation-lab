# Automatic Instrumentation Lab

This project is a multi-language lab designed to explore automatic instrumentation techniques across different programming environments. It includes implementations in Go, Java, Node.js, and PHP, each demonstrating unique aspects of instrumentation.

## Instrumentation Mechanisms

### Go Compile-Time Instrumentation

**Location**: `go-compile-time/` directory
**Structure**:
- `app/` - Contains the standalone Fibonacci application
- `build-tool.go` - Minimal AST-based build tool for instrumentation
- `Makefile` and `build-integrated.sh` - Build automation

**Prerequisites**: Go 1.16+

**Mechanism**: AST (Abstract Syntax Tree) manipulation during build

- Uses Go's `go/ast` package to parse and modify the source code
- Injects a timing function and a `defer` statement into the `fibonacci` function
- Adds the `time` import to the file
- Demonstrates compile-time code transformation for automatic tracing

**How to Run**:

```bash
cd go-compile-time
make run
```

This will:
- Build the integrated instrumentation tool
- Instrument the Fibonacci app at build time
- Run the instrumented binary

**What gets instrumented:**
- The `fibonacci` function is wrapped with a `defer trace_fibonacci()()` call
- A timing function is injected to measure execution time
- The `time` import is added if not present

**Example Output:**
```
Function fibonacci took: 9.625µs
Function fibonacci took: 41ns
...
55
```

This approach is similar to how production tools like [OpenTelemetry Go Compile Instrumentation](https://github.com/open-telemetry/opentelemetry-go-compile-instrumentation) work, but is simplified for tutorial and demonstration purposes.

### Go eBPF Instrumentation

**Location**: `go-ebpf/` directory
**Files**:

- `fibonacci.go`: Fibonacci sequence generator
- `trace.bt`: BPF trace script for function instrumentation
- `Dockerfile`: Docker setup for eBPF environment

**Prerequisites**: Go 1.16+, Linux kernel with eBPF support, Docker (optional)

**Mechanism**: Extended Berkeley Packet Filter (eBPF) for runtime tracing

- Uses bpftrace to attach uprobes/uretprobes to function entry/exit points
- Non-intrusive instrumentation that doesn't require code modification
- Measures function execution time using kernel-level tracing
- Demonstrates system-level observability without application changes

**How to Run**:

**Docker (recommended for macOS/Windows)**:

```bash
cd go-ebpf/
# Build the Docker image
docker buildx build --platform linux/arm64 -t fibtrace .

# Run the container with eBPF tracing
docker run --rm -it --privileged fibtrace

# Or run just the Fibonacci application
docker run --rm -it --privileged fibtrace /app/fibonacci
```

**Linux (native)**:

```bash
cd go-ebpf/
# Build the Go application
go build -o fibonacci fibonacci.go

# Run with bpftrace (requires root)
sudo bpftrace trace.bt &
./fibonacci
```

### Java Agent Instrumentation

**Location**: `java/` directory
**Files**:

- `agent/src/main/java/com/example/FibonacciAgent.java`: Java agent for instrumenting the Fibonacci application
- `app/src/main/java/com/example/FibonacciApp.java`: Main application demonstrating Fibonacci sequence
- `build.gradle`: Gradle build configuration
- `gradlew`: Gradle wrapper script

**Prerequisites**: JDK 21+

**Mechanism**: Java Instrumentation API with ByteBuddy

- Uses Java's `java.lang.instrument` package for runtime bytecode modification
- ByteBuddy library for dynamic bytecode generation and manipulation
- Implements a Java agent that intercepts method calls at runtime
- Tracks call stack and execution time for specific methods
- Demonstrates JVM-level instrumentation without source code changes

**How to Run**:

```bash
cd java/
# Build the project
./gradlew build

# Run without agent (normal execution)
java -jar app/build/libs/app-1.0.0.jar 5

# Run with agent (instrumented execution)
java -javaagent:agent/build/libs/agent-1.0.0-all.jar -jar app/build/libs/app-1.0.0.jar 5
```

### Node.js Module Instrumentation

**Location**: `nodejs/` directory
**Files**:

- `app.js`: Main application file
- `fib.js`: Fibonacci sequence logic
- `instrumentation.js`: Instrumentation logic for Node.js
- `package.json`: Node.js package configuration

**Prerequisites**: Node.js 14+ with npm

**Mechanism**: Module loading interception with require-in-the-middle

- Uses `require-in-the-middle` library to hook into Node.js module loading
- Intercepts module exports and wraps functions with instrumentation
- Monkey-patches functions to add timing and logging
- Demonstrates how OpenTelemetry-style instrumentation works in Node.js
- Non-intrusive approach that works with existing code

**How to Run**:

```bash
cd nodejs/
# Install dependencies
npm install

# Run without instrumentation (normal execution)
node app.js 5

# Run with instrumentation (instrumented execution)
node -r ./instrumentation.js app.js 5
```

### PHP Zend Observer Instrumentation

**Location**: `php-zend-observer/` directory
**Files**:

- `fibonacci.php`: Demo script for PHP instrumentation
- `observer.c`: C extension for PHP instrumentation
- `php_observer.h`: Header file for the PHP observer
- `php.ini`: PHP configuration file for the demo
- `config.m4`: Build configuration for the extension

**Prerequisites**: PHP 8.0+ with development headers and build tools

**Mechanism**: PHP Zend Observer API with C extension

- Implements a PHP C extension using the Zend Observer API
- Hooks into PHP function execution at the engine level
- Observes function entry and exit points without code modification
- Measures execution time and provides cumulative timing statistics
- Demonstrates low-level PHP engine instrumentation

**How to Run**:

```bash
cd php-zend-observer/

# Build the PHP extension
make clean
phpize
./configure --enable-observer
make

# Run without instrumentation (normal execution)
php fibonacci.php 5

# Run with instrumentation (instrumented execution)
export PHPRC=$(pwd)
php fibonacci.php 5
```

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request for any enhancements or bug fixes.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
