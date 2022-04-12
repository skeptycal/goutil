#!/bin/zsh

exists() { command -v "$1" >/dev/null 2>&1; }
require() {
    MESSAGE=${2:="$1 is not installed.  Aborting."}
    command -v "$1" >/dev/null 2>&1 || { echo >&2 $MESSAGE; exit 1; }
}

die () {
    MESSAGE=${1:="An error occurred.  Aborting."}
    { echo >&2 $MESSAGE; exit 1; }
}

first_field() {
    # split lines using tab and save first field
    # cat repos.csv | cut -d $'\t' -f 1 | sort >|templist.csv
    #  ... ok ... tab is the default so no need for delimiter
    [ -e "$1" ] || { die "File not found ($1). Aborting."; }
    cat "$1" | cut -f 1 | sort
}

require gh

REPO_LIMIT=2500
DEFAULT_ORG="skeptycal"

ORG=${1:=$DEFAULT_ORG}

echo "GitHub Repo Management"
echo ""
gh --version
echo ""
echo "Getting github repo info for org '${ORG}':"
echo ""

# get up to 2500 repos
# gh repo list skeptycal -L 2500 | sort >|repos.csv
echo "Getting list of forks... "

gh repo list --fork -L $REPO_LIMIT | sort >|forks.csv
wc -l forks.csv
echo ""

echo "Getting list of source repos... "
gh repo list --source -L $REPO_LIMIT | sort >|sources.csv
wc -l sources.csv
echo ""

first_field forks.csv >|"forks_list.csv"
first_field sources.csv >|"sources_list.csv"
echo ""
