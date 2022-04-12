#!/bin/zsh

function whichprog {
	type -p type |cut -d ' ' -f 3
}

function pause(){
 /usr/bin/read -s -n 1 -p "Press any key to continue . . ."
 echo ""
}

REPO="https://github.com/skeptycal/goutil.git"

echo "\$SHELL: ${SHELL}"
echo "\$REPO: ${REPO}"

echo "create git submodules from all directories within the current working directory ..."
echo "press CTRL+C to exit ..."
pause

function create_git_submodules() {
	[[ -n "$1" ]] && cd "$1"
	echo "\$PWD: $PWD"
	for f in $(dirlist); do
		echo "dir: $f"
		echo git submodule add ./$f
	done;
	[[ -n "$1" ]] && cd -
}

function recurse_git_sub_add() {

}


# add all folders (non hidden) to go.work
# dirlist |xargs go work use ./{}

# add all folders (non hidden) to git submodules
# TODO something doesn't work here ...
# dirlist |xargs git submodule add $REPO {}

create_git_submodules .
