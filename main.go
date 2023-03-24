package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	filename := flag.String("file", "in.md", "File to read from")
	output := flag.String("output", "out.md", "File to write to")
	flag.Parse()

	if err := readText(*filename, *output); err != nil {
		fmt.Fprintf(os.Stderr, "fmtText: %v", err)
		os.Exit(1)
	}
}

func readText(infile, outfile string) error {
	file, err := os.Open(infile)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		text := fmtText(line)
		if err := writeText(outfile, text); err != nil {
			return err
		}
	}

	return scanner.Err()
}

func fmtText(line string) string {
	dollar := 0
	pos := 0
	text := ""
	for i := 0; i < len(line); i++ {
		// dollarが0の時は、$が出てきたら$の前に空白を追加し、dollarを1にする
		// dollarが1の時は、$が出てきたら$の後に空白を追加し、dollarを0にする
		if line[i] == '$' {
			if dollar == 0 {
				text += line[pos:i] + " $"
				pos = i + 1
				dollar = 1
			} else {
				text += line[pos:i] + "$ "
				pos = i + 1
				dollar = 0
			}
		}
	}
	text += line[pos:]
	return text
}

func writeText(filename, line string) error {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(line + "\n")
	return err
}
