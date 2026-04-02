package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const version = "0.1.0"

func main() {
	showVersion := flag.Bool("V", false, "Show version")
	flag.BoolVar(showVersion, "version", false, "Show version")
	showHelp := flag.Bool("h", false, "Show help")
	flag.BoolVar(showHelp, "help", false, "Show help")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "inxi-go %s — system information tool\n\nUsage: inxi-go [options]\n\nOptions:\n", version)
		flag.PrintDefaults()
	}
	flag.Parse()

	switch {
	case *showHelp:
		flag.Usage()
	case *showVersion:
		fmt.Printf("inxi-go %s\n", version)
	default:
		printMachine()
		printSystem()
		printCPU()
		printMemory()
		printGraphics()
		printAudio()
	}
}

// kv returns "key: value", or "" when value is empty.
func kv(key, val string) string {
	if val == "" {
		return ""
	}
	return key + ": " + val
}

// printSection prints a labelled block. Each element of lines is rendered as a
// space-joined row of key-value pairs. The label appears only on the first row.
func printSection(label string, lines ...[]string) {
	const width = 11
	indent := strings.Repeat(" ", width+1)
	first := true
	for _, pairs := range lines {
		pairs = compact(pairs)
		if len(pairs) == 0 {
			continue
		}
		row := strings.Join(pairs, "  ")
		if first {
			fmt.Printf("%-*s %s\n", width, label+":", row)
			first = false
		} else {
			fmt.Printf("%s%s\n", indent, row)
		}
	}
}

// compact removes empty strings from a slice without allocating when unneeded.
func compact(s []string) []string {
	out := s[:0:len(s)]
	for _, v := range s {
		if v != "" {
			out = append(out, v)
		}
	}
	return out
}

// readFile returns the trimmed content of a file, or "" on any error.
func readFile(path string) string {
	b, err := os.ReadFile(path)
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(b))
}

// run executes a command and returns its trimmed stdout, or "" on error.
func run(name string, args ...string) string {
	out, err := exec.Command(name, args...).Output()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(out))
}
