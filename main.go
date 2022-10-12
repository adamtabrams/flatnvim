package main

import (
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/neovim/go-client/nvim"
)

func main() {
	const version string = "v0.2.1"

	log.SetFlags(log.Ldate | log.Ltime)
	log.SetPrefix(version + " ")

	logfile := os.Getenv("FLATNVIM_LOGFILE")
	if logfile != "" {
		if f, err := os.OpenFile(logfile, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0664); err != nil {
			log.Printf("FLATNVIM_LOGFILE is set to %v but cannot be opened: %v\n", logfile, err)
		} else {
			defer f.Close()
			log.SetOutput(f)
		}
	}

	err := os.Setenv("FLATNVIM_VERSION", version)
	if err != nil {
		log.Fatalln("unable to set FLATNVIM_VERSION environment variable")
	}

	addr := os.Getenv("NVIM")
	if addr == "" {
		editor := os.Getenv("FLATNVIM_EDITOR")

		path, err := exec.LookPath(editor)
		if err != nil {
			log.Fatalf("command '%v' from FLATNVIM_EDITOR is not in $PATH\n", editor)
		}

		cmd := exec.Command(path, os.Args[1:]...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			log.Fatalln(err)
		}
		os.Exit(0)
	}

	files := os.Args[1:]
	if len(files) == 0 {
		log.Fatalln("no arguments given")
	}

	v, err := nvim.Dial(addr)
	if err != nil {
		log.Fatalln(err)
	}
	defer v.Close()

	b := v.NewBatch()
	b.Command(":let flatnvim_buf=bufname()")

	var wd string
	if dir, err := os.Getwd(); err != nil {
		log.Println("could not get current working directory")
	} else {
		wd = dir + string(os.PathSeparator)

	}

	for _, file := range files {
		log.Println(file)
		if file == "--" {
			continue
		}
		if len(wd) > 0 {
			file = strings.TrimPrefix(file, wd)
		}
		b.Command(":e " + file)
	}

	b.Command(":exe 'bd! '.flatnvim_buf")

	// TODO fix airline
	extraCmd := os.Getenv("FLATNVIM_EXTRA_COMMAND")
	if extraCmd != "" {
		b.Command(extraCmd)
	}

	if err := b.Execute(); err != nil {
		log.Fatalln(err)
	}
}
