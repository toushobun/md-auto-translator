package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	srcDir string
	outDir string
	target string
)

// go run main.go --src=./testFile/original/toutest.md --out=./testFile/result
func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339Nano})
	flag.StringVar(&srcDir, "src", "", "source directory")
	flag.StringVar(&outDir, "out", "", "out directory")
	flag.StringVar(&target, "target", "en", "target languages(en/ja/...)")
}

func run() error {

	return filepath.WalkDir(srcDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || !strings.HasSuffix(path, ".md") {
			return nil
		}

		// TODO: process file
		createScaffold(path, srcDir, outDir, target)
		return nil
	})
}

func createScaffold(srcPath, srcRoot, outRoot, target string) {
	targetPath := strings.Replace(srcPath, srcRoot, outRoot, 1)

	if _, err := os.Stat(targetPath); err == nil {
		return // already exists
	}

	data, err := os.ReadFile(srcPath)
	if err != nil {
		fmt.Println("read failed:", srcPath)
		return
	}

	fm := string(data)

	var buf bytes.Buffer
	buf.WriteString(fm)
	buf.WriteString("\n\n")
	buf.WriteString(fmt.Sprintf("<!-- TODO: translate from %s -->\n", srcPath))

	err = os.MkdirAll(filepath.Dir(targetPath), 0755)
	if err != nil {
		fmt.Println("mkdir failed:", targetPath)
		return
	}

	err = os.WriteFile(targetPath, buf.Bytes(), 0644)
	if err != nil {
		fmt.Println("write failed:", targetPath)
		return
	}

	fmt.Println("generated:", targetPath)
}

func main() {
	flag.Parse()
	if err := run(); err != nil {
		log.Error().Err(err).Send()
	}
}
