#!/bin/bash

green='\033[0;32m'
red='\033[0;31m'
reset='\033[0m'

# Test 1: Snake to Camel
echo "Running Snake to Camel conversion..."
go run main.go test_assets/snake.py sc > test_assets/.tmp/_sc.test
go run main.go test_assets/.tmp/_sc.test cs > test_assets/.tmp/_snake.py

sc_test=$(diff test_assets/snake.py test_assets/.tmp/_snake.py)

if [ -z "$sc_test" ]; then
    echo -e "${green}Pass Test: Snake to Camel${reset}"
    rm test_assets/.tmp/_sc.test
    rm test_assets/.tmp/_snake.py
else
    echo -e "${red}Fail Test: Snake to Camel${reset}"
    diff test_assets/snake.py test_assets/.tmp/_snake.py
    exit
fi



# Test 2: Snake to Pascal
echo "Running Snake to Pascal conversion..."
go run main.go test_assets/snake.py sp > test_assets/.tmp/_sc.test
go run main.go test_assets/.tmp/_sc.test ps > test_assets/.tmp/_snake.py

sp_test=$(diff test_assets/snake.py test_assets/.tmp/_snake.py)

if [ -z "$sp_test" ]; then
    echo -e "${green}Pass Test: Snake to Pascal${reset}"
    rm test_assets/.tmp/_sc.test
    rm test_assets/.tmp/_snake.py
else
    echo -e "${red}Fail Test: Snake to Pascal${reset}"
    diff test_assets/snake.py test_assets/.tmp/_snake.py
    exit
fi



# Test 3: Camel to Snake
echo "Running Camel to Snake conversion..."
go run main.go test_assets/camel.js cs > test_assets/.tmp/_cs.test
go run main.go test_assets/.tmp/_cs.test sc > test_assets/.tmp/_camel.js

cs_test=$(diff test_assets/camel.js test_assets/.tmp/_camel.js)

if [ -z "$cs_test" ]; then
    echo -e "${green}Pass Test: Camel To Snake${reset}"
    rm test_assets/.tmp/_cs.test
    rm test_assets/.tmp/_camel.js
else
    echo -e "${red}Fail Test: Camel to Snake${reset}"
    diff test_assets/camel.js test_assets/.tmp/_camel.js
    exit
fi



# Test 4: Camel to Pascal
echo "Running Camel to Pascal conversion..."
go run main.go test_assets/camel.js cp > test_assets/.tmp/_cp.test
go run main.go test_assets/.tmp/_cp.test pc > test_assets/.tmp/_camel.js

cs_test=$(diff test_assets/camel.js test_assets/.tmp/_camel.js)

if [ -z "$cs_test" ]; then
    echo -e "${green}Pass Test: Camel To Pascal${reset}"
    rm test_assets/.tmp/_cp.test
    rm test_assets/.tmp/_camel.js
else
    echo -e "${red}Fail Test: Camel to Pascal${reset}"
    diff test_assets/camel.js test_assets/.tmp/_camel.js
    exit
fi


# Test 5: Pascal to Snake
echo "Running Pascal to Snake conversion..."
go run main.go test_assets/pascal.f90 ps > test_assets/.tmp/_ps.test
go run main.go test_assets/.tmp/_ps.test sp > test_assets/.tmp/_pascal.f90

cs_test=$(diff test_assets/pascal.f90 test_assets/.tmp/_pascal.f90)

if [ -z "$cs_test" ]; then
    echo -e "${green}Pass Test: Pascal To Snake${reset}"
    rm test_assets/.tmp/_ps.test
    rm test_assets/.tmp/_pascal.f90
else
    echo -e "${red}Fail Test: Pascal To Snake${reset}"
    diff test_assets/pascal.f90 test_assets/.tmp/_pascal.f90
    exit
fi



# Test 6: Pascal to Camel
echo "Running Pascal to Camel conversion..."
go run main.go test_assets/pascal.f90 pc > test_assets/.tmp/_pc.test
go run main.go test_assets/.tmp/_pc.test cp > test_assets/.tmp/_pascal.f90

cs_test=$(diff test_assets/pascal.f90 test_assets/.tmp/_pascal.f90)

if [ -z "$cs_test" ]; then
    echo -e "${green}Pass Test: Pascal To Camel${reset}"
    rm test_assets/.tmp/_pc.test
    rm test_assets/.tmp/_pascal.f90
else
    echo -e "${red}Fail Test: Pascal To Camel${reset}"
    diff test_assets/pascal.f90 test_assets/.tmp/_pascal.f90
    exit
fi



