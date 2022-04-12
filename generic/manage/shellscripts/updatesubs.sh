#!/usr/bin/env zsh
# -*- coding: utf-8 -*-
    # shellcheck shell=bash
    # shellcheck source=/dev/null
    # shellcheck disable=2178,2128,2206,2034
#? -----------------------------> updatesubs.sh - update all git submodules
	#*  tools for repo management for macOS with zsh
	#*  tested on macOS Big Sur and zsh 5.8
	#*	copyright (c) 2021 Michael Treanor
	#*	MIT License - https://www.github.com/skeptycal
#? -----------------------------> https://www.github.com/skeptycal

#? -----------------------------> utilities
    # IS_PROD=0
	# DEV_PROD=0
    # _DEV_GOMAKE=0
	# default_repo_version='v0.1.0'
	# REPO_TEMPLATE_PATH="/Users/skeptycal/go/src/github.com/skeptycal/gorepotemplate"

	# _options() {
	# 	parsed_options=$(
	# 	getopt -n "$0" -o t:Rf -- "$@"
	# 	) || return 1
	# 	eval "set -- $parsed_options"
	# 	while [ "$#" -gt 0 ]; do
	# 	case $1 in
	# 		(-[Rf]) shift;;
	# 		(-t) shift 2;;
	# 		(--) shift; break;;
	# 		(*) exit 1 # should never be reached.
	# 	esac
	# 	done
	# 	echo "Now, the arguments are $*"
	# }

    # is_prod() { [ $IS_PROD -eq 0 ]; }
    # is_dev() { [ $IS_PROD -ne 0 ]; }

	# export go_version() { go version |cut -d ' ' -f 3 | cut -d 'o' -f 2; }

	# export repo_version() {
    #     # echo $(git describe --tags $(git rev-list --tags --max-count=1))
    #     if git describe --tags > /dev/null 2>&1; then
    #         git tag | sort -V | tail -n 1
    #     else
    #         echo "$default_repo_version"
    #     fi
	# }

	# usage() {
    #     if [ -z "$1" ]; then
    #         white "usage: ${MAIN:-}${0} ${DARKGREEN:-}app [args]"
    #         return 1
    #     fi

    #     local app="$1"
    #     shift
    #     white "usage: ${MAIN:-}${app} ${DARKGREEN:-}${@}"
	# }

	# bump() {
    #     local parsed_options=$(
	# 	getopt -n "$0" -o hi -- "$@"
	# 	) || return 1

    #     local vv=
    #     local dev=

    #     vv=$(echo $(repo_version) | cut -d '-' -f 1)
    #     [ -z $vv ] && vv='v0.1.0'
    #     local vv=${vv#v*}
    #     local major=${vv%%.*}
    #     vv=${vv#*.}
    #     local minor=${vv%%.*}
    #     vv=${vv#*.}
    #     local patch=${vv%%.*}

    #     case "$1" in
    #         major)
    #             major=$(( major + 1 ))
    #             minor=0
    #             patch=0
    #             ;;

    #         minor)
    #             minor=$(( minor + 1 ))
    #             patch=0
    #             ;;

    #         patch)
    #             patch=$(( patch + 1 ))
    #             ;;

    #         dev)
    #             shift
    #             devtag "$@"
    #             return
    #             ;;

    #         *)
    #             echo "current version: $(_get_version)"
    #             usage bump '[major|minor|patch|dev][message]'
    #             dbinfo "default case \$version: $version"
    #             return 0
    #             ;;

    #     esac

    #     printf -v version "v%s.%s.%s" $major $minor $patch

    #     echo "new version: $version"
    #     git tag "$version"
    #     git push origin --tags
    #     git push origin --all

	# }

	devtag () {
		_setup_variables
		local vv=$(_get_version)
		local version_file=${REPO_PATH}/${REPO_NAME}/.VERSION
		echo $version_file
		dbinfo "\$version_file: $version_file"
		local dev=
		local message=
		if [ -n "$1" ]
		then
			message="(GoBot) devtag version ${version}: ${1}"
			shift
			echo "$message"
		fi
		dbinfo "\$message: $message"
		vv=${vv%%-*}
		[ -z $vv ] && vv='v0.1.0'
		dbinfo "\$vv: $vv"
		printf -v dev "%16.16s" $(/opt/homebrew/bin/gdate +'%s%N')
		dbinfo "\$dev: $dev"
		version="${vv}-${dev}"
		dbinfo "\$version: $version"
		echo $version >| $version_file
		go mod tidy > /dev/null 2>&1
		go doc >| go.doc
		git add $version_file
		git add go.mod
		git add go.sum
		git add go.doc
		git tag "$version"
		if [ -n "$message" ]
		then
			git commit -m "$message"
		else
			git commit -m "(GoBot) devtag version ${version}"
		fi
		git push origin --tags
		git push origin --all
		white "new Git version tag: ${MAIN:-}$version"
	}

	clean_template_name() {
        if [ -z "$1" ]; then
            usage $0 '[files]'
            return 0
        fi
        for f in $@; do
            sed -i '' -e "s/gorepotemplate/${REPO_NAME}/g" ${f}
        done
	}

	template() {

		# i=0
		# until [ "$((i=$i+1))" -gt "$#" ]
		# do case "$1"                   in
		# --Recursive) set -- "$@" "-R"  ;;
		# --file)      set -- "$@" "-f"  ;;
		# --target)    set -- "$@" "-t"  ;;
		# *)           set -- "$@" "$1"  ;;
		# esac; shift; done

		# local parsed_options=$(
		# getopt -n "$0" -o bfhiv -- "$@"
		# ) || return 1

		# attn "\$parsed_options: $parsed_options"

        _setup_variables || dbecho "could not setup variables"
		local _usage_string="[--force|--bump|--help|--version] [files]"
		local _force=

        case "$1" in

            -v|--version|version)

				local version=$(_get_version) >/dev/null 2>&1
				if [ -z "$version" ]; then
					usage "$0 (no version set)"
				else
					usage "$0 version ${version} (in repo ${REPO_NAME})"
				fi
				return 0
                ;;

            -h|--help|help)
                usage $0 "$_usage_string"
                return 0
                ;;

            -f|--force)
				_force=1
                shift
                ;;

            --bump|bump)
                shift
                bump "$@"
                return
                ;;

            *)
				dbecho "template default action reached"
                # usage template '[--init|--bump|--force|--version|--help]'
                # return 0
                ;;
        esac


		me "Setup Local Repo Files for (${REPO_NAME})"
			dbinfo "\${REPO_PATH}: ${REPO_PATH}"
			dbinfo "\${REPO_NAME}: ${REPO_NAME}"
			dbinfo "\${LOCAL_USER_PATH}: ${LOCAL_USER_PATH}"
			dbinfo "\$1: $1"
			dbinfo "\$@: $@"

        if [ -z "$1" ]; then
			#*************** handle default repo files
				# Repo root-level directories to ignore
					#! .git		Git repo - never copy this
					#! bak		Repo specific - never copy this

				# repo folders to copy
					#* .github	GitHub actions and workflows - use these
					#/ cmd		Examples - may copy or recreate (needs renaming in parts)
					#* docs		Copy initial GitHub Pages site

				# generated later (not copied)
					#! VERSION
					#! coverage.txt
					#! go.doc
					#! go.mod

				#/ (this is the list for the 'template' function)
				# repo named file (special case: must be renamed)
				# rename after processing ...
					#/ gorepotemplate.go

				# All repo root-level directories
					#* .editorconfig
					#* .gitignore
					#* CODE_OF_CONDUCT.md
					#* LICENSE
					#* README.md
					#* SECURITY.md
					#* contributing.md
					#* example.go
					#* go.test.sh needs chmod +x (if cp -a doesn't put it??)
					#* idea.md

					#/ files=( .editorconfig .gitignore CODE_OF_CONDUCT.md LICENSE README.md SECURITY.md contributing.md example.go gorepotemplate.go go.test.sh idea.md )
				files=( .editorconfig .gitignore CODE_OF_CONDUCT.md LICENSE README.md SECURITY.md contributing.md example.go gorepotemplate.go go.test.sh idea.md cmd/example/gorepotemplate/main.go docs/* )
				dbinfo "\$files: $files"
        else
			#*************** handle cli arg repo files
			files=$@
			dbinfo "\$files: $files"
           	# usage $0 '[files]'
			# return 1
        fi

        for f in $files; do
			local src=${LOCAL_TEMPLATE_PATH}/${f}
			local dst=${REPO_PATH}/${REPO_NAME}/${f}
			cp -ab $src $dst
			dbecho cp -ab $src $dst

            clean_template_name $dst
            dbecho clean_template_name $dst
        done;
		[ -e gorepotemplate.go ] && mv -b gorepotemplate.go ${REPO_NAME}.go

	}

#? -----------------------------> gomake setup
    _setup_environment() {
        dbecho "Setup Environment"

        #/ ******* test setup
            #/ if the current directory is gomake_test, a special version
            #/ of this script will run for testing purposes
            #/ if PWD == gomake_test, clear test directory and remake everything ...
            if [[ ${PWD##*/} = "gomake_test" ]]; then
                warn "gomake_test directory found ... entering test mode."
                _DEV_GOMAKE=1
                cd ~
                rm -rf ~/gomake_test
                mkdir ~/gomake_test
                cd ~/gomake_test
            fi
        #/ ******* end test setup

        #* gh must be authenticated to use this script.
		_gh_auth_username
		if [[ -z ${_gh_user} ]]; then
			attn "gh must be authenticated to use this script."
            exists gh || ( gh --help; return 1; )
			gh auth login
            (( $? )) && ( attn "error running 'gh auth login' ... check https://cli.github.com/ "; return 1; )
            _gh_auth_username
            [[ -z ${_gh_user} ]] && ( attn "cannot log in to GitHub"; return 1; )
		fi

        dbinfo "\$_gh_user (GitHub authenticated user): ${_gh_user}"
        dbinfo "\${PWD} (current directory): ${PWD}"
    }
	_setup_environment

    _setup_variables() {
		_setup_environment
        #* set repo variables
		#* general information
			YEAR=$( date +'%Y'; )
            dbinfo "\${YEAR}: ${YEAR}"

		#* local repo information
			REPO_PATH="${PWD%/*}"
			REPO_NAME="${PWD##*/}"

			# "/Users/skeptycal/go/src/github.com/skeptycal/gorepotemplate"
            LOCAL_USER_PATH="${GOPATH}/src/github.com/$(whoami)"
            LOCAL_TEMPLATE_PATH="${LOCAL_USER_PATH}/gorepotemplate"

            EXAMPLE_PATH="cmd/example/${REPO_NAME}"
            EXAMPLE_FILE="${EXAMPLE_PATH}/main.go"

            dbecho "Setup Local Repo (${REPO_NAME})"
            dbinfo "\${REPO_PATH}: ${REPO_PATH}"
            dbinfo "\${REPO_NAME}: ${REPO_NAME}"
            dbinfo "\${LOCAL_USER_PATH}: ${LOCAL_USER_PATH}"
            dbinfo "\${LOCAL_TEMPLATE_PATH}: ${LOCAL_TEMPLATE_PATH}"
            dbinfo "\${EXAMPLE_PATH}: ${EXAMPLE_PATH}"
            dbinfo "\${EXAMPLE_FILE}: ${EXAMPLE_FILE}"

		#* github repo information
            GITHUB_TEMPLATE_PATH="https://github.com/skeptycal/gorepotemplate"
            GITHUB_URL="https://github.com/${_gh_user}"
            GITHUB_REPO_URL="${GITHUB_URL}/${REPO_NAME}"
            GO_GET_URL="github.com/${_gh_user}/${REPO_NAME}"
            GITHUB_DOCS_URL="${GITHUB_REPO_URL}/docs"
            PAGES_URL="https://${_gh_user}.github.io/${REPO_NAME}"

            dbecho "Setup Remote Repo (${REPO_NAME})"
            dbinfo "\${GITHUB_TEMPLATE_PATH}: ${GITHUB_TEMPLATE_PATH}"
            dbinfo "\${_gh_user}: ${_gh_user}"
            dbinfo "\${GITHUB_URL}: ${GITHUB_URL}"
            dbinfo "\${GITHUB_REPO_URL}: ${GITHUB_REPO_URL}"
            dbinfo "\${GO_GET_URL}: ${GO_GET_URL}"
            dbinfo "\${GITHUB_DOCS_URL}: ${GITHUB_DOCS_URL}"
            dbinfo "\${PAGES_URL}: ${PAGES_URL}"

        #* file header blurbs
			BLURB_GO=$( _file_blurb )
			BLURB_INI=$( _file_blurb '#' )

            dbinfo "\${BLURB_GO}: ${BLURB_GO}"
            dbinfo "\${BLURB_INI}: ${BLURB_INI}"
    }

	mkd() {
		if [[ -n "$1" ]]; then
            mkdir -p "$1" >/dev/null 2>&1 || ( warn "error creating directory $1"; return 1 )
            cd "$1" || ( warn "error creating directory $1"; return 1 )
        fi
		return 0
	}

    _setup_local() {
        dbecho "Setup local repo"

        #* check and setup local repo directory
        # create if needed and CD if possible
        if [[ -n "$1" ]]; then
            mkdir -p "$1" >/dev/null 2>&1
            cd "$1" || ( warn "error creating directory $1"; return 1 )
        fi

		mkd "$1" || return 1

        # directory must be empty (certain parts of this setup can be run on existing repos)
        [ -n "$(ls -A ${PWD})" ] && ( warn "directory not empty"; return 1; )

    	#* Initial repo setup
			git init
            dbinfo "\$?: $? - git init"

		return 0
    }
    _setup_remote() {
        #* create remote repo from template (I use GitHub ... change it if you want)
            dbecho "gh repo create ${REPO_NAME} -y --public --template $GITHUB_TEMPLATE_PATH"
            gh repo create ${REPO_NAME} -y --public --template $GITHUB_TEMPLATE_PATH

            if (( $? )); then
                warn "error creating GitHub remote repo ${REPO_NAME}"
                is_prod && return 1; # dev mode may continue with existing repo
            fi

            # this should be done by gh ... but it doesn't always work
            dbecho git remote add origin "${GITHUB_REPO_URL}"
            git remote add origin "${GITHUB_REPO_URL}" >/dev/null 2>&1

            if (( $? )); then
                warn "error adding remote repository";
                is_prod && return 1; # dev mode may continue with existing repo
            fi

            # in the case of using an existing remote ... dev only
            dbecho git pull origin main --rebase
            git pull origin main --rebase >/dev/null 2>&1

            if (( $? )); then
                warn "error syncing remote repository"
                is_prod && return 1; # dev mode may continue with existing repo
            fi

        #* .gitignore and initial commit
            makeGI
            git add .gitignore -f
            git commit -m "initial commit"

            if (( $? )); then
                warn "error with initial commit";
                is_prod && return 1; # dev mode may continue with existing repo
            fi

            # push initial repository changes
            dbecho git push --set-upstream origin main
            git push --set-upstream origin main


            if (( $? )); then
                warn "error with initial remote repo push";
                is_prod && return 1; # dev mode may continue with existing repo
            fi
    }
    _setup_dirs() {
        # based on the unofficial and evolving https://github.com/golang-standards/project-layout

        # remove template placeholder example files
        rm -rf go.sum
        rm -rf go.mod
        rm -rf go.doc
        rm -rf gorepotemplate.go


        mkdir -p "$EXAMPLE_PATH"
		cp ${LOCAL_TEMPLATE_PATH}/cmd/example/gorepotemplate/main.go "$EXAMPLE_FILE"
		clean_template_name "$EXAMPLE_FILE"
		rm -rf cmd/example/gorepotemplate

        #* .gitignore and initial commit
        git add --all && git commit -m 'GoBot: setup directory tree and remove template examples'
        git push origin main
    }
    _make_files() {
		#* GitHub repo files
        GO_VERSION=$(go_version)

        # the default list of files is for template is:
            # files=( .editorconfig .gitignore CODE_OF_CONDUCT.md LICENSE README.md SECURITY.md contributing.md example.go gorepotemplate.go go.test.sh idea.md cmd/example/gorepotemplate/main.go docs/* )
        template

        git add --all && git commit -m "GoBot: setup repo files from template"

		#* GitHub Pages site setup
			mkdir -p docs
			git checkout -b gh-pages
            git fetch
			git push origin gh-pages

            git checkout main

            git add --all && git commit -m "GoBot: create GitHub Pages folder branch"

        #* dev branch and initial dev version
            git checkout -b dev
            git fetch
            git push origin dev

            git checkout main

			# comes from template now
            # mkdir -p .github
            # cd .github
            # template .github/*
            # cd ISSUE_TEMPLATE
            # template .github/ISSUE_TEMPLATE/*
            # cd ..
            # cd workflows
            # template .github/workflows/*
            # cd ${REPO_PATH}/${REPO_NAME}

            git add --all && devtag "GoBot: dev branch and directory tree setup"

		    #* Go module setup
            go mod init
            go get -u "${GO_GET_URL}"
            go mod tidy

            git add --all && git commit -m "GoBot: Go module setup"

            go doc >|go.doc
            chmod +x go.test.sh
            ./go.test.sh

            git add --all && git commit -m "GoBot: initial docs and test run"

    }

#? -----------------------------> gomake main
    _gomake() {
        attn "GoMake make_private is set to $make_private. The GitHub repository created will be private."
        _setup_variables "$@"
        _setup_environment "$@"
        _setup_local "$@"
        _setup_remote "$@"
        _setup_dirs "$@"
        _make_files "$@"
        devtag
    }
    gomake_menu() {
		dbinfo "args: $@"
        case "$1" in

            -v|--version|version)
				dbinfo " $0 $@"
                echo "gomake $version"
                return 0
                ;;

            -h|--help|help)
                echo "Usage: gomake [reponame] [--files] [--up]"
                return 0
                ;;

            --files|files)
                shift
                _make_files "$@"
                return
                ;;

            --devtag|devtag)
                shift
                devtag "$@"
                return
                ;;

            --bump|bump)
                shift
                version "$@"
                return
                ;;

            init)
                shift
                if [[ $1 == 'private' ]]; then
                    make_private="private"
                    shift
                fi
                _gomake "$@"
                ;;

            *)
                usage gomake "$_gomake_usage_string"
                return 0
                ;;
        esac
    }

	#!------------------------> main
	#     #! repo testing ...
	#     # alias streamtest='cdgo; del stream; mkd stream; gomake'
	#     version=$(_get_version)
	#     make_private='public'
	#     _gomake_usage_string="[init [private]|files|devtag|bump|help|version]"
	#     _setup_variables
	# SET_DEBUG=0
	# 	dbecho "run gomake_menu ..."
	#     gomake_menu "$@"

	# 	gomake() {
	# 		. $(which gomake.sh) "$@"
	# 	}

#? -----------------------------> submodule update
	update_repo() {
		msg=${1:-"GitBot: auto updates"}
		git add -A
		git commit -m "$msg"
		$(devtag)
		# git push --set-upstream origin $(git_current_branch)
	}

	ws_update () {
		# update all submodules that do not begin with '.'
		for d in $(dirlist); do
			cd $d
			update_repo "GoBot: submodule updates"
			git submodule add ./$d
			cd -
		done

		# update workspace management utilities repo
		cd .manage
		update_repo "update workspace management repo"
		cd -

		# update workspace repo
		update_repo "update Go utilities workspace repo"
	}

	_foreach() {
		eval "$@"
		true
		return 0
	}

	gitforeach() {
		git submodule foreach $(_foreach "$@")
	}

#!------------------------> main
ws_update 2>&1 /dev/null
