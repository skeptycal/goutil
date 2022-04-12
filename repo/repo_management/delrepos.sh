#!/bin/zsh

DEFAULT_ORG="skeptycal"

#* GitHub repository naming conventions:
    # Repository names
        # Max length: 100 code points
        # All code points must be either a hyphen (-), an underscore (_), a period (.), or an ASCII alphanumeric code point
        # Must be unique per-user and/or per-organization
        # Note: sequences of invalid code points are automatically replaced by a single hyphen (-) Note: length checking is performed after replacement
        #
    # This was verified through checking automatically-generated aliases with repository names.
    #
    # Reference: https://github.com/isiahmeadows/github-limits

# Delete a GitHub repository.
    # With no argument, deletes the current repository.
    # (in this function, blank repo ($1) is not allowed.)
    # Otherwise, deletes the specified repository.
    #
    # Deletion requires authorization with the "delete_repo" scope.
    # To authorize, run "gh auth refresh -s delete_repo"
delrepo() {
    [[ -z "$1" ]] && return 1
    gh repo delete "$DEFAULT_ORG/$1" --confirm
}