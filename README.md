# spellingbee
Solver for the NYTimes Spelling Bee (I'm ashamed of myself for writing this)


## Compilation

To build, do:

`go build .`


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
