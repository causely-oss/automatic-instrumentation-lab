# Automatic Instrumentation Lab

[![Maintained by Causely](https://img.shields.io/badge/Maintained%20by-Causely.ai-blue)](https://www.causely.ai)
[![License: Apache 2.0](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

This repository contains a lab to explore different techniques of automatic instrumentation. [Automatic Instrumentation](https://opentelemetry.io/docs/concepts/glossary/#automatic-instrumentation) _"[r]efers to telemetry collection methods that do not require the end-user to modify application's source code. Methods vary by programming language, and examples include bytecode injection or monkey patching."_

You will find the following techniques:

- [Monkey Patching](#monkey-patching-nodejs) [📝](./nodejs/) [🎬](https://asciinema.org/a/AqBkvGMjXqLuYIKtP3QdfKHTS)
- [Byte Code Instrumentation](#byte-code-instrumentation-java) [📝](./java/) [🎬](https://asciinema.org/a/CbCsDOaheyM4NvPLFq3G2RVyT)
- [Compile-Time Instrumentation](#compile-time-instrumentation-go) [📝](./go-compile-time/) [🎬](https://asciinema.org/a/CZKVkiWCenE42OOPdJYKB9SSK)
- [eBPF based instrumentation](#ebpf-based-instrumentation-go) [📝](./go-ebpf/) [🎬](https://asciinema.org/a/6Z3nJP8mg0ZlZEpiAT4QHrQXj)
- [Observer API (PHP)](#php-observer-api-php) [📝](./php/) [🎬](https://asciinema.org/a/OX9193zka7CKb7J1yTdQy5dpx)

The examples in this repository are for educational purpose to provide insights into those different techniques and to uncover how they work. The code here should not be used as a starting point for real implementations, since certain details are skipped or tooling is available that is better suited for real world use cases. If you want to learn more about this topic, after going through this lab you should take a look into implementations by the OpenTelemetry project, e.g.

- Monkey Patching
  - [OpenTelemetry JavaScript Instrumentation Libraries](https://github.com/open-telemetry/opentelemetry-js-contrib/tree/main/packages)
  - [OpenTelemetry Python Instrumentation Libraries](https://github.com/open-telemetry/opentelemetry-python-contrib/tree/main/instrumentation)
- Byte Code Instrumentation
  - [OpenTelemetry Java Instrumentation Libraries](https://github.com/open-telemetry/opentelemetry-java-instrumentation/tree/main/instrumentation)
  - [OpenTelemetry .NET Instrumentation Libraries](https://github.com/open-telemetry/opentelemetry-dotnet-contrib/tree/main/src)
- eBPF-based Instrumentation
  - [OpenTelemetry eBPF Instrumentation (OBI)](https://github.com/open-telemetry/opentelemetry-ebpf-instrumentation)
- Compile-time Instrumentation 
  - [OpenTelemetry Go Compile-Time Instrumentation](https://github.com/open-telemetry/opentelemetry-go-compile-instrumentation)
- PHP Observer API Instrumentation
  - [OpenTelemetry PHP Instrumentation Libraries](https://github.com/open-telemetry/opentelemetry-php-contrib/tree/main/src/Instrumentation)

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

**Demo**: Watch the instrumented application in action on [asciinema](https://asciinema.org/a/AqBkvGMjXqLuYIKtP3QdfKHTS)

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

**Demo**: Watch the instrumented application in action on [asciinema](https://asciinema.org/a/CbCsDOaheyM4NvPLFq3G2RVyT)

**OpenTelemetry Projects**:

- [OpenTelemetry Java](https://opentelemetry.io/docs/zero-code/java/agent/) - Uses bytecode instrumentation for automatic instrumentation
- [OpenTelemetry .NET](https://opentelemetry.io/docs/zero-code/dotnet/) - Uses CIL instrumentation for automatic instrumentation

### Compile-Time Instrumentation (Go)

Compile-time instrumentation involves modifying source code during the build process to inject observability code. This technique uses Abstract Syntax Tree (AST) manipulation to transform code before compilation, ensuring that instrumentation is baked into the final binary. This approach provides zero runtime overhead and works well with statically compiled languages.

**Languages**: This technique is most effective with statically compiled languages like Go, Rust, C++, and other languages that support AST manipulation during compilation.

**How to Run the Tutorial**:

```bash
cd go-compile-time
make build
./fibonacci
```

**Demo**: Watch the instrumented application in action on [asciinema](https://asciinema.org/a/CZKVkiWCenE42OOPdJYKB9SSK)

**OpenTelemetry Projects**:

- [OpenTelemetry Go Compile Instrumentation](https://github.com/open-telemetry/opentelemetry-go-compile-instrumentation) - Official OpenTelemetry project for Go compile-time instrumentation

### eBPF-based Instrumentation (Go)

eBPF (Extended Berkeley Packet Filter) instrumentation leverages kernel-level tracing capabilities to observe application behavior without modifying the application code. This technique uses BPF programs that run in the kernel to attach probes to function entry and exit points, providing deep system-level observability with minimal overhead.

**Languages**: eBPF instrumentation is language-agnostic and can be applied to almost any language running on Linux. It works at the system level, making it independent of the application's programming language.

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
go go build -gcflags="all=-N -l" -o fibonacci fibonacci.go

# Run with bpftrace (requires root)
sudo bpftrace trace.bt &
./fibonacci
```

**Demo**: Watch the instrumented application in action on [asciinema](https://asciinema.org/a/6Z3nJP8mg0ZlZEpiAT4QHrQXj)

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

**Demo**: Watch the instrumented application in action on [asciinema](https://asciinema.org/a/OX9193zka7CKb7J1yTdQy5dpx)

**OpenTelemetry Projects**:

- [OpenTelemetry PHP](https://opentelemetry.io/docs/zero-code/php/) - Uses the PHP Observer API for automatic instrumentation

## Testing All Instrumentations

To verify that all instrumentations work correctly, you can run the comprehensive test script:

```bash
# Make the script executable (if not already)
chmod +x test-all-instrumentations.sh

# Run all tests
./test-all-instrumentations.sh
```

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request for any enhancements or bug fixes.

## License

This project is licensed under the Apache 2.0 License - see the [LICENSE](LICENSE) file for details.
