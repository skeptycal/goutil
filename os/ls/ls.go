package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	log "
)

const (
	NL  rune = '\n'
	NUL rune = 0
)

// var (
// 	pathFlag, excludeList                    string
// 	sizeFlag, oneFlag, zeroFlag, versionFlag bool
// 	home                                     string = getHome()
// 	pwd                                      string = getPWD()
// 	outputSEP                                string = " "
// )

// func init() {
// 	flag.StringVar(&pathFlag, "root", "", "root path to use for searches")
// 	flag.StringVar(&pathFlag, "r", "", "root path to use for searches (short)")
// 	flag.StringVar(&excludeList, "x", "", "list of excluded patterns")

// 	flag.BoolVar(&versionFlag, "version", false, "version of the package")
// 	flag.BoolVar(&versionFlag, "v", false, "version of the package (short)")
// 	flag.BoolVar(&sizeFlag, "size", false, "print the allocated size of each file, in blocks")
// 	flag.BoolVar(&zeroFlag, "0", false, "end each output line with NUL, not newline")
// 	flag.BoolVar(&oneFlag, "1", false, "list one file per line")

// 	flag.Parse()

// 	// files := flag.Args()[0]
// 	// fmt.Printf("files: %s\n", files)

// 	flag.PrintDefaults()

// 	if versionFlag {
// 		fmt.Printf("Version: %s\n", "fakeVersion 1.0.0")
// 		os.Exit(0)
// 	}

// 	if oneFlag {
// 		outputSEP = string(NL)
// 	}

// 	if zeroFlag {
// 		outputSEP = string(NUL)
// 	}

// 	if len(os.Args) > 2 {
// 		pathFlag = os.Args[1]
// 	} else {
// 		pathFlag = getPWD()
// 	}
// }

func getHome() (d string) {
	d, _ = os.UserHomeDir()
	return
}

func Env(s string) string {
	return os.ExpandEnv(s)
}

func getPWD() string {
	return Env(`$PWD`)
}

func Ls(root, pattern string) (list []os.FileInfo, err error) {
	err = filepath.Walk(root,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			list = append(list, info)
			return nil
		})
	return
}

func DirsRecursive(root, pattern string) (list []os.FileInfo, err error) {
	err = filepath.Walk(root,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				list = append(list, info)
			}
			return nil
		})
	return list, err
}

// ReadNames performs a directory walk to print all file names
//
// using fileinfo is probably slower ...
func ReadNames(path string) (err error) {
	fmt.Println("ReadNames for ", path)
	err = filepath.Walk(path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			fmt.Println(path, info.Size())
			return nil
		})
	if err != nil {
		log.Println(err)
	}
	return
}

// ReadNamesOnly performs a quicker filesearch returning only a slice of names
func ReadNamesOnly(path string) (err error) {
	fmt.Println("ReadNamesOnly for ", path)
	fmt.Println("=========================")
	const maxNames = -1
	file, err := os.Open(path)
	if err != nil {
		return
	}
	names, err := file.Readdirnames(maxNames)
	log.Println(strings.Join(names, "\n"))
	return
}

// func main() {
// 	useColor := flag.Bool("color", false, "display colorized output")
// 	flagAll := flag.Bool("A", false, "almost-all : do not list hidden files")
// 	flagAuthor := flag.Bool("author", false, "with -l, print the author of each file")
// 	flagEscape := flag.Bool("b", false, "print C-style escapes for nongraphic characters")
// 	flagBlockSize := flag.Int("block-size", 1024, "with -l, scale sizes by SIZE when printing them")
// 	flagIgnoreBackups := flag.Bool("B", false, "do not list implied entries ending with ~")
// 	flagColumns := flag.Bool("C", false, "list entries by columns")
// 	flagColor := flag.String("color", "always", "colorize the output; WHEN can be 'always' (default if omitted), 'auto', or 'never'; more info below")
// 	flagDir := flag.Bool("d", false, "list directories themselves, not their contents")
// 	flagClassify := flag.Bool("F", false, "append indicator (one of */=>@|) to entries")
// 	flagDirsFirst := flag.Bool("group-directories-first", false, "group directories before files")
// 	flagHuman := flag.Bool("h", false, "with -l and -s, print sizes like 1K 234M 2G etc.")
// 	flagSi := flag.Bool("si", false, "but use powers of 1000 not 1024")
// 	flag.Parse()
// 	flagFollow := flag.Bool("H", false, "follow symbolic links listed on the command line")
// 	flagHyperlink := flag.String("hyperlink", "always", "hyperlink file names; WHEN can be 'always' (default if omitted), 'auto', or 'never'")
// 	flagInode := flag.Bool("inode", false, "print the index number of each file")
// 	flagIgnorePattern := flag.String("I", "", "do not list implied entries matching shell PATTERN")
// 	flagLong := flag.Bool("l", false, "use a long listing format")
// 	flagIndicator := flag.Bool("p", false, "append / indicator to directories")
// 	flagReverse := flag.Bool("r", false, "reverse order while sorting")
// 	flagRecursive := flag.Bool("R", false, "list subdirectories recursively")
// 	flagSortSize := flag.Bool("S", false, "sort by file size, largest first")
// 	flagSortTime := flag.Bool("t", false, "sort by time, newest first")
// 	flagWidth := flag.Int("w", 0, "set output width to COLS.  0 means no limit")
// 	flagSortExtension := flag.Bool("X", false, "sort alphabetically by entry extension")
// 	flagContext := flag.Bool("Z", false, "print any security context of each file")
// 	flagOne := flag.Bool("1", false, "list one file per line.  Avoid '\n' with -q or -b")
// 	flagHelp := flag.Bool("h", false, "display this help and exit")
// 	flagHelp2 := flag.Bool("help", false, "display this help and exit")
// 	flagVersion := flag.Bool("v", false, "output version information and exit")
// 	flagVersion2 := flag.Bool("version", false, "output version information and exit")
// 	flag.Parse()

// 	if *useColor {
// 		colorize(ColorBlue, "Hello, DigitalOcean!")
// 		return
// 	}

// 	path := "."
// 	err := ReadNames(path)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	err = ReadNamesOnly(path)
// 	if err != nil {
// 		log.Println(err)
// 	}
// }

// Usage: ls [OPTION]... [FILE]...
//
// Exit status:
//  0  if OK,
//  1  if minor problems (e.g., cannot access subdirectory),
//  2  if serious trouble (e.g., cannot access command-line argument).
//
// List information about the FILEs (the current directory by default).
// Sort entries alphabetically if none of -cftuvSUX nor --sort is specified.
//
// The SIZE argument is an integer and optional unit (example: 10K is 10*1024).
// Units are K,M,G,T,P,E,Z,Y (powers of 1024) or KB,MB,... (powers of 1000).
// Binary prefixes can be used, too: KiB=K, MiB=M, and so on.
//
// The TIME_STYLE argument can be full-iso, long-iso, iso, locale, or +FORMAT.
// FORMAT is interpreted like in date(1).  If FORMAT is FORMAT1<newline>FORMAT2,
// then FORMAT1 applies to non-recent files and FORMAT2 to recent files.
// TIME_STYLE prefixed with 'posix-' takes effect only outside the POSIX locale.
// Also the TIME_STYLE environment variable sets the default style to use.
//
// Using color to distinguish file types is disabled both by default and
// with --color=never.  With --color=auto, ls emits color codes only when
// standard output is connected to a terminal.  The LS_COLORS environment
// variable can change the settings.  Use the dircolors command to set it.
//
// Modeled after GNU coreutils: <https://www.gnu.org/software/coreutils/>
// Full documentation <https://www.gnu.org/software/coreutils/ls>
func ls() error {
	return nil
}
