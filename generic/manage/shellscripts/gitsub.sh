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

#? -----------------------------> imports
	. $(which gitutil.sh) || . ./gitutil.sh

	#* DEV_PROD, _DEV_GOMAKE, default_repo_version, parsed_options, tag, major, minor, patch, dev
	#* REPO_TEMPLATE_PATH="/Users/skeptycal/go/src/github.com/skeptycal/gorepotemplate"
	#	_last() 				#/ # _last prints the last field of the input
	#	_cut() 					#/ # _cut deletes matching substrings from a string
	#	_pen() 					#/ # _pen prints the penultimate field of the input
	#	_is_git_repo() 			#/ # _is_git_repo returns true if it is run within a git repository.
	#	_tag() 					#/ # _tag prints the latest repo tag
	#	_tag_version() 			#/ # _tag_version prints the latest git tag's version number
	#	push_tags() 			#/ # push_tags pushes tags and any changes to remote.
	#	go_version() 			#/ # go_version returns the current Go version number
	#	repo_version() 			#/ # repo_version returns the semver version of the repository...
	#	_process_version()		#/ # _process_version gets the current repo semver
	#	usage_simple ()			#/ # usage_simple is a basic usage message
	#	usage() 				#/ # usage is an ANSI colored usage message
	#	_options() 				#/ # _options processes CLI arguments
	#	bump() 					#/ # bump [major|minor|patch|dev] [message]

	#	jq 						#/ # commandline JSON processor [version 1.6]
	#	dirlist 				#/ list of non-hidden directories

#? -----------------------------> functions

_init() { : }

_foreach() {
	local force=0
	if [[ "$1" == "-f" ]]; then
		force=1
		shift
	fi

	attn "\$force: $force"
	attn "commands: $@"

	for f in $(dirlist); do
		cd $f
		attn "\$PWD: $PWD"
		$@
		cd -
		attn "\$PWD: $PWD"
	done;
}

ugaa () {
	_foreach -f "git add -A"
}


_init
usage "_foreach" "[-f][command]"
_foreach "$@"
ugaa
