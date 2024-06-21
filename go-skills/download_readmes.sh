#!/bin/bash

# Base URL of the repository
BASE_URL="https://learn.zone01kisumu.ke/git/root/public/src/branch/master/subjects"

# List of directories to extract README.md from
directories=(
    "sortwordarr"
    "ispowerof2"
    "findprevprime"
    "lastrune"
    "strlen"
    "displayalrevm"
    "searchreplace"
    "lastword"
    "countdown"
    "displayalpham"
    "displayfirstparam"
    "displaylastparam"
    "hello"
    "paramcount"
    "printdigits"
    "max"
    "wdmatch"
    "firstrune"
    "reduceint"
    "rot13"
    "alphamirror"
    "chunk"
    "compare"
    "doop"
    "foldint"
    "tabmult"
    "inter"
    "piglatin"
    "printbits"
    "reversebits"
    "romannumbers"
    "swapbits"
    "union"
    "gcd"
    "printhex"
    "repeatalpha"
    "addprimesum"
    "fprime"
    "hiddenp"
    "revwstr"
    "rostring"
    "split"
)

# Directory to save the renamed README.md files
destination_dir="./extracted_readmes"
mkdir -p "$destination_dir"

# Loop through directories and download README.md
for dir in "${directories[@]}"; do
    url="$BASE_URL/$dir/README.md"
    output_file="$destination_dir/README.md.$dir"
    
    # Download the README.md file and save it with the new name
    wget -q -O "$output_file" "$url"

    # Check if the download was successful
    if [[ ! -s "$output_file" ]]; then
        echo "README.md not found in $dir or download failed"
        rm -f "$output_file"
    fi
done
