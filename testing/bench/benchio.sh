#!/bin/zsh

go build -gcflags='-m -m' io | grep 'escapes to heap'

# escapes to heap
eth() {
    # only works for packages in $GOROOT (/usr/local/go/src)
    # go build -gcflags='-m -m' io 2>&1 | grep 'escapes to heap'
    go build -gcflags='-m -m' "$1" 2>&1 | grep 'escapes to heap'
}

# function too complex
ftc() {
    go build -gcflags='-m -m' "$1" 2>&1 | grep 'function too complex'
}

# is in bounds (inserting bounds checks)
iib() {
        go build -gcflags='-d=ssa/check_bce/debug=1' "$1" 2>&1
}

# bounds check proof
bcp() {
    go build -gcflags='-d=ssa/prove/debug=2' "$1" 2>&1
}