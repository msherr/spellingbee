# spellingbee
Solver for the NYTimes Spelling Bee (I'm ashamed of myself for writing this)


To build, do:

`go build .`

## Usage

```
Spelling bee cheater!  Shame on you.
usage: spellingbee-cheater [-h|--help] [-d|--dict "<value>"] -p|--puzzle
                           "<value>"

                           solves NYTimes spelling bees

Arguments:

  -h  --help    Print help information
  -d  --dict    path to dictionary file. Default: /usr/share/dict/words
  -p  --puzzle  puzzle letters (put center letter first)
  ```
  
  Example:
  `./spellingbee -p shameful`
  
  Note that you should put the center letter first.
