#!/bin/zsh

usage() {
cat << EOF
Usage:
  $0 [-u [username]] [-p]
  Options:
    -u <username> : Optionally specify the new username to set password for.
    -p : Prompt for a new password.
EOF
}

cat <<PACK >|f.go
package main
func HelloWorld() {
    println("hello, world!")
}

func main () {
  HelloWorld()
}
PACK

GOSSAFUNC=HelloWorld go build
# dumped SSA to ssa.html
open ssa.html