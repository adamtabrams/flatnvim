package main

import (
	"log"
	"os"
	"os/exec"

	"github.com/neovim/go-client/nvim"
)

func main() {
	log.SetFlags(0)

	const version string = "v0.2.0"
	err := os.Setenv("FLATNVIM_VERSION", version)
	if err != nil {
		log.Fatalf("%v: unable to set FLATNVIM_VERSION environment variable\n", version)
	}

	addr := os.Getenv("NVIM_LISTEN_ADDRESS")
	if addr == "" {
		editor := os.Getenv("FLATNVIM_EDITOR")

		path, err := exec.LookPath(editor)
		if err != nil {
			log.Fatalf("%v: command '%v' from FLATNVIM_EDITOR is not in $PATH\n", version, editor)
		}

		cmd := exec.Command(path, os.Args[1:]...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			log.Fatalf("%v: %v\n", version, err)
		}
		os.Exit(0)
	}

	files := os.Args[1:]
	if len(files) == 0 {
		log.Fatalf("%v: no arguments given\n", version)
	}

	v, err := nvim.Dial(addr)
	if err != nil {
		log.Fatalf("%v: %v\n", version, err)
	}
	defer v.Close()

	b := v.NewBatch()
	for _, file := range files {
		b.Command(":e " + file)
	}
	b.Command(":b term://")
	b.Command(":bw!")

	extraCmd := os.Getenv("FLATNVIM_EXTRA_COMMAND")
	if extraCmd != "" {
		b.Command(extraCmd)
	}

	if err := b.Execute(); err != nil {
		log.Fatalf("%v: %v\n", version, err)
	}
}
