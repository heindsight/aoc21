#!/bin/sh

set -eu

for answer_file in answers/*.txt; do
    puzzle="${answer_file%.txt}"
    puzzle="${puzzle#answers/}"
    infile="input/${puzzle%[ab]}.txt"
    ./aoc21 "${puzzle}" < "${infile}" | diff -u - "${answer_file}" || :
done
