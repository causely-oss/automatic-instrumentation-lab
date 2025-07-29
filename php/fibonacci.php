<?php

function fib($n) {
    if ($n <= 1) {
        return $n;
    }
    return fib($n - 1) + fib($n - 2);
}

// Get the number from command line argument or use default
$n = isset($argv[1]) ? (int)$argv[1] : 10;

echo "Calculating the {$n}th Fibonacci number...\n";
$result = fib($n);
echo "Result: $result\n"; 