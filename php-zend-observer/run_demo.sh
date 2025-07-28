#!/bin/bash

# Clean previous builds
make clean

# Prepare the build environment
phpize

# Configure the build
./configure --enable-observer

# Compile the extension
make

# Set the PHPRC environment variable to use the local php.ini
export PHPRC=$(pwd)

# Run the PHP demo script
php demo.php 