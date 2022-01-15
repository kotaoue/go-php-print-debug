package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var directories = flag.String("directories", "", "directories to search. if you specify more than one, separate them with commas. eg. app,vendor")

func main() {
	flag.Parse()

	if err := Main(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func Main() error {
	find := false
	for _, v := range strings.Split(*directories, ",") {
		b, err := walk(v)
		if err != nil {
			return err
		}
		if b && !find {
			find = true
		}
	}

	if find {
		return errors.New("found a print debug, check the output above")
	}

	return nil
}

func walk(root string) (bool, error) {
	fmt.Printf("# root: %s\n", root)

	find := false
	err := filepath.Walk(root, func(path string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		switch filepath.Ext(path) {
		case ".php":
			fmt.Printf("## path: %s\n", path)
			b, err := search(path)
			if err != nil {
				return err
			}
			if b && !find {
				find = true
			}
		}

		return nil
	})

	return find, err
}

func search(path string) (bool, error) {
	file, err := os.Open(path)
	if err != nil {
		return false, err
	}

	fs := bufio.NewScanner(file)
	i := 0
	find := false
	for fs.Scan() {
		i++
		if isPrintDebug(fs.Text()) {
			fmt.Printf("- [ ] find print debug@%s:%d, %s\n", path, i, strings.TrimSpace(fs.Text()))
			find = true
		}
	}

	return find, fs.Err()
}

func isPrintDebug(s string) bool {
	for _, v := range []string{"print", "print_r", "var_dump", "var_export", "echo"} {
		if strings.Contains(s, v) {
			i := strings.Index(s, v) + len(v)
			if b := (s[i:i+1] == "(" || s[i:i+1] == " "); b {
				return b
			}
		}
	}
	return false
}
