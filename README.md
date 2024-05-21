# spellingbee
Solver for the NYTimes Spelling Bee (I'm ashamed of myself for writing this)

[Micah Sherr](mailto:msherr@cs.georgetown.edu) is responsible for this horrible thing.


## Compilation

To build, do:

```bash
go mod init spellingbee
go mod tidy
go build
```

Note that you should have some dictionary installed.  By default, we look for a word-list (i.e., a file with words in it, delimited by newline characters) at `/usr/share/dict/words`, although this can be overridden using the `-d` or `--dict` option.

This file seems to exist on Macs (at least on my Mac).  On Ubuntu, you can install a suitable word-list via `sudo apt install wamerican-small`.


## Usage

```
usage: spellingbee [-h|--help] [-d|--dict "<value>"] -p|--puzzle "<value>"
                   [-q|--quiet]

                   solves NYTimes spelling bees

Arguments:

  -h  --help    Print help information
  -d  --dict    path to dictionary file. Default: /usr/share/dict/words
  -p  --puzzle  puzzle letters (put center letter first)
  -q  --quiet   quiet mode; only output words
  ```
  
  Example:
  `./spellingbee -p shueamfl`
  
  Note that you should put the center letter first.
