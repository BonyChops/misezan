#!/bin/sh

jq -r '.basic[], .extra[] | "\(.a),\(.b),\(.ans)"' ../test/test_cases.json > test_cases.csv
jq -r '.extraFloat[] | "\(.a),\(.b),\(.ans)"' ../test/test_cases.json > test_cases_float.csv