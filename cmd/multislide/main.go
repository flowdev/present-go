package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

type lineKind int

const (
	kindText = lineKind(iota)
	kindTitle
	kindStep
	kindNote
)

type presentLine struct {
	kind lineKind
	text string
}

var titleLine = regexp.MustCompile(`^(\*|##)\s+`)
var stepLine = regexp.MustCompile(`\s+\|\s*$`)
var noteLine = regexp.MustCompile(`^:\s+`)

func main() {
	if len(os.Args) > 1 && (os.Args[1] == "-h" || os.Args[1] == "--help") {
		fmt.Println("Usage: " + os.Args[0])
		fmt.Println("")
		fmt.Println("This command reads:  present.mslide")
		fmt.Println("This command writes: present.slide")
		fmt.Println("                     notes.txt")
		os.Exit(0)
	}

	f, err := os.Open("present.mslide")
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Unable to open file 'present.mslide': %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	lines, err := readLines(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Unable to read input slides: %v\n", err)
		os.Exit(2)
	}
	fmt.Fprintf(os.Stderr, "INFO: Found %d lines of input.\n", len(lines))

	plines := parseLines(lines)
	if err = writeNotes(plines); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
		os.Exit(3)
	}
	lines = handleSteps(plines)
	fmt.Fprintf(os.Stderr, "INFO: Got %d lines of output.\n", len(lines))

	if err = writeSlides(lines); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Unable to write slides: %v\n", err)
		os.Exit(4)
	}
}

func readLines(r io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(r)
	result := make([]string, 0, 8192)
	for scanner.Scan() {
		result = append(result, scanner.Text()) // without '\n'
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return result, nil
}

func parseLines(lines []string) []presentLine {
	result := make([]presentLine, len(lines))
	hasTitle := false // steps before the first title aren't steps
	for i, l := range lines {
		if titleLine.MatchString(l) {
			result[i] = presentLine{kind: kindTitle, text: l}
			hasTitle = true
		} else if hasTitle && stepLine.MatchString(l) {
			result[i] = presentLine{kind: kindStep, text: stepLine.ReplaceAllLiteralString(l, "")}
		} else if noteLine.MatchString(l) {
			result[i] = presentLine{kind: kindNote, text: noteLine.ReplaceAllLiteralString(l, "")}
		} else {
			result[i] = presentLine{kind: kindText, text: l}
		}
	}
	return result
}

func handleSteps(plines []presentLine) []string {
	result := make([]string, 0, 16*len(plines))
	i := 0

	for i < len(plines) {
		page := findPage(plines, i)
		result = handlePage(result, page)
		i += len(page)
		fmt.Fprintf(os.Stderr, "INFO: Handled page with %d lines (%d lines altogether).\n", len(page), len(result))
	}
	return result
}
func findPage(plines []presentLine, start int) []presentLine {
	i := start
	if i < len(plines) && plines[i].kind == kindTitle {
		i++
	}
	for ; i < len(plines) && plines[i].kind != kindTitle; i++ {
	}
	if i == len(plines) && i > 0 && strings.TrimSpace(plines[i-1].text) != "" {
		page := make([]presentLine, 0, i-start+1)
		page = append(page, plines[start:i]...)
		page = append(page, presentLine{kind: kindText, text: ""})
		return page
	}
	return plines[start:i]
}
func handlePage(result []string, page []presentLine) []string {
	if len(page) == 0 {
		return result
	}
	more := true
	for more {
		i := 0
		result, i = copyLinesUntilFirstStep(result, page)
		result = handleFirstStep(result, page, i)
		result, more = copyNonSteps(result, page, i+1)
	}
	return result
}
func copyLinesUntilFirstStep(result []string, page []presentLine) ([]string, int) {
	i := 0
	for ; i < len(page) && page[i].kind != kindStep; i++ {
		result = append(result, page[i].text)
	}
	return result, i
}
func handleFirstStep(result []string, page []presentLine, i int) []string {
	if i < len(page) {
		page[i].kind = kindText
		result = append(result, page[i].text)
	}
	return result
}
func copyNonSteps(result []string, page []presentLine, i int) ([]string, bool) {
	more := false
	for ; i < len(page); i++ {
		if page[i].kind != kindStep {
			result = append(result, page[i].text)
		} else {
			more = true
		}
	}
	return result, more
}

func writeSlides(lines []string) (err error) {
	f, err := os.Create("present.slide")
	if err != nil {
		return fmt.Errorf("unable to create file 'present.slide': %w", err)
	}
	defer func() { err = f.Close() }()

	for _, l := range lines {
		_, err := fmt.Fprintln(f, l)
		if err != nil {
			return fmt.Errorf("unable to write to 'present.slide': %w", err)
		}
	}
	return
}

func writeNotes(lines []presentLine) (err error) {
	var f *os.File

	f, err = os.Create("notes.txt")
	if err != nil {
		return fmt.Errorf("unable to create file 'notes.txt': %w", err)
	}
	defer func() { err = f.Close() }()

	first := true
	title := ""
	for _, l := range lines {
		if l.kind == kindTitle {
			title = l.text
		}
		if l.kind == kindNote {
			if title != "" {
				if !first {
					_, err := fmt.Fprintln(f, "")
					if err != nil {
						return fmt.Errorf("unable to write to 'notes.txt': %w", err)
					}
					_, err = fmt.Fprintln(f, "")
					if err != nil {
						return fmt.Errorf("unable to write to 'notes.txt': %w", err)
					}
				}
				_, err := fmt.Fprintln(f, title)
				if err != nil {
					return fmt.Errorf("unable to write to 'notes.txt': %w", err)
				}
				_, err = fmt.Fprintln(f, "")
				if err != nil {
					return fmt.Errorf("unable to write to 'notes.txt': %w", err)
				}
				title = ""
			}
			_, err := fmt.Fprintln(f, l.text)
			if err != nil {
				return fmt.Errorf("unable to write to 'notes.txt': %w", err)
			}
			first = false
		}
	}
	return
}
