// Package cli implements the thinkpixelmemctl command-line interface.
// Administrative operations use the public ThinkPixelMEM API; there is no
// storage or internal-service back door.
package cli

import (
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

const (
	Name    = "thinkpixelmemctl"
	Version = "devel"
)

var ErrNotImplemented = errors.New("API operation is not implemented")

type command struct{ path, summary string }

var commands = []command{
	{"context retrieve", "Retrieve a governed ContextPack"},
	{"index rebuild", "Rebuild a disposable retrieval index"},
	{"ingestion status", "Show ingestion status"},
	{"memory correct", "Create a correcting memory revision"},
	{"memory forget", "Forget a memory"},
	{"memory get", "Get a memory"},
	{"memory inspect", "Inspect memory provenance and revisions"},
	{"memory quarantine", "Quarantine a memory"},
	{"memory-space create", "Create a MemorySpace"},
	{"memory-space describe", "Describe a MemorySpace"},
	{"profile inspect", "Inspect a derived profile"},
}

// Run executes the CLI and returns a process exit code. It never terminates the
// process itself, which keeps command behavior testable and embeddable.
func Run(args []string, stdout, stderr io.Writer) int {
	if stdout == nil || stderr == nil {
		return 2
	}
	if len(args) == 0 || isHelp(args[0]) {
		writeRootHelp(stdout)
		return 0
	}
	if args[0] == "version" || args[0] == "--version" || args[0] == "-version" {
		fmt.Fprintf(stdout, "%s %s\n", Name, Version)
		return 0
	}

	matched, complete := match(args)
	if matched == "" {
		fmt.Fprintf(stderr, "%s: unknown command %q\nRun %q for usage.\n", Name, strings.Join(args, " "), Name+" help")
		return 2
	}
	if !complete && (len(args) == 1 || (len(args) == 2 && isHelp(args[1]))) {
		writeCommandHelp(stdout, matched)
		return 0
	}
	if !complete {
		fmt.Fprintf(stderr, "%s: unknown command %q\nRun %q for usage.\n", Name, strings.Join(args, " "), Name+" "+matched+" --help")
		return 2
	}
	if len(args) > len(strings.Fields(matched)) && isHelp(args[len(args)-1]) {
		writeCommandHelp(stdout, matched)
		return 0
	}
	if len(args) != len(strings.Fields(matched)) {
		fmt.Fprintf(stderr, "%s %s: arguments are not available in the Phase 1 skeleton\n", Name, matched)
		return 2
	}
	fmt.Fprintf(stderr, "%s %s: %v; this command will use the public API\n", Name, matched, ErrNotImplemented)
	return 1
}

func match(args []string) (string, bool) {
	want := strings.Join(args, " ")
	var prefix string
	for _, candidate := range commands {
		if want == candidate.path {
			return candidate.path, true
		}
		if args[0] == strings.Fields(candidate.path)[0] {
			prefix = args[0]
		}
	}
	return prefix, false
}

func isHelp(value string) bool { return value == "help" || value == "--help" || value == "-h" }

func writeRootHelp(output io.Writer) {
	fmt.Fprintf(output, "Usage: %s <command>\n\nAdminister ThinkPixelMEM through its public API.\n\nCommands:\n", Name)
	groups := make(map[string]bool)
	for _, candidate := range commands {
		groups[strings.Fields(candidate.path)[0]] = true
	}
	names := []string{"help", "version"}
	for group := range groups {
		names = append(names, group)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Fprintf(output, "  %s\n", name)
	}
	fmt.Fprintf(output, "\nRun %q for command help.\n", Name+" <command> --help")
}

func writeCommandHelp(output io.Writer, group string) {
	fmt.Fprintf(output, "Usage: %s %s <command>\n\n", Name, group)
	for _, candidate := range commands {
		if strings.HasPrefix(candidate.path, group+" ") {
			fmt.Fprintf(output, "  %-12s %s\n", strings.TrimPrefix(candidate.path, group+" "), candidate.summary)
		}
	}
}
