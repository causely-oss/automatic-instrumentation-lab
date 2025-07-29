#!/bin/bash

# Build
make clean
phpize
./configure --enable-observer
make

# Run
export PHPRC=$(pwd)
php fibonacci.php 