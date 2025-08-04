#!/bin/bash

# Build
# make clean
phpize
./configure --enable-observer
make

# We need to make sure that the php.ini which loads the created library is found when the application is run
export PHPRC=$(pwd)
# Run
php fibonacci.php 
