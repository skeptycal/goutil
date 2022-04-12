#!/usr/bin/env zsh
# -*- coding: utf-8 -*-
    # shellcheck shell=bash
    # shellcheck source=/dev/null
    # shellcheck disable=2178,2128,2206,2034
#? -----------------------------> gitsub.sh - add/remove/update git submodules
	#*  tools for repo management using macOS with zsh
	#*  tested on macOS v21.2.0 and zsh zsh v5.8 (x86_64-apple-darwin21.0)
	#*	copyright (c) 2021 Michael Treanor
	#*	MIT License - https://www.github.com/skeptycal

local usage='usage: $0 <package name>'
local outpath=$PWD/.testoutput
mkdir -p $outpath
local outfile=$outpath/heap.txt

exists() { command -v "$1" 2>&1; }
die () { echo ${*:-"die (pid = $$)"} && exit 1; }
noargs() {  [[ -z "$1" ]] && die $usage; }

escapes_to_heap() {
    noargs "$1"
    echo "noargs: $1"
    _package="${1:-"./main.go"}"; shift;
    go build -gcflags='-m -m' "${package}" >|$outfile
    cat $outfile | grep "escapes"
}

_main() {
    escapes_to_heap "$1"
}

# TODO does this need to be $@ to pass further build options???
_main "$1"