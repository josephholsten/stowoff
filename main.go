package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/josephholsten/stowoff/app"
)

var packageName string
var stowDir string
var sourceDir string

func init() {
	homeDir := os.Getenv("HOME")
	flag.StringVar(&packageName, "package", "local", "package name")
	flag.StringVar(&stowDir, "stow-dir", filepath.Join(homeDir, ".dotfiles"), "directory into which to stow files")
	flag.StringVar(&sourceDir, "source-dir", homeDir, "directory from which to import files")
}

func main() {
	// log.SetFlags(0) // disable time in output
	// filter := &logutils.LevelFilter{
	// 	Levels: []logutils.LogLevel{
	// 		"DEBUG",
	// 		"INFO",
	// 		"NOTICE",
	// 		"WARN",
	// 		"ERROR",
	// 		"CRITICAL",
	// 		"ALERT",
	// 		"PANIC",
	// 	},
	// 	MinLevel: "WARN",
	// 	Writer:   os.Stderr,
	// }
	// log.SetOutput(filter)

	flag.Parse()
	args := flag.Args()
	if 0 == len(args) {
		usage()
	}
	command := args[0]
	switch command {
	case "import":
		importCommand(args[1:])
	case "load":
		loadCommand(args[1:])
	default:
		usage()
	}
}

func usage() {
	fmt.Printf("usage")
	os.Exit(1)
}

func importCommand(args []string) {
	// TODO: support multiple files
	file := args[0]

	err := importFile(file, packageName)
	if err != nil {
		log.Fatalf("importFile: %v", err)
	}

	fmt.Printf("imported successfully, now link by running:\n")
	fmt.Printf("  stow %v\n", packageName)
}

// packageDir MAY already exist
// packageDir MAY be created
// packageDir MUST be a directory
// sourcePath MUST exist
// sourcePath MUST NOT be symbolic link
// destPath MUST NOT exist
func importFile(file, pkg string) error {
	packageDir := filepath.Join(stowDir, pkg)
	sourcePath := filepath.Join(sourceDir, file)
	destPath := filepath.Join(packageDir, file)

	// optionally create packageDir
	if err := os.Mkdir(packageDir, 0750); err != nil && !os.IsExist(err) {
		return err
	}

	// verify sane environment {
	stat, err := os.Lstat(packageDir)
	if err != nil {
		return err
	}
	if !stat.IsDir() {
		return fmt.Errorf("package path %v is not a directory", packageDir)
	}

	stat, err = os.Lstat(sourcePath)
	if err != nil {
		return err
	}
	if stat.Mode()&os.ModeSymlink != 0 {
		return fmt.Errorf("source path %v is a symbolic link", sourcePath)
	}

	if _, err = os.Stat(destPath); err == nil || !os.IsNotExist(err) {
		return fmt.Errorf("destination path %v exists", destPath)
	}
	// }

	// Apply changes
	if err = os.Link(sourcePath, destPath); err != nil {
		return err
	}

	if err = os.Remove(sourcePath); err != nil {
		return err
	}

	return nil
}

func loadCommand(args []string) {
	if 0 == len(args) {
		log.Fatalf("usage: stowoff load app\n")
	}

	cfg, err := loadApp(args[0])
	if err != nil {
		log.Fatalf("loadCommand: %v", err)
	}
	fmt.Printf("cfg: %v\n", cfg)
}

