# Automatic Instrumentation Lab

This repository contains a lab to explore different techniques of automatic instrumentation. [Automatic Instrumentation](https://opentelemetry.io/docs/concepts/glossary/#automatic-instrumentation) _"[r]efers to telemetry collection methods that do not require the end-user to modify application's source code. Methods vary by programming language, and examples include bytecode injection or monkey patching."_

The examples in this repository are for educational purpose to provide insights into those different techniques and to uncover how they work. The code here should not be used as a starting point for real implementations, since certain details are skipped or tooling is available that is better suited for real world use cases. It is also recommended

You will find the following techniques:

- [Monkey Patching](#monkey-patching-nodejs)
- [Byte Code Instrumentation](#byte-code-instrumentation-java)
- [Compile-Time Instrumentation](#compile-time-instrumentation-go)
- [eBPF based instrumentation](#ebpf-based-instrumentation-go)
- [Observer API (PHP)](#php-observer-api-php)

## Instrumentation Techniques

### Monkey Patching (Node.js)

Monkey patching is a technique that involves modifying or extending existing code at runtime by replacing functions, methods, or modules with instrumented versions. This approach intercepts function calls and adds observability without requiring changes to the original source code.

**Languages**: This technique is commonly used in dynamic languages like JavaScript (Node.js), Python, Ruby, and other interpreted languages that support runtime modification of functions and modules.

**How to Run the Tutorial**:

```bash
cd nodejs/
# Install dependencies
npm install

# Run without instrumentation (normal execution)
node app.js 5

# Run with instrumentation (instrumented execution)
node -r ./instrumentation.js app.js 5
```

**OpenTelemetry Projects**:

- [OpenTelemetry JavaScript](https://opentelemetry.io/docs/zero-code/js/) - Uses monkey patching for automatic instrumentation of Node.js applications
- [OpenTelemetry Python](https://opentelemetry.io/docs/zero-code/python/) - Employs monkey patching for automatic instrumentation

### Byte Code Instrumentation (Java)

Byte code instrumentation involves modifying the compiled bytecode of applications at runtime to add observability hooks. This technique works at the JVM level, allowing instrumentation of any Java application without source code access. It leverages the Java Instrumentation API to transform classes as they are loaded by the JVM.

**Languages**: This technique is primarily used in JVM-based languages (Java, Kotlin, Scala, Groovy) and can also be applied to other bytecode-based languages like .NET (CIL instrumentation).

**How to Run the Tutorial**:

```bash
cd java/
# Build the project
./gradlew build

# Run without agent (normal execution)
java -jar app/build/libs/app-1.0.0.jar 5

# Run with agent (instrumented execution)
java -javaagent:agent/build/libs/agent-1.0.0-all.jar -jar app/build/libs/app-1.0.0.jar 5
```

**OpenTelemetry Projects**:

- [OpenTelemetry Java](https://opentelemetry.io/docs/zero-code/java/agent/) - Uses bytecode instrumentation for automatic instrumentation
- [OpenTelemetry .NET](https://opentelemetry.io/docs/zero-code/dotnet/) - Uses CIL instrumentation for automatic instrumentation

### Compile-Time Instrumentation (Go)

Compile-time instrumentation involves modifying source code during the build process to inject observability code. This technique uses Abstract Syntax Tree (AST) manipulation to transform code before compilation, ensuring that instrumentation is baked into the final binary. This approach provides zero runtime overhead and works well with statically compiled languages.

**Languages**: This technique is most effective with statically compiled languages like Go, Rust, C++, and other languages that support AST manipulation during compilation.

**How to Run the Tutorial**:

```bash
cd go-compile-time
make run
```

This will:

- Build the integrated instrumentation tool
- Instrument the Fibonacci app at build time
- Run the instrumented binary

**OpenTelemetry Projects**:

- [OpenTelemetry Go Compile Instrumentation](https://github.com/open-telemetry/opentelemetry-go-compile-instrumentation) - Official OpenTelemetry project for Go compile-time instrumentation

### eBPF-based Instrumentation (Go)

eBPF (Extended Berkeley Packet Filter) instrumentation leverages kernel-level tracing capabilities to observe application behavior without modifying the application code. This technique uses BPF programs that run in the kernel to attach probes to function entry and exit points, providing deep system-level observability with minimal overhead.

**Languages**: eBPF instrumentation is language-agnostic and can be applied to any compiled language running on Linux (C, C++, Go, Rust, etc.). It works at the system level, making it independent of the application's programming language.

**How to Run the Tutorial**:

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

**OpenTelemetry Projects**:

- [OpenTelemetry eBPF](https://github.com/open-telemetry/opentelemetry-ebpf-instrumentation) - Official OpenTelemetry project for eBPF-based instrumentation

### PHP Observer API (PHP)

The PHP Observer API is a low-level instrumentation technique that hooks directly into the PHP engine's execution flow. This approach uses C extensions to observe function calls at the Zend engine level, providing deep visibility into PHP application behavior without requiring code modifications. It operates at the language runtime level, similar to how other languages implement their instrumentation APIs.

**Languages**: This technique is specific to PHP and leverages the Zend Observer API introduced in PHP 8.0+. Similar approaches exist in other languages through their respective runtime APIs (e.g., Python's sys.settrace, Ruby's TracePoint).

**How to Run the Tutorial**:

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

**OpenTelemetry Projects**:

- [OpenTelemetry PHP](https://opentelemetry.io/docs/zero-code/php/) - Uses the PHP Observer API for automatic instrumentation

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request for any enhancements or bug fixes.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
