package datatypes

import (
	"fmt"
	"strings"

	"github.com/fatih/structs"
	"github.com/skeptycal/goutil/types"
	"golang.org/x/sys/cpu"
)

type Any = types.Any

func CPUOptionsAvailable() string {
	sb := strings.Builder{}
	defer sb.Reset()

	for _, v := range cpuOptionMap {
		s := structs.New(v)
		for kk, vv := range s.Map() {
			if b, ok := vv.(bool); ok {
				if b {
					sb.WriteString(fmt.Sprintf("%v: %v\n", kk, vv))
				}
			}
		}
	}

	return sb.String()
}

func CPUOptions() string {
	sb := strings.Builder{}
	defer sb.Reset()

	for _, v := range cpuOptionMap {
		s := structs.New(v)
		for kk, vv := range s.Map() {
			if b, ok := vv.(bool); ok {
				if b || !b {
					sb.WriteString(fmt.Sprintf("%v: %v\n", kk, vv))
				}
			}
		}
	}

	return sb.String()
}

func HasAVX2() bool {
	return cpu.X86.HasAVX2
}

var cpuOptionMap = map[string]Any{
	"cpu.ARM":     cpu.ARM,
	"cpu.ARM64":   cpu.ARM64,
	"cpu.MIPS64X": cpu.MIPS64X,
	"cpu.PPC64":   cpu.PPC64,
	"cpu.S390X":   cpu.S390X,
	"cpu.X86":     cpu.X86,
}
