#!/usr/bin/zsh

update_git_dirs () {
	TEMPLATE_DIR=~$GOPATH/src/github.com/$(whoami)/gorepotemplate
	EXCLUDES='.virtualenvs node_modules .venv .git'
	SEARCH_PATH=$PWD
	# cd "$HOME" || return
	lime "Locating all git repos in $SEARCH_PATH" ..." >&6
	git_dirs=$(find . -type d -name ".virtualenvs" -prune -o -name ".git" | sed 's/\.git//')
	green ${git_dirs//.\//\\n}
	for i in git_dirs
	do
		attn "Going into $i" >&6
		cd "$i" || return
		gitit "Gitit Bot: weekly update - minor / formatting" >&6
		git stash >&6
		git pull origin master --rebase >&6
		git stash apply >&6
		[ -f .pre-commit-config.yaml ] || cp $TEMPLATE_DIR/.pre-commit-config.yaml .
		pre-commit autoupdate
		gitit >&6
		cd ~
	done
	cd "$SEARCH_PATH"
} 6>&1 > /dev/null 2>&1
