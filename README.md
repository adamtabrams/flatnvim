# About

The goal of `flatnvim` is to make it easy to get the most out of [Neovim](https://neovim.io/)'s
[terminal emulator](https://neovim.io/doc/user/nvim_terminal_emulator.html).
When you open files from within the `Neovim` terminal, `flatnvim` will automatically add them
to the current instance instead of creating a new nested one.
This makes it easy to use all your favorite programs naturally from within `Neovim`.

# Usage

## Clone & Build
You'll first need the `flatnvim` binary on you system.

```sh
git clone https://github.com/adamtabrams/flatnvim.git
cd flatnvim
./build.sh
```

## Environment Variables
There are 3 environment variable to use with `flatnvim`.
Add them to the config file for your shell (.bashrc, .profile, .zprofile, etc).

### Required: set the path to the actual editor.
```sh
export FLATNVIM_EDITOR="nvim"
```
- When `flatnvim` is called from your regular terminal, it will just pass through to this editor.

### Recommended: set the path to the `flatnvim` binary as your default terminal editor.
```sh
export EDITOR="$HOME/repos/flatnvim/bin/flatnvim"
```
- If you don't know what this should be, use the path printed by the build.sh script.
- In addition, I personally make aliases to this: `alias vim="$EDITOR`

### Optional: set an extra Neovim command to be executed when preventing nested instances.
```sh
export FLATNVIM_EXTRA_COMMAND="if exists(':AirlineRefresh') == 2 | AirlineRefresh | endif"
```
- I use the command above to refresh my status bar.


## Neovim Configs

### Recommended: disable status line when in terminal mode.
```sh
autocmd TermOpen * setlocal laststatus=0 noshowmode noruler
  \| autocmd TermClose * setlocal laststatus=2 showmode ruler

```
