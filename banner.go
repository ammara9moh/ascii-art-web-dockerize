package main

import (
	"strings"
)

const charHeight = 8

func renderASCII(s string, banner map[rune][]string) string {
	var sb strings.Builder
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		if line == "" {
			sb.WriteString("\n")
			continue
		}
		for row := 0; row < charHeight; row++ {
			for _, ch := range line {
				if art, ok := banner[ch]; ok {
					sb.WriteString(art[row])
				} else {
					sb.WriteString("        ")
				}
			}
			sb.WriteString("\n")
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func loadBanner(path string) (map[rune][]string, error) {
	data, err := appFiles.ReadFile(path)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")
	banner := make(map[rune][]string)
	ch := 32
	for i := 0; i < len(lines); {
		if lines[i] == "" {
			i++
			continue
		}
		if i+charHeight > len(lines) {
			break
		}
		banner[rune(ch)] = lines[i : i+charHeight]
		i += charHeight
		ch++
	}
	return banner, nil
}
