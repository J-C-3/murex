//go:build !windows
// +build !windows

package man

import (
	"bufio"
	"compress/gzip"
	"os"
	"os/exec"
	"regexp"
	"sort"
	"strings"
)

var (
	rxMatchManSection   = regexp.MustCompile(`/man[1678]/`)
	rxMatchFlagsEscaped = regexp.MustCompile(`\\f[BI]((\\-|-)[a-zA-Z0-9]|(\\-\\-|--)[\\\-a-zA-Z0-9]+).*?\\f[RP]`)
	rxMatchFlagsQuoted  = regexp.MustCompile(`\.IP "(.*?)"`)
	rxMatchFlagsDarwin  = regexp.MustCompile(`\.It Fl ([a-zA-Z0-9])`)
	rxMatchFlagsOther   = regexp.MustCompile(`\.B (.*?)`)
	//rxMatchFlagsOther   = regexp.MustCompile(`\.B (.*?)\\fR`)

	rxMatchGetFlag = regexp.MustCompile(`(--[\-a-zA-Z0-9]+)`)

	rxReplaceMarkup = regexp.MustCompile(`\.[a-zA-Z]+(\s|)`)
)

// GetManPages executes `man -w` to locate the manual files
func GetManPages(exe string) []string {
	// Get paths
	cmd := exec.Command("man", "-w", exe)
	b, err := cmd.Output()
	if err != nil {
		return nil
	}

	s := strings.TrimSpace(string(b))
	if s == exe {
		return nil
	}

	return strings.Split(s, ":")
}

func validMan(path string) bool {
	return !rxMatchManSection.MatchString(path) &&
		!strings.HasSuffix(path, "test/cat.1.gz")
}

// ParseFlags runs the parser to locate any flags with hyphen prefixes
func ParseFlags(paths []string) (flags []string) {
	// Parse man pages
	fMap := make(map[string]bool)
	for i := range paths {
		if validMan(paths[i]) {
			continue
		}
		parseFlags(&fMap, paths[i])
	}

	for f := range fMap {
		flags = append(flags, f)
	}
	sort.Strings(flags)
	return
}

// old parsing
func parseFlags(flags *map[string]bool, filename string) {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		return
	}

	var scanner *bufio.Scanner

	if len(filename) > 3 && filename[len(filename)-3:] == ".gz" {
		gz, err := gzip.NewReader(file)
		defer gz.Close()
		if err != nil {
			return
		}

		scanner = bufio.NewScanner(gz)
	} else {
		scanner = bufio.NewScanner(file)
	}

	for scanner.Scan() {
		s := scanner.Text()

		match := rxMatchFlagsEscaped.FindAllStringSubmatch(s, -1)
		for i := range match {
			if len(match[i]) == 0 {
				continue
			}

			s := strings.Replace(match[i][1], `\`, "", -1)
			if strings.HasSuffix(s, "fR") || strings.HasSuffix(s, "fP") {
				s = s[:len(s)-2]
			}
			(*flags)[s] = true
		}

		match = rxMatchFlagsQuoted.FindAllStringSubmatch(s, -1)
		for i := range match {
			if len(match[i]) == 0 {
				continue
			}

			flag := rxMatchGetFlag.FindAllStringSubmatch(match[i][1], -1)
			for j := range flag {
				if len(flag[j]) == 0 {
					continue
				}

				(*flags)[flag[j][1]] = true
			}
		}

		match = rxMatchFlagsDarwin.FindAllStringSubmatch(s, -1) // eg `cat` on OSX
		for i := range match {
			if len(match[i]) == 0 {
				continue
			}

			(*flags)["-"+match[i][1]] = true
		}

		match = rxMatchFlagsOther.FindAllStringSubmatch(s, -1)
		for i := range match {
			if len(match[i]) == 0 {
				continue
			}

			//// Fix \^ seen on some OSX man pages
			//match[i][1] = strings.Replace(match[i][1], `\^`, "", -1)

			flag := rxMatchGetFlag.FindAllStringSubmatch(match[i][1], -1)
			for j := range flag {
				if len(flag[j]) == 0 {
					continue
				}

				(*flags)[flag[j][1]] = true
			}
		}
	}

	return
}

// ParseSummary runs the parser to locate a summary
func ParseSummary(paths []string) string {
	for i := range paths {
		if validMan(paths[i]) {
			continue
		}
		desc := parseSummary(paths[i])
		if desc != "" {
			return desc
		}
	}

	return ""
}

func parseSummary(filename string) string {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		return ""
	}

	var scanner *bufio.Scanner

	if len(filename) > 3 && filename[len(filename)-3:] == ".gz" {
		gz, err := gzip.NewReader(file)
		defer gz.Close()
		if err != nil {
			return ""
		}

		scanner = bufio.NewScanner(gz)
	} else {
		scanner = bufio.NewScanner(file)
	}

	var (
		read bool
		desc string
	)

	for scanner.Scan() {
		s := scanner.Text()

		if strings.Contains(s, "SYNOPSIS") {
			if len(desc) > 0 && desc[len(desc)-1] == '-' {
				desc = desc[:len(desc)-1]
			}
			return strings.TrimSpace(desc)
		}

		if read {
			// Tidy up man pages generated from reStructuredText
			if strings.HasPrefix(s, `\\n[rst2man-indent`) ||
				strings.HasPrefix(s, `\\$1 \\n`) ||
				strings.HasPrefix(s, `level \\n`) ||
				strings.HasPrefix(s, `level margin: \\n`) {
				continue
			}

			s = strings.Replace(s, ".Nd ", " - ", -1)
			s = strings.Replace(s, "\\(em ", " - ", -1)
			s = strings.Replace(s, " , ", ", ", -1)
			s = strings.Replace(s, "\\fB", "", -1)
			s = strings.Replace(s, "\\fR", "", -1)
			if strings.HasSuffix(s, " ,") {
				s = s[:len(s)-2] + ", "
			}
			s = rxReplaceMarkup.ReplaceAllString(s, "")
			s = strings.Replace(s, "\\", "", -1)

			if strings.HasPrefix(s, `.`) {
				continue
			}

			desc += s
		}

		if strings.Contains(s, "NAME") {
			read = true
		}
	}

	return ""
}
