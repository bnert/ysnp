# You Shall Not Pass! (ysnp)
![Gandalf Being Gandalf](https://media.giphy.com/media/8abAbOrQ9rvLG/giphy.gif)

This is a client only, cli based password manager with a pretty simple interface.
```
$ ysnp help
Usage:
  ysnp [command]

Available Commands:
  help        Help about any command
  read        Read password file
  write       Write password file

Flags:
  -h, --help   help for ysnp

Use "ysnp [command] --help" for more information about a command.
```

### `read` sub-command
```
$ ysnp read -h
Read password file

Usage:
  ysnp read [from file.ysnp] [flags]

Flags:
  -h, --help   help for read
```

### `write` sub-command
```
$ ysnp write -h
Write password file

Usage:
  ysnp write [to file.ysnp] [flags]

Flags:
  -h, --help   help for write
```

## Building
### Pre-reqs
1. `go` toolchain in order to build the application: https://golang.org/dl/
2. `make` executable (provided with most distros)

### Build with make
```
$ make build-release
```

### Build using `go`
```
$ go build -o ysnp *.go
```

### Installing
Once the binary is built, simple copy or symlink the binary into `/usr/local/bin` (May require `sudo`):
Ex.
```
$ cp <path to>/ysnp /usr/local/bin
```

If needed, you may need to add `/usr/local/bin` to your PATH:
```
$ export PATH=$PATH:/usr/loca/bin
```
**NOTE**: I highly suggest adding this to your `$HOME/.bashrc` or `$HOME/.profile` if it is not already present.
