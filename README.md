# About

The goal of `flatnvim` is to make it easy to use [Neovim](https://neovim.io/)'s
[terminal emulator](https://neovim.io/doc/user/nvim_terminal_emulator.html).
When you open files from within the `Neovim` terminal, `flatnvim` will automatically add them
to the current instance instead of creating a new nested one.
Now you can easily use all your favorite command line programs inside `Neovim`.

# Usage

## Clone & Build
You'll first need the `flatnvim` binary on you system.

```sh
git clone https://github.com/adamtabrams/flatnvim.git
cd flatnvim
./build.sh
```

## Environment Variables
There are 4 environment variables `flatnvim` can use.
Add them to the config file for your shell (.bashrc, .profile, .zprofile, etc).

### Required: set the path to the actual editor.
```sh
export FLATNVIM_EDITOR="nvim"
```
- When `flatnvim` is called from your regular terminal, it will just pass arguments to this editor.

### Recommended: set the path to the `flatnvim` binary as your default terminal editor.
```sh
export EDITOR="$HOME/repos/flatnvim/bin/flatnvim"
```
- If you don't know what this should be, use the path printed by the build.sh script.
- In addition, you can make an alias to `flatnvim`: `alias vim="$EDITOR`

### Optional: set a log file for `flatnvim` to use.
```sh
export FLATNVIM_LOGFILE="$HOME/repos/flatnvim/log.txt"
```
- If this is not set, any error messages are just printed.

### Optional: set an extra `Neovim` command to be executed when preventing nested instances.
```sh
export FLATNVIM_EXTRA_COMMAND="echo 'it works' | sleep"
```
- Just in case you want `flatnvim` to do something extra before it opens file in the parent instance.


## Neovim Configs

### Recommended: disable status line when in terminal mode.
```viml
autocmd TermOpen * setlocal laststatus=0 noshowmode noruler
  \| autocmd TermClose * setlocal laststatus=2 showmode ruler

```

### Recommended: quickly exit terminal if commands were successful.
```viml
autocmd TermClose * if !v:event.status | exe 'bdelete! '..expand('<abuf>') | endif
```

### Optional: fixes issue with Airline that occurs when window focus is lost.
```viml
autocmd TermLeave * AirlineRefresh
```


## Neovim Function

Adding this function to your `Neovim` config makes it easier to access the terminal and command line programs.

```viml
function! TempTerm(...)
    let command = get(a:, 1)
    exe "terminal ".command
    return ""
endfunction
```

```viml
nnoremap <silent> gt  :call TempTerm(" ")<CR>
nnoremap <silent> gL  :call TempTerm("lazygit")<CR>
nnoremap <silent> gl  :call TempTerm("lf")<CR>
```
