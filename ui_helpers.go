package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/term"
)

var re_ansi_escape = regexp.MustCompile(`\x1b\[[0-9;]*[A-Za-z]`)

func visibleLen(val string) int {
	clean := re_ansi_escape.ReplaceAllString(val, "")
	count := 0
	runes := []rune(clean)
	for _, r := range runes {
		if unicode.IsPrint(r) {
			count++
		}
	}
	return count
}

func terminalWidth() (int, error) {
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	return width, err
}

// Pads the left side of the string so that the result is at least min_length
// visible runes wide
func LeftPadLine(val string, min_length int, wrap_with rune) string {
	if len(val) >= min_length {
		return val
	}

	var result strings.Builder
	padded_len := len(val)
	for ; min_length-padded_len > 0; padded_len++ {
		result.WriteRune(wrap_with)
	}

	result.WriteString(val)
	return result.String()
}

// Pads the right side of the string so that the result is at least min_length
// visible runes wide
func RightPadLine(val string, min_length int, wrap_with rune) string {
	if len(val) >= min_length {
		return val
	}

	var result strings.Builder
	result.WriteString(val)

	padded_len := len(val)
	for ; min_length-padded_len > 0; padded_len++ {
		result.WriteRune(wrap_with)
	}
	return result.String()
}

// Wraps all strings in a single-line border
func WrapBox(val string, x_pad int, y_pad int) string {
	lines := strings.Split(val, "\n")
	maxlinelen := visibleLen(lines[0])

	for _, line := range lines[1:] {
		candidate := visibleLen(line)
		if maxlinelen < candidate {
			maxlinelen = candidate
		}
	}
	maxlinelen += 2 * x_pad
	empty_line := fmt.Sprintf("│%s│\n", strings.Repeat(" ", maxlinelen))
	sx_pad := strings.Repeat(" ", x_pad)
	sy_pad := strings.Repeat(empty_line, y_pad)

	var b_vert_border strings.Builder
	b_vert_border.WriteString(strings.Repeat("─", maxlinelen))
	vert_border := b_vert_border.String()

	var result strings.Builder
	fmt.Fprintf(&result, "┌%s┐\n%s", vert_border, sy_pad)
	for _, line := range lines {
		fmt.Fprintf(
			&result, "│%s%s%s│\n",
			sx_pad,
			RightPadLine(line, maxlinelen-2*x_pad, ' '),
			sx_pad)
	}
	fmt.Fprintf(&result, "%s└%s┘", sy_pad, vert_border)

	return result.String()
}

// Centers all of the lines in val based on the terminal width
func CenterBox(val string) (string, error) {
	width, err := terminalWidth()
	if err != nil {
		return "", err
	}

	var result strings.Builder
	lines := strings.SplitSeq(val, "\n")
	for line := range lines {
		half_w := visibleLen(line) / 2

		left_pad := ""
		if left_w := width/2 - half_w; left_w > 0 {
			left_pad = strings.Repeat(" ", left_w)
		}

		fmt.Fprintf(&result, "%s%s\n", left_pad, line)
	}

	return result.String(), nil
}
