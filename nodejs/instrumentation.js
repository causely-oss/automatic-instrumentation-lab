// Automatic instrumentation using require-in-the-middle
// This demonstrates how OpenTelemetry actually instruments modules

const hook = require("require-in-the-middle");
const path = require("path");

console.log("Fibonacci Instrumentation loaded!");

let callCounter = 0;
const startTimes = new Map();

const fibPath = path.join(__dirname, "fib.js");

console.log("Fib path:", fibPath);

// Hook into module loading - try different patterns
hook([fibPath], (exports, name, basedir) => {
  console.log(`Instrumenting module: ${name} from ${basedir}`);

  // Get the original fibonacci function from the exports object
  const originalFibonacci = exports.fibonacci;

  if (!originalFibonacci) {
    console.log("No fibonacci function found in exports");
    return exports;
  }

  // Create an instrumented version
  function instrumentedFibonacci(...args) {
    const callId = ++callCounter;
    const startTime = process.hrtime.bigint();
    startTimes.set(callId, startTime);

    // Call the original function
    const result = originalFibonacci.apply(this, args);

    // Calculate duration
    const endTime = process.hrtime.bigint();
    const duration = endTime - startTimes.get(callId);
    startTimes.delete(callId);

    console.log(
      `Call ${callId}: fibonacci(${args[0]}) took ${duration} nanoseconds, result: ${result}`,
    );

    return result;
  }

  exports.fibonacci = instrumentedFibonacci;

  console.log("Monkey patching completed successfully");

  return exports;
});
