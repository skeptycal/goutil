# gorepotemplate

> This is the initial directory tree for gorepotemplate. Use the make_tree_md.sh script ([GNU-tree required][get_tree]) to update it if you wish. It is safe to delete this file.

### Directory Structure

```sh
.
├── .editorconfig
├── .github
│   ├── ISSUE_TEMPLATE
│   │   ├── bug_report.md
│   │   └── feature_request.md
│   └── workflows
│       └── go.yml
├── .gitignore
├── .vscode
│   └── settings.json
├── CODE_OF_CONDUCT.md
├── LICENSE
├── README.md
├── SECURITY.md
├── contributing.md
├── generic
│   ├── .gitmodules
│   ├── generictypes
│   │   ├── asserts.go
│   │   ├── bubblesort_test.go
│   │   ├── constraints
│   │   │   ├── constraints.go
│   │   │   └── constraints_test.go
│   │   ├── dict
│   │   │   ├── components.go
│   │   │   ├── dict.go
│   │   │   ├── dict_test.go
│   │   │   └── example
│   │   │       └── main.go
│   │   ├── examples.go
│   │   ├── generic.go
│   │   ├── generic_test.go
│   │   ├── kinds.go
│   │   ├── list
│   │   │   ├── bubblesort.go
│   │   │   ├── convert.go
│   │   │   └── list.go
│   │   ├── sampletypes.go
│   │   ├── sequence
│   │   │   ├── examples
│   │   │   │   └── main.go
│   │   │   ├── examples.go
│   │   │   ├── sequence.go
│   │   │   └── sequence_test.go
│   │   ├── so_answer1.md
│   │   ├── so_questions.go
│   │   ├── soexamples
│   │   │   └── 71677581
│   │   │       ├── .editorconfig
│   │   │       ├── .github
│   │   │       │   ├── FUNDING.yml
│   │   │       │   ├── ISSUE_TEMPLATE
│   │   │       │   │   ├── bug_report.md
│   │   │       │   │   └── feature_request.md
│   │   │       │   ├── dependabot.yml
│   │   │       │   └── workflows
│   │   │       │       ├── ci.yml
│   │   │       │       └── go.yml
│   │   │       ├── CODE_OF_CONDUCT.md
│   │   │       ├── LICENSE
│   │   │       ├── README.md
│   │   │       ├── SECURITY.md
│   │   │       ├── contributing.md
│   │   │       ├── coverage.txt
│   │   │       ├── go.test.sh
│   │   │       ├── main.go
│   │   │       ├── main.go.bak
│   │   │       ├── main_test.go
│   │   │       └── profile.out
│   │   ├── soexamples.bak
│   │   │   ├── main.go
│   │   │   ├── main.go.bak
│   │   │   └── main_test.go
│   │   ├── sort
│   │   │   ├── floats.go
│   │   │   ├── generic.go
│   │   │   ├── ints.go
│   │   │   ├── sort.go
│   │   │   ├── stable.go
│   │   │   ├── stdlib.go
│   │   │   ├── stdlibinternal.go
│   │   │   └── strings.go
│   │   ├── stack.go
│   │   ├── stack_test.go
│   │   └── types.go
│   ├── main.go
│   ├── main_test.go
│   ├── manage
│   │   ├── .testoutput
│   │   │   └── heap.txt
│   │   ├── data
│   │   │   ├── aliases.txt
│   │   │   ├── coverage.txt
│   │   │   ├── env.txt
│   │   │   └── gh_repo_create_help.txt
│   │   ├── examples
│   │   │   ├── escapes
│   │   │   │   └── main.go
│   │   │   └── gitit
│   │   │       ├── main.go
│   │   │       ├── main_test.go
│   │   │       └── temp
│   │   ├── ghshell
│   │   │   ├── .testoutput
│   │   │   │   └── heap.txt
│   │   │   ├── cmd.go
│   │   │   ├── errorcontrol.go
│   │   │   ├── errors.go
│   │   │   ├── ghshell.go
│   │   │   ├── ghshell_test.go
│   │   │   ├── gitit.go
│   │   │   ├── gobot.go
│   │   │   ├── initgobot.go
│   │   │   ├── logging.go
│   │   │   ├── rand.go
│   │   │   ├── unix_rusage.1
│   │   │   ├── util.go
│   │   │   └── util_test.go
│   │   ├── gnuflags
│   │   │   └── flags.go
│   │   ├── main.go
│   │   ├── profile.out
│   │   └── shellscripts
│   │       ├── gitit.sh
│   │       ├── gitsub.sh
│   │       ├── gitutil.sh
│   │       ├── go.test.sh
│   │       ├── gobuildflagsoutput.sh
│   │       ├── update.sh
│   │       ├── updatesubs.sh
│   │       └── workspace_init.sh
│   └── nongeneric
│       └── nongeneric.go
├── go.doc
├── go.mod
├── go.sum
├── go.test.sh
├── make_tree_md.sh
├── os
│   ├── .gitmodules
│   ├── basicfile
│   │   ├── basicfile.go
│   │   ├── datafile.go
│   │   ├── direntry.go
│   │   ├── errors.go
│   │   ├── fileinfo.go
│   │   ├── fileinfoaliases.go
│   │   ├── filemode.go
│   │   ├── fileops.go
│   │   ├── fileunix.go
│   │   ├── gofile.go
│   │   ├── gofileerror.go
│   │   ├── internal.go
│   │   ├── llrb_avg.go
│   │   ├── llrb_avg_test.go
│   │   ├── textfile.go
│   │   ├── types.go
│   │   ├── util.go
│   │   └── ~bench_results.csv
│   ├── gofile
│   │   ├── .editorconfig
│   │   ├── .github
│   │   │   ├── FUNDING.yml
│   │   │   ├── ISSUE_TEMPLATE
│   │   │   │   ├── bug_report.md
│   │   │   │   └── feature_request.md
│   │   │   ├── dependabot.yml
│   │   │   └── workflows
│   │   │       ├── ci.yml
│   │   │       └── go.yml
│   │   ├── CODE_OF_CONDUCT.md
│   │   ├── LICENSE
│   │   ├── README.md
│   │   ├── SECURITY.md
│   │   ├── VERSION
│   │   ├── cmd
│   │   │   ├── dir
│   │   │   │   ├── main.go
│   │   │   │   ├── main_test.go
│   │   │   │   └── temp1
│   │   │   │       ├── temp
│   │   │   │       └── temp2
│   │   │   │           └── temp
│   │   │   └── examples
│   │   │       └── errorlog
│   │   │           └── main.go
│   │   ├── constants.go
│   │   ├── contributing.md
│   │   ├── copy.go
│   │   ├── copybenchmarks
│   │   │   ├── copy_test.go
│   │   │   ├── fakeDst
│   │   │   └── fakeSrc
│   │   ├── coverage.txt
│   │   ├── dir_options.go
│   │   ├── dirlist.go
│   │   ├── errors.go
│   │   ├── fileops.go
│   │   ├── go.test.sh
│   │   ├── gofile.go
│   │   ├── gofileerror.go
│   │   ├── gofileerror_test.go
│   │   ├── internal.go
│   │   ├── internal_test.go
│   │   ├── logging.go
│   │   ├── profile.out
│   │   └── types.go
│   ├── gofile2
│   │   ├── .editorconfig
│   │   ├── .github
│   │   │   ├── FUNDING.yml
│   │   │   ├── ISSUE_TEMPLATE
│   │   │   │   ├── bug_report.md
│   │   │   │   └── feature_request.md
│   │   │   ├── dependabot.yml
│   │   │   └── workflows
│   │   │       ├── ci.yml
│   │   │       └── go.yml
│   │   ├── CODE_OF_CONDUCT.md
│   │   ├── LICENSE
│   │   ├── README.md
│   │   ├── SECURITY.md
│   │   ├── VERSION
│   │   ├── basicfile.go
│   │   ├── cmd
│   │   │   ├── dir
│   │   │   │   ├── main.go
│   │   │   │   └── main_test.go
│   │   │   └── examples
│   │   │       └── errorlog
│   │   │           └── main.go
│   │   ├── constants.go
│   │   ├── contributing.md
│   │   ├── copy.go
│   │   ├── copybenchmarks
│   │   │   ├── copy_test.go
│   │   │   ├── fakeDst
│   │   │   └── fakeSrc
│   │   ├── coverage.txt
│   │   ├── dir_options.go
│   │   ├── dirlist.go
│   │   ├── exported.go
│   │   ├── fakeDst
│   │   ├── fakeSrc
│   │   ├── file_error.go
│   │   ├── file_error_test.go
│   │   ├── fileops.go
│   │   ├── fileops_test.go
│   │   ├── go.test.sh
│   │   ├── gofile.go
│   │   ├── internal.go
│   │   ├── internal_test.go
│   │   ├── logerrors.go
│   │   └── types.go
│   ├── goshell
│   │   ├── .editorconfig
│   │   ├── .github
│   │   │   ├── FUNDING.yml
│   │   │   ├── ISSUE_TEMPLATE
│   │   │   │   ├── bug_report.md
│   │   │   │   └── feature_request.md
│   │   │   ├── dependabot.yml
│   │   │   └── workflows
│   │   │       ├── codeql-analysis.yml
│   │   │       └── go.yml
│   │   ├── .pre-commit-config.yaml
│   │   ├── CODE_OF_CONDUCT.md
│   │   ├── LICENSE
│   │   ├── README.md
│   │   ├── SECURITY.md
│   │   ├── cmd
│   │   │   └── example
│   │   │       └── goshell
│   │   │           └── main.go
│   │   ├── contributing.md
│   │   ├── coverage.txt
│   │   ├── defaults.go
│   │   ├── defaults_test.go
│   │   ├── docs
│   │   │   ├── _config.yml
│   │   │   ├── docs.md
│   │   │   ├── index.html
│   │   │   └── template.md
│   │   ├── example.go
│   │   ├── go.test.sh
│   │   ├── homedir.go
│   │   ├── homedir_test.go
│   │   ├── idea.md
│   │   ├── internal
│   │   │   └── fixture
│   │   │       └── test.go
│   │   ├── make_tree_md.sh
│   │   └── tree.md
│   ├── ls
│   │   ├── .github
│   │   │   ├── FUNDING.yml
│   │   │   ├── ISSUE_TEMPLATE
│   │   │   │   ├── bug_report.md
│   │   │   │   └── feature_request.md
│   │   │   ├── dependabot.yml
│   │   │   └── workflows
│   │   │       ├── ci.yml
│   │   │       └── go.yml
│   │   ├── CODE_OF_CONDUCT.md
│   │   ├── LICENSE
│   │   ├── README.md
│   │   ├── SECURITY.md
│   │   ├── ansi.go
│   │   ├── contributing.md
│   │   ├── git.go
│   │   ├── ls.go
│   │   ├── ls.txt
│   │   └── main.go
│   ├── memfile
│   │   ├── LICENSE
│   │   ├── memfile.go
│   │   ├── memfile_test.go
│   │   ├── uniuri.go
│   │   ├── util.go
│   │   └── util_test.go
│   ├── osargsutils
│   │   ├── executable.go
│   │   └── executable_test.go
│   ├── path
│   │   ├── cmd
│   │   │   └── path.go
│   │   └── shpath.go
│   ├── redlogger
│   │   ├── redlog
│   │   │   └── main.go
│   │   └── redlogger.go
│   └── shpath
│       ├── cmd
│       │   └── path.go
│       ├── osutil.go
│       ├── shpath.go
│       ├── shpath_test.go
│       ├── stringlists.go
│       └── stringutil.go
├── repo
│   ├── defaults
│   │   ├── .editorconfig
│   │   ├── .github
│   │   │   ├── FUNDING.yml
│   │   │   ├── ISSUE_TEMPLATE
│   │   │   │   ├── bug_report.md
│   │   │   │   └── feature_request.md
│   │   │   ├── dependabot.yml
│   │   │   └── workflows
│   │   │       ├── codeql-analysis.yml
│   │   │       └── go.yml
│   │   ├── .pre-commit-config.yaml
│   │   ├── CODE_OF_CONDUCT.md
│   │   ├── LICENSE
│   │   ├── README.md
│   │   ├── SECURITY.md
│   │   ├── ansi.go
│   │   ├── cmd
│   │   │   ├── benchmarks
│   │   │   │   └── formatbenchmarks
│   │   │   │       ├── main.go
│   │   │   │       └── main_test.go
│   │   │   └── example
│   │   │       ├── color_example
│   │   │       │   └── main.go
│   │   │       └── defaults
│   │   │           └── main.go
│   │   ├── contributing.md
│   │   ├── coverage.txt
│   │   ├── debug.go
│   │   ├── debug_test.go
│   │   ├── defaults.go
│   │   ├── defaults_test.go
│   │   ├── docs
│   │   │   ├── _config.yml
│   │   │   ├── docs.md
│   │   │   ├── index.html
│   │   │   └── template.md
│   │   ├── example.go
│   │   ├── flags.go
│   │   ├── go.test.sh
│   │   ├── idea.md
│   │   ├── internal.go
│   │   ├── make_tree_md.sh
│   │   ├── settings.go
│   │   ├── test_utils.go
│   │   ├── test_utils_test.go
│   │   ├── trace.go
│   │   ├── tree.md
│   │   ├── types.go
│   │   ├── types_test.go
│   │   ├── utils.go
│   │   └── utils_test.go
│   ├── errorlogger
│   │   ├── .editorconfig
│   │   ├── .github
│   │   │   ├── FUNDING.yml
│   │   │   ├── ISSUE_TEMPLATE
│   │   │   │   ├── bug_report.md
│   │   │   │   └── feature_request.md
│   │   │   └── workflows
│   │   │       └── go.yml
│   │   ├── CODE_OF_CONDUCT.md
│   │   ├── LICENSE
│   │   ├── README.md
│   │   ├── SECURITY.md
│   │   ├── cmd
│   │   │   └── example
│   │   │       └── executable
│   │   │           └── main.go
│   │   ├── contributing.md
│   │   ├── coverage.txt
│   │   ├── docs
│   │   │   ├── _config.yml
│   │   │   ├── docs.md
│   │   │   ├── index.html
│   │   │   └── template.md
│   │   ├── error_func.go
│   │   ├── error_func_test.go
│   │   ├── errorlogger.go
│   │   ├── errorlogger_test.go
│   │   ├── example.go
│   │   ├── go.test.sh
│   │   ├── idea.md
│   │   ├── json_formatter.go
│   │   ├── json_formatter_test.go
│   │   ├── level.go
│   │   ├── logrus_types.go
│   │   ├── profile.out
│   │   ├── test_info.go
│   │   ├── test_info_test.go
│   │   ├── text_formatter.go
│   │   └── text_formatter_test.go
│   ├── fake_repo
│   │   ├── LICENSE
│   │   └── fakerepo.go
│   ├── ghrepo
│   ├── gitcommits
│   │   ├── .editorconfig
│   │   ├── .github
│   │   │   ├── FUNDING.yml
│   │   │   ├── ISSUE_TEMPLATE
│   │   │   │   ├── bug_report.md
│   │   │   │   └── feature_request.md
│   │   │   ├── dependabot.yml
│   │   │   └── workflows
│   │   │       ├── codeql-analysis.yml
│   │   │       └── go.yml
│   │   ├── .pre-commit-config.yaml
│   │   ├── CODE_OF_CONDUCT.md
│   │   ├── LICENSE
│   │   ├── README.md
│   │   ├── SECURITY.md
│   │   ├── cmd
│   │   │   └── example
│   │   │       └── gitcommits
│   │   │           └── gitcommits.go
│   │   ├── contributing.md
│   │   ├── coverage.txt
│   │   ├── docs
│   │   │   ├── _config.yml
│   │   │   ├── docs.md
│   │   │   ├── index.html
│   │   │   └── template.md
│   │   ├── example.go
│   │   ├── gitcommits.go
│   │   ├── go.test.sh
│   │   ├── idea.md
│   │   ├── make_tree_md.sh
│   │   └── tree.md
│   ├── gitroot
│   │   ├── find
│   │   │   ├── find.go
│   │   │   └── find_test.go
│   │   └── root.go
│   ├── goconfig
│   │   ├── flags.go
│   │   ├── goconfig.go
│   │   └── goconfig_test.go
│   ├── gogithub
│   │   ├── .editorconfig
│   │   ├── .github
│   │   │   ├── FUNDING.yml
│   │   │   ├── ISSUE_TEMPLATE
│   │   │   │   ├── bug_report.md
│   │   │   │   └── feature_request.md
│   │   │   ├── dependabot.yml
│   │   │   └── workflows
│   │   │       ├── codeql-analysis.yml
│   │   │       └── go.yml
│   │   ├── .pre-commit-config.yaml
│   │   ├── CODE_OF_CONDUCT.md
│   │   ├── LICENSE
│   │   ├── README.md
│   │   ├── SECURITY.md
│   │   ├── auth.go
│   │   ├── builderpool.go
│   │   ├── cmd
│   │   │   └── example
│   │   │       └── gogithub
│   │   │           └── example.go
│   │   ├── contributing.md
│   │   ├── coverage.txt
│   │   ├── docs
│   │   │   ├── _config.yml
│   │   │   ├── docs.md
│   │   │   ├── index.html
│   │   │   └── template.md
│   │   ├── example.go
│   │   ├── exec.go
│   │   ├── exec_test.go
│   │   ├── go.test.sh
│   │   ├── gogithub.go
│   │   ├── idea.md
│   │   ├── job.go
│   │   ├── job_test.go
│   │   ├── make_tree_md.sh
│   │   ├── sbWriter.go
│   │   ├── sbpool.go
│   │   └── tree.md
│   ├── gomake
│   │   ├── datafile.go
│   │   ├── example.go
│   │   ├── gi_list
│   │   ├── gogithub.go
│   │   ├── gomake.go
│   │   ├── internal.go
│   │   ├── pprint.go
│   │   ├── template_files
│   │   │   ├── README.md
│   │   │   ├── index.html
│   │   │   └── screenshot.css
│   │   ├── templates.go
│   │   └── util.go
│   ├── gomod
│   │   ├── .gitmodules
│   │   ├── checks
│   │   │   └── checks.go
│   │   ├── gen
│   │   │   └── genzfunc.go
│   │   ├── main.go
│   │   └── mod
│   │       ├── mod.go
│   │       └── mod_test.go
│   ├── repo_management
│   │   ├── cleanlist
│   │   │   └── cleanlist.go
│   │   ├── config
│   │   │   ├── config.go
│   │   │   ├── flags.go
│   │   │   ├── syncmap.go
│   │   │   └── util.go
│   │   ├── countargs.sh
│   │   ├── del_list.sh
│   │   ├── delrepos.sh
│   │   ├── forks.csv
│   │   ├── forks_list.csv
│   │   ├── ghnamecheck
│   │   │   ├── ghnamecheck_test.go
│   │   │   └── main.go
│   │   ├── make_repo_list.sh
│   │   ├── sources.csv
│   │   └── sources_list.csv
│   ├── seeker
│   │   ├── cli.go
│   │   ├── cmd
│   │   │   └── server
│   │   │       ├── serve.go
│   │   │       └── serve_test.go
│   │   ├── config.go
│   │   ├── handlers.go
│   │   ├── http.go
│   │   ├── mapper.go
│   │   ├── seeker.go
│   │   ├── seeker.sh
│   │   ├── time.go
│   │   ├── timeMap.go
│   │   ├── timemapper.go
│   │   └── util.go
│   └── util2
│       ├── .github
│       │   └── workflows
│       │       ├── codeql-analysis.yml
│       │       ├── go.yml
│       │       ├── greetings.yml
│       │       ├── label.yml
│       │       └── stale.yml
│       ├── LICENSE
│       ├── datatools
│       │   ├── cmd
│       │   │   └── EmailDomains
│       │   │       ├── emaildomains
│       │   │       └── main.go
│       │   ├── compare
│       │   │   └── interface.go
│       │   ├── format
│       │   │   ├── email.go
│       │   │   ├── email_test.go
│       │   │   ├── numberformatting.go
│       │   │   ├── numberformatting_test.go
│       │   │   └── sample.txt
│       │   ├── math
│       │   │   └── polynomial
│       │   │       ├── cmd
│       │   │       │   └── poly
│       │   │       │       ├── main.go
│       │   │       │       └── poly
│       │   │       ├── polynomial.go
│       │   │       └── polynomial_test.go
│       │   └── mysql
│       │       ├── LICENSE
│       │       ├── config.go
│       │       ├── license.go
│       │       └── mysql.go
│       ├── devtools
│       │   ├── cli
│       │   │   └── main.go
│       │   ├── gogit
│       │   │   ├── _example
│       │   │   │   ├── cli_examples
│       │   │   │   │   └── github_api_response.json
│       │   │   │   └── escape-seq
│       │   │   │       └── main.go
│       │   │   ├── gogit.go
│       │   │   └── gogit_test.go
│       │   └── testing
│       │       └── testing.go
│       ├── gofile
│       │   ├── cmd
│       │   │   ├── dir
│       │   │   │   ├── main.go
│       │   │   │   └── main_test.go
│       │   │   └── redlog
│       │   │       └── main.go
│       │   ├── fileops.go
│       │   ├── fileops_test.go
│       │   ├── gofile.go
│       │   ├── json
│       │   │   └── json.go
│       │   └── redlogger
│       │       └── redlogger.go
│       ├── template
│       │   ├── LICENSE
│       │   └── license.go
│       └── webtools
│           ├── LICENSE
│           ├── getpage
│           │   └── googlesearch.css
│           ├── http
│           │   └── downloadurl.go
│           └── youtube
│               └── gotube
│                   ├── gotube
│                   └── main.go
├── sequence
│   ├── sequence
│   └── sequence_test.go
├── soexamples
│   └── 71677581
└── tree.md

164 directories, 496 files
```

[get_tree]: (http://mama.indstate.edu/users/ice/tree/)
