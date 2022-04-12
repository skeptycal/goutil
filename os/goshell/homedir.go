package goshell

// The functionality in this file (homedir and caching) are modeled from directly
// from https://github.com/mitchellh/go-homedir

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

// defaultDisableCache is the default value for disableCache.
const defaultDisableCache = false

// type strErrConvertFunc func(fn func() (string, error)) string

var (
	HOME    = NewHome
	dirfunc func() (string, error)
)

func init() {
	if runtime.GOOS == "windows" {
		dirfunc = dirWindows
	} else {
		// Unix-like system, so just assume Unix
		dirfunc = dirUnix
	}
}

func NewHome() Dir {
	return &homedir{
		disableCache: defaultDisableCache,
	}
}

type homedir struct {
	homedirCache string
	disableCache bool
	cacheLock    sync.RWMutex
}

type Dir interface {
	Abs() string
	Base() string
	Reset()
	SetCache(disabled bool)
	// Ls(pattern string) []string
}

// Abs returns the home directory for the executing user.
//
// This uses an OS-specific method for discovering the home directory.
// An error is returned if a home directory cannot be detected.
func (p *homedir) Abs() string {
	if !p.disableCache {
		p.cacheLock.RLock()
		cached := p.homedirCache
		p.cacheLock.RUnlock()
		if cached != "" {
			return cached
		}
	}

	defer p.unlock()
	p.lock()

	p.homedirCache = result
	return result
}

func (p *homedir) Base() string {
	return filepath.Base(p.Abs())
}

// SetCache specifies whether directory name cache should be disabled.
// The default is false (cache enabled).
func (p *homedir) SetCache(disabled bool) {

	// a lock is acquired to prevent changing p.disableCache
	// concurrently with the updating of homedirCache.
	defer p.cacheLock.Unlock()
	p.cacheLock.Lock()

	p.disableCache = disabled
}

// Reset clears the cache, forcing the next call to Dir to re-detect
// the home directory. This generally never has to be called, but can be
// useful in tests if you're modifying the home directory via the HOME
// env var or something.
func (p *homedir) Reset() {
	defer p.cacheLock.Unlock()
	p.lock()
	homedirCache = ""
}

// lock locks rw for writing. If the lock is already locked for reading
// or writing, lock blocks until the lock is available.
func (p *homedir) lock() {
	p.cacheLock.Lock()
}

// Unlock unlocks rw for writing. It is a run-time error if rw is not
// locked for writing on entry to Unlock.
//
// As with Mutexes, a locked RWMutex is not associated with a particular
// goroutine. One goroutine may RLock (Lock) a RWMutex and then arrange
// for another goroutine to RUnlock (Unlock) it.
func (p *homedir) unlock() {
	p.cacheLock.Unlock()
}

func (p *homedir) get() (cached string) {
	if !p.disableCache {
		p.cacheLock.RLock()
		cached = p.homedirCache
		p.cacheLock.RUnlock()
		if cached != "" {
			return
		}
	}

	return p.getNoCache()
}

func (p *homedir) getNoCache() string {
	return ""
}

/// ------------------------------------------------------ old version
// DisableCache will disable caching of the home directory. Caching is enabled
// by default.
var DisableCache bool

var homedirCache string
var cacheLock sync.RWMutex

// GetHomeDir returns the home directory for the executing user.
//
// This uses an OS-specific method for discovering the home directory.
// An error is returned if a home directory cannot be detected.
func GetHomeDir() (string, error) {
	if !DisableCache {
		cacheLock.RLock()
		cached := homedirCache
		cacheLock.RUnlock()
		if cached != "" {
			return cached, nil
		}
	}

	cacheLock.Lock()
	defer cacheLock.Unlock()

	var result string
	var err error
	if runtime.GOOS == "windows" {
		result, err = dirWindows()
	} else {
		// Unix-like system, so just assume Unix
		result, err = dirUnix()
	}

	if err != nil {
		return "", err
	}
	homedirCache = result
	return result, nil
}

// Expand expands the path to include the home directory if the path
// is prefixed with `~`. If it isn't prefixed with `~`, the path is
// returned as-is.
func Expand(path string) (string, error) {
	if len(path) == 0 {
		return path, nil
	}

	if path[0] != '~' {
		return path, nil
	}

	if len(path) > 1 && path[1] != '/' && path[1] != '\\' {
		return "", errors.New("cannot expand user-specific home dir")
	}

	dir, err := GetHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(dir, path[1:]), nil
}

// Reset clears the cache, forcing the next call to Dir to re-detect
// the home directory. This generally never has to be called, but can be
// useful in tests if you're modifying the home directory via the HOME
// env var or something.
func Reset() {
	cacheLock.Lock()
	defer cacheLock.Unlock()
	homedirCache = ""
}

func dirUnix() (string, error) {
	homeEnv := "HOME"
	if runtime.GOOS == "plan9" {
		// On plan9, env vars are lowercase.
		homeEnv = "home"
	}

	// First prefer the HOME environmental variable
	if home := os.Getenv(homeEnv); home != "" {
		return home, nil
	}

	var stdout bytes.Buffer

	// If that fails, try OS specific commands
	if runtime.GOOS == "darwin" {
		cmd := exec.Command("sh", "-c", `dscl -q . -read /Users/"$(whoami)" NFSHomeDirectory | sed 's/^[^ ]*: //'`)
		cmd.Stdout = &stdout
		if err := cmd.Run(); err == nil {
			result := strings.TrimSpace(stdout.String())
			if result != "" {
				return result, nil
			}
		}
	} else {
		cmd := exec.Command("getent", "passwd", strconv.Itoa(os.Getuid()))
		cmd.Stdout = &stdout
		if err := cmd.Run(); err != nil {
			// If the error is ErrNotFound, we ignore it. Otherwise, return it.
			if err != exec.ErrNotFound {
				return "", err
			}
		} else {
			if passwd := strings.TrimSpace(stdout.String()); passwd != "" {
				// username:password:uid:gid:gecos:home:shell
				passwdParts := strings.SplitN(passwd, ":", 7)
				if len(passwdParts) > 5 {
					return passwdParts[5], nil
				}
			}
		}
	}

	// If all else fails, try the shell
	stdout.Reset()
	cmd := exec.Command("sh", "-c", "cd && pwd")
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		return "", err
	}

	result := strings.TrimSpace(stdout.String())
	if result == "" {
		return "", errors.New("blank output when reading home directory")
	}

	return result, nil
}

func dirWindows() (string, error) {
	// First prefer the HOME environmental variable
	if home := os.Getenv("HOME"); home != "" {
		return home, nil
	}

	// Prefer standard environment variable USERPROFILE
	if home := os.Getenv("USERPROFILE"); home != "" {
		return home, nil
	}

	drive := os.Getenv("HOMEDRIVE")
	path := os.Getenv("HOMEPATH")
	home := drive + path
	if drive == "" || path == "" {
		return "", errors.New("HOMEDRIVE, HOMEPATH, or USERPROFILE are blank")
	}

	return home, nil
}
