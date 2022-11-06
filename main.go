package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/neovim/go-client/nvim"
)

func main() {
	const version string = "v0.2.1"

	log.SetFlags(log.Ldate | log.Ltime)
	log.SetPrefix(version + " ")

	logfile := os.Getenv("FLATNVIM_LOGFILE")
	if logfile != "" {
		if f, err := os.OpenFile(logfile, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0o664); err == nil {
			defer f.Close()
			log.SetOutput(f)
		} else {
			log.Printf("FLATNVIM_LOGFILE is set to %v but cannot be opened: %v\n", logfile, err)
		}
	}

	err := os.Setenv("FLATNVIM_VERSION", version)
	if err != nil {
		log.Printf("unable to set FLATNVIM_VERSION environment variable: %v\n", err)
	}

	addr := os.Getenv("NVIM")
	if addr == "" {
		editor := os.Getenv("FLATNVIM_EDITOR")

		path, err := exec.LookPath(editor)
		if err != nil {
			log.Panicf("command '%v' from FLATNVIM_EDITOR is not in $PATH: %v\n", editor, err)
		}

		cmd := exec.Command(path, os.Args[1:]...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			log.Panicln(err)
		}

		return
	}

	files := os.Args[1:]
	if len(files) == 0 {
		log.Panicln("no arguments given")
	}

	v, err := nvim.Dial(addr)
	if err != nil {
		log.Panicf("unable to connect to parent nvim instance: %v\n", err)
	}
	defer v.Close()

	b := v.NewBatch()
	b.Command("let flatnvim_buf=bufname()")

	var curDir string
	if dir, err := os.Getwd(); err == nil {
		curDir = dir + string(os.PathSeparator)
	} else {
		log.Printf("could not get current directory: %v\n", err)
	}

	var vimDir string
	if dir, err := v.Exec("pwd", true); err == nil {
		vimDir = dir + string(os.PathSeparator)
	} else {
		log.Printf("could not get directory of nvim: %v\n", err)
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Printf("unable to get user home directory: %v\n", err)
	}

	for _, file := range files {
		if file == "--" {
			continue
		}

		path := trimPath(file, curDir, vimDir, homeDir)
		b.Command("e " + path)
	}

	b.Command("exe 'bd! '.flatnvim_buf")

	extraCmd := os.Getenv("FLATNVIM_EXTRA_COMMAND")
	if extraCmd != "" {
		b.Command(extraCmd)
	}

	if err := b.Execute(); err != nil {
		log.Panicf("unable to execute nvim commands: %v\n", err)
	}
}

func trimPath(path, curDir, vimDir, homeDir string) string {
	if len(curDir) > 0 && !filepath.IsAbs(path) {
		path = curDir + path
	}

	if len(vimDir) > 0 {
		path = strings.TrimPrefix(path, vimDir)
	}

	if len(homeDir) > 0 && strings.HasPrefix(path, homeDir) {
		path = strings.Replace(path, homeDir, "~", 1)
	}

	path = strings.ReplaceAll(path, "/./", "/")
	path = strings.TrimPrefix(path, "./")

	return path
}
