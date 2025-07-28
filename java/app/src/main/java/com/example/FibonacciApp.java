package com.example;

public class FibonacciApp {
    public static void main(String[] args) {
        int n = 10; // default value
        
        // If a command line argument is provided, use it
        if (args.length > 0) {
            try {
                n = Integer.parseInt(args[0]);
            } catch (NumberFormatException e) {
                System.err.println("Error: Invalid number provided. Using default value 10.");
            }
        }
        
        System.out.println("Calculating the " + n + "th Fibonacci number...");
        long result = fibonacci(n);
        System.out.println("Result: " + result);
    }
    
    public static long fibonacci(int n) {
        if (n <= 1) {
            return n;
        }
        return fibonacci(n - 1) + fibonacci(n - 2);
    }
} 