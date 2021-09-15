package main

import (
	"log"
	"os"
	"os/exec"

	"github.com/neovim/go-client/nvim"
)

func main() {
	log.SetFlags(0)

	addr := os.Getenv("NVIM_LISTEN_ADDRESS")
	if addr == "" {
		editor := os.Getenv("FLATNVIM_EDITOR")

		path, err := exec.LookPath(editor)
		if err != nil {
			log.Fatalf("command '%v' set with environment variable FLATNVIM_EDITOR is not in $PATH\n", editor)
		}

		cmd := exec.Command(path, os.Args[1:]...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}

	files := os.Args[1:]
	if len(files) == 0 {
		log.Fatal("no arguments given")
	}

	v, err := nvim.Dial(addr)
	if err != nil {
		log.Fatal(err)
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
		log.Fatal(err)
	}
}
