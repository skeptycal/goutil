#!/usr/bin/env zsh
# -*- coding: utf-8 -*-
    # shellcheck shell=bash
    # shellcheck source=/dev/null
    # shellcheck disable=2178,2128,2206,2034
#? -----------------------------> updatesubs.sh - add/remove/update git submodules
	#*  tools for repo management using macOS with zsh
	#*  tested on macOS v21.2.0 and zsh zsh v5.8 (x86_64-apple-darwin21.0)
	#*	copyright (c) 2021 Michael Treanor
	#*	MIT License - https://www.github.com/skeptycal
#? -----------------------------> https://www.github.com/skeptycal

	. $(which ansi_colors.sh) || return

#? -----------------------------> constants and flags
	DEV_PROD=0 		# set to 1 for production builds
    _DEV_GOMAKE=0	# set to 1 for gomake debug mode
	export default_repo_version='v0.1.0'
	local parsed_options=
	local tag=""
	local major=
	local minor=
	local patch=
	local dev=

	REPO_TEMPLATE_PATH="/Users/skeptycal/go/src/github.com/skeptycal/gorepotemplate"

	_init() {
		:
	}

#? -----------------------------> utilities

	# jq - commandline JSON processor [version 1.6]

		# Usage:	jq [options] <jq filter> [file...]
		# 	jq [options] --args <jq filter> [strings...]
		# 	jq [options] --jsonargs <jq filter> [JSON_TEXTS...]

		# jq is a tool for processing JSON inputs, applying the given filter to
		# its JSON text inputs and producing the filter's results as JSON on
		# standard output.

		# The simplest filter is ., which copies jq's input to its output
		# unmodified (except for formatting, but note that IEEE754 is used
		# for number representation internally, with all that that implies).

		# For more advanced filters see the jq(1) manpage ("man jq")
		# and/or https://stedolan.github.io/jq

		# Example:

		# 	$ echo '{"foo": 0}' | jq .
		# 	{
		# 		"foo": 0
		# 	}

		# For a listing of options, use jq --help.

	# _last prints the last field of the input based
	# on the separator in $1
	_last() {
		[[ -z "$1" ]] && return 1
		sep="$1"
		shift
		# echo "\$sep: $sep"
		# echo "\$@: $@"
		echo ${${@}##*${sep}}
		# echo "$@" | awk -F ${sep:-' '} '{print $NF}' # sucks
	}

	# _cut prints the last field of the input based
	# on the separator in $1
	_cut() {
		[[ -z "$1" ]] && return 1
		sep="$1"
		shift
		# echo "remove all '$sep' from $@"
		echo ${${@}//${sep}/}
	}

	# _pen prints the penultimate field of the inpu based
	# on the separator in $1
	_pen() {
		[[ -z "$1" ]] && return 1
		sep="$1"
		shift
		# echo "\$sep: $sep"
		# echo "\$@: $@"
		awk -F $sep '{print $NF-1}' "$@"
	}

	# _is_git_repo returns true if it is run within a git repository.
	# all text output is redirected to /dev/null
	_is_git_repo() { git status; } >/dev/null 2>&1

	# _tag prints the latest repo tag
	# errNo 128 is returned if git fails
	# any error text is sent to /dev/null
	_tag() { git rev-list --tags --max-count=1; } 2> /dev/null

	# _tag_version prints the latest git tag's version number
	# errNo 128 is returned if git fails
	# any error text is sent to /dev/null
	_tag_version() { git describe --tags $(_tag); } 2> /dev/null

	# push_tags pushes tags and any changes to remote.
	# all text output is redirected to /dev/null
	push_tags() {
		_is_git_repo || return 1
		git push origin --tags
        git push origin --all
	} >/dev/null 2>&1

	# go_version returns the current Go version number
	# errNo is returned if Go fails (not installed)
	go_version() { go version |cut -d ' ' -f 3 | cut -d 'o' -f 2; } 2> /dev/null

	# repo_version returns the semver version of the repository,
	# if the repo has tags. Otherwise, the $default_repo_version
	# is returned. (usually v0.1.0)
	repo_version() { $(_tag_version) || echo $default_repo_version; }

	# _process_version gets the current repo semver
	# version and breaks it up into individual variables:
	# major, minor, patch, dev
	_process_version() {
		dev=$(echo $(repo_version) | cut -d '-' -f 2)
		local vv=$(echo $(repo_version) | cut -d '-' -f 1)
        [ -z $vv ] && vv='v0.1.0'
        local vv=${vv#v*}
         major=${vv%%.*}
        vv=${vv#*.}
         minor=${vv%%.*}
        vv=${vv#*.}
         patch=${vv%%.*}
	}

	usage_simple () { me "Usage: $0 [options]"; }

	usage() {
        if [ -z "$1" ]; then
            white "usage: ${MAIN:-}${0} ${DARKGREEN:-}command [args]"
            return 1
        fi

        local app="$1"
        shift
        white "usage: ${MAIN:-}${app} ${DARKGREEN:-}${@}"
	}

	_options() {
		parsed_options=$(
		getopt -n "$0" -o t:Rf -- "$@"
		) || return 1
		eval "set -- $parsed_options"
		while [ "$#" -gt 0 ]; do
		case $1 in
			(-[Rf]) shift;;
			(-t) shift 2;;
			(--) shift; break;; # require '--'
			(*) exit 1 # should never be reached.
		esac
		done
		echo "Now, the arguments are $*"
	}

	# bump increments the version number depending on
	# the args passed in.
	#
	# bump [major|minor|patch|dev] [message]
	# default is 'dev' and no message (automated message)
	bump() {

		_options
        local vv=
        local dev=

		_process_version

        case "$1" in
            major)
                major=$(( major + 1 ))
                minor=0
                patch=0
                ;;

            minor)
                minor=$(( minor + 1 ))
                patch=0
                ;;

            patch)
                patch=$(( patch + 1 ))
                ;;

            dev)
                shift
                devtag "$@"
                return
                ;;

            *)
                echo "current version: $(_get_version)"
                usage $0 "bump [major|minor|patch|dev][message]"
                dbinfo "default case \$version: $version"
                return 0
                ;;

        esac

        printf -v version "v%s.%s.%s" $major $minor $patch

        echo "new version: $version"
        git tag "$version"
        git push origin --tags
        git push origin --all

	}
