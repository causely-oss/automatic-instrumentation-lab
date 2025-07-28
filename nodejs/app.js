// Main application that uses the fibonacci function
const { fibonacci } = require("./fib.js");

const n = process.argv[2] || 10;

console.log(`Calculating the ${n}th Fibonacci number...`);
const result = fibonacci(n);
console.log("Result:", result);
