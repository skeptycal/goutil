#!/bin/zsh

function countArguments() {
    echo "${#@}"
}

wordlist="one two three four five"

echo "normal substitution, no quotes:"
countArguments $wordlist
echo ""
# -> 5

echo "substitution with quotes:"
countArguments "$wordlist"
echo ""
# -> 1