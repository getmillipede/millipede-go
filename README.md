millipede-go :bug:
==================

[![Build Status](https://travis-ci.org/getmillipede/millipede-go.svg?branch=master)](https://travis-ci.org/getmillipede/millipede-go)
[![Circle CI](https://img.shields.io/circleci/project/getmillipede/millipede-go.svg)](https://circleci.com/gh/getmillipede/millipede-go)
[![GoDoc](https://godoc.org/github.com/getmillipede/millipede-go/millipede?status.svg)](https://godoc.org/github.com/getmillipede/millipede-go/millipede)
[![Coverage Status](https://coveralls.io/repos/getmillipede/millipede-go/badge.svg?branch=master&service=github)](https://coveralls.io/github/getmillipede/millipede-go?branch=master)
![License](https://img.shields.io/github/license/getmillipede/millipede-go.svg)

![](https://raw.githubusercontent.com/getmillipede/millipede-go/master/assets/millipede.gif)

Golang version of [getmillipede](https://github.com/getmillipede/) with some exclusive features.

## Components

- **./**: the millipede Golang API
- [./cmd/millipede-go](https://github.com/getmillipede/millipede-go/tree/master/cmd/millipede-go): the millipede CLI
- [./cmd/millipede-web](https://github.com/getmillipede/millipede-go/tree/master/cmd/millipede-go): the millipede web server

## CLI Usage

```command
$ millipede-go -h
NAME:
   millipede-go - Print a beautiful millipede

USAGE:
   millipede-go [global options] command [command options] [arguments...]

VERSION:
   1.3.0-dev (HEAD)

AUTHOR(S):
   Millipede crew <https://github.com/getmillipede/millipede-go>

COMMANDS:
   help, h     Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --reverse, -r                           reverse the millipede
   --opposite, -o                          go the opposite direction
   --chameleon                             the millipede use its environment color
   --rainbow                               the millipede live with care bears
   --zalgo                                 invoke the hive-mind representing chaos
   --animate                               he is alive !
   --center, -C                            millipede in the midle
   --skin, --template, -s, -t "default"    millipede skin (default, frozen, love, corporate, musician, bocal, ascii, inception, humancentipede, finger)
   --width, -w "3"                         millipede width
   --curve, -c "4"                         millipede curve size
   --random                                RANDOMZ!
   --help, -h                              show help
   --version, -v                           print the version
```

### Examples

```command
$ millipede-go
    ╚⊙ ⊙╝
  ╚═(███)═╝
 ╚═(███)═╝
╚═(███)═╝
 ╚═(███)═╝
  ╚═(███)═╝
   ╚═(███)═╝
    ╚═(███)═╝
    ╚═(███)═╝
   ╚═(███)═╝
  ╚═(███)═╝
 ╚═(███)═╝
╚═(███)═╝
 ╚═(███)═╝
  ╚═(███)═╝
   ╚═(███)═╝
    ╚═(███)═╝
    ╚═(███)═╝
   ╚═(███)═╝
  ╚═(███)═╝
 ╚═(███)═╝
```

```command
$ millipede-go 42
    ╚⊙ ⊙╝
  ╚═(███)═╝
 ╚═(███)═╝
╚═(███)═╝
 ╚═(███)═╝
  ╚═(███)═╝
   ╚═(███)═╝
    ╚═(███)═╝
    ╚═(███)═╝
   ╚═(███)═╝
  ╚═(███)═╝
 ╚═(███)═╝
╚═(███)═╝
 ╚═(███)═╝
  ╚═(███)═╝
   ╚═(███)═╝
    ╚═(███)═╝
    ╚═(███)═╝
   ╚═(███)═╝
  ╚═(███)═╝
 ╚═(███)═╝
╚═(███)═╝
 ╚═(███)═╝
  ╚═(███)═╝
   ╚═(███)═╝
    ╚═(███)═╝
    ╚═(███)═╝
   ╚═(███)═╝
  ╚═(███)═╝
 ╚═(███)═╝
╚═(███)═╝
 ╚═(███)═╝
  ╚═(███)═╝
   ╚═(███)═╝
    ╚═(███)═╝
    ╚═(███)═╝
   ╚═(███)═╝
  ╚═(███)═╝
 ╚═(███)═╝
╚═(███)═╝
 ╚═(███)═╝
  ╚═(███)═╝
   ╚═(███)═╝
```

```command
$ millipede-go --reverse
 ╔═(███)═╗
  ╔═(███)═╗
   ╔═(███)═╗
    ╔═(███)═╗
    ╔═(███)═╗
   ╔═(███)═╗
  ╔═(███)═╗
 ╔═(███)═╗
╔═(███)═╗
 ╔═(███)═╗
  ╔═(███)═╗
   ╔═(███)═╗
    ╔═(███)═╗
    ╔═(███)═╗
   ╔═(███)═╗
  ╔═(███)═╗
 ╔═(███)═╗
╔═(███)═╗
 ╔═(███)═╗
  ╔═(███)═╗
    ╔⊙ ⊙╗
```

```command
$ millipede-go --opposite
    ╚⊙ ⊙╝
  ╚═(███)═╝
   ╚═(███)═╝
    ╚═(███)═╝
    ╚═(███)═╝
   ╚═(███)═╝
  ╚═(███)═╝
 ╚═(███)═╝
╚═(███)═╝
 ╚═(███)═╝
  ╚═(███)═╝
   ╚═(███)═╝
    ╚═(███)═╝
    ╚═(███)═╝
   ╚═(███)═╝
  ╚═(███)═╝
 ╚═(███)═╝
╚═(███)═╝
 ╚═(███)═╝
  ╚═(███)═╝
   ╚═(███)═╝
```

```command
$ millipede-go --skin=bocal --opposite --reverse
   ╔═(🐟🐟🐟)═╗
  ╔═(🐟🐟🐟)═╗
 ╔═(🐟🐟🐟)═╗
╔═(🐟🐟🐟)═╗
 ╔═(🐟🐟🐟)═╗
  ╔═(🐟🐟🐟)═╗
   ╔═(🐟🐟🐟)═╗
    ╔═(🐟🐟🐟)═╗
    ╔═(🐟🐟🐟)═╗
   ╔═(🐟🐟🐟)═╗
  ╔═(🐟🐟🐟)═╗
 ╔═(🐟🐟🐟)═╗
╔═(🐟🐟🐟)═╗
 ╔═(🐟🐟🐟)═╗
  ╔═(🐟🐟🐟)═╗
   ╔═(🐟🐟🐟)═╗
    ╔═(🐟🐟🐟)═╗
    ╔═(🐟🐟🐟)═╗
   ╔═(🐟🐟🐟)═╗
  ╔═(🐟🐟🐟)═╗
    ╔⊙ ⊙╗
```

```command
$ millipede-go --skin=ascii --width=7
    \o     o/
  |=(#######)=|
 |=(#######)=|
|=(#######)=|
 |=(#######)=|
  |=(#######)=|
   |=(#######)=|
    |=(#######)=|
    |=(#######)=|
   |=(#######)=|
  |=(#######)=|
 |=(#######)=|
|=(#######)=|
 |=(#######)=|
  |=(#######)=|
   |=(#######)=|
    |=(#######)=|
    |=(#######)=|
   |=(#######)=|
  |=(#######)=|
 |=(#######)=|
```

```command
$ millipede-go --skin=bocal --curve=8 --opposite
          ╚⊙ ⊙╝
        ╚═(🐟🐟🐟)═╝
       ╚═(🐟🐟🐟)═╝
      ╚═(🐟🐟🐟)═╝
     ╚═(🐟🐟🐟)═╝
    ╚═(🐟🐟🐟)═╝
   ╚═(🐟🐟🐟)═╝
  ╚═(🐟🐟🐟)═╝
 ╚═(🐟🐟🐟)═╝
╚═(🐟🐟🐟)═╝
 ╚═(🐟🐟🐟)═╝
  ╚═(🐟🐟🐟)═╝
   ╚═(🐟🐟🐟)═╝
    ╚═(🐟🐟🐟)═╝
     ╚═(🐟🐟🐟)═╝
      ╚═(🐟🐟🐟)═╝
       ╚═(🐟🐟🐟)═╝
        ╚═(🐟🐟🐟)═╝
       ╚═(🐟🐟🐟)═╝
      ╚═(🐟🐟🐟)═╝
     ╚═(🐟🐟🐟)═╝
```

```command
$ millipede-go --skin=bocal --curve=0
  ╚⊙ ⊙╝
╚═(🐟🐟🐟)═╝
╚═(🐟🐟🐟)═╝
╚═(🐟🐟🐟)═╝
╚═(🐟🐟🐟)═╝
╚═(🐟🐟🐟)═╝
╚═(🐟🐟🐟)═╝
╚═(🐟🐟🐟)═╝
╚═(🐟🐟🐟)═╝
╚═(🐟🐟🐟)═╝
╚═(🐟🐟🐟)═╝
╚═(🐟🐟🐟)═╝
╚═(🐟🐟🐟)═╝
╚═(🐟🐟🐟)═╝
╚═(🐟🐟🐟)═╝
╚═(🐟🐟🐟)═╝
╚═(🐟🐟🐟)═╝
╚═(🐟🐟🐟)═╝
╚═(🐟🐟🐟)═╝
╚═(🐟🐟🐟)═╝
╚═(🐟🐟🐟)═╝
```


```command
$ millipede-go --zalgo
  ╚⊙̲͋ ⊙̴╝͇͠
╚═(̀█̴█̴█̨)̛̻═͜╝͍͡
 ╚̘═̶(ͪ́█̛ͯ█҉█͚́)̷̰═͙͜╝̴̦
  ╚̨═(͉͋█͘█͜█̲͘)̷═̢╝̴
   ╚═̯(̴█͢█҉█͟)͢═̫͝╝͏
    ╚͝═(̨█͜█̶█̜͠)̯̉́═̢╝̧͎͛
   ╚̩ͩ═̵(ͦ█̴█̡█̤͟)̧̟═̸͕̓╝̕
  ╚͟═̨(͟█̞͠█̕█́͢)̪ͪ͟═̡͎╝̸
 ╚═͞(█̴̚█̨̾█̓͜)͜═̴╝̡
╚̲═̨(█͠█҉̲█̴̮̓)̸̭═̕╝̴͎
 ╚͏═҉̬(̀█̛̬ͥ█̶͕█̷)͝═͈͠╝́
  ╚̵═(█̚͘█́█̴͚)̷̱═̶̺̎╝̛͉
   ╚═́(̙̀█̨█̢█ͧ͘)̵═̸╝҉
    ╚̛͉̔═̶̻(█̡͔█̓͘█̢͐)̡̜═͘╝̀
   ╚̷═̫(̞͗█̢͇█̴͉█̶̥)̷═̤̑͡╝̝̏͡
  ╚═͕(ͭ█̛█̸█̢͇)͞═͎͠╝͞
 ╚̧═̀(҉̙█̍͏█̟͘█͏)̵̫═̡̱╝̷
╚═̸(̉͟█̸█͝█̢)͟═̴͎╝̼̑͞
 ╚═(̂█̵͕█͏█̛)̧═̶̿╝̖̉͞
  ╚̡═̴ͦ(̦█̨█̶̥█҉)̛̮═͡╝̴
   ╚═(̨█͢█̨█̸)̡═̕╝̛̋

```

---

```command
$ millipede-go --animate
```
![](https://raw.githubusercontent.com/getmillipede/millipede-go/master/assets/millipede.gif)

---

```command
$ millipede-go --animate --rainbow --curve=8 --width=5 --skin=humancentipede --chameleon 21
```
![](https://raw.githubusercontent.com/getmillipede/millipede-go/master/assets/millipede-full.gif)

---

```command
$ millipede-go --animate --zalgo
```
![](https://raw.githubusercontent.com/getmillipede/millipede-go/master/assets/millipede-zalgo.gif)

## Web usage

```command
$ millipede-web
2015/07/30 18:57:57 About to listen on 4242. Go to http://127.0.0.1:4242/
```

```command
$ curl localhost:4242
  ╚⊙ ⊙╝
╚═(███)═╝
 ╚═(███)═╝
  ╚═(███)═╝
   ╚═(███)═╝
    ╚═(███)═╝
   ╚═(███)═╝
  ╚═(███)═╝
 ╚═(███)═╝
╚═(███)═╝
 ╚═(███)═╝
  ╚═(███)═╝
   ╚═(███)═╝
    ╚═(███)═╝
   ╚═(███)═╝
  ╚═(███)═╝
 ╚═(███)═╝
╚═(███)═╝
 ╚═(███)═╝
  ╚═(███)═╝
   ╚═(███)═╝
$ curl localhost:4242/\?skin=bocal
╚⊙ ⊙╝
╚═(🐟🐟🐟)═╝
 ╚═(🐟🐟🐟)═╝
  ╚═(🐟🐟🐟)═╝
   ╚═(🐟🐟🐟)═╝
    ╚═(🐟🐟🐟)═╝
   ╚═(🐟🐟🐟)═╝
  ╚═(🐟🐟🐟)═╝
 ╚═(🐟🐟🐟)═╝
╚═(🐟🐟🐟)═╝
 ╚═(🐟🐟🐟)═╝
  ╚═(🐟🐟🐟)═╝
   ╚═(🐟🐟🐟)═╝
    ╚═(🐟🐟🐟)═╝
   ╚═(🐟🐟🐟)═╝
  ╚═(🐟🐟🐟)═╝
 ╚═(🐟🐟🐟)═╝
╚═(🐟🐟🐟)═╝
 ╚═(🐟🐟🐟)═╝
  ╚═(🐟🐟🐟)═╝
  ╚═(🐟🐟🐟)═╝
```

---

## Install

Install using Golang

```command
$ go get github.com/getmillipede/millipede-go/...
```
---

Install git version using homebrew

```console
$ brew tap getmillipede/millipede
$ brew install millipede-go --HEAD
```

---

Install a released version using homebrew

```console
$ brew tap getmillipede/millipede
$ brew install millipede-go
```

---

Download a statically compiled binary

- Go to [github.com/getmillipede/millipede-go/releases/latest](https://github.com/getmillipede/millipede-go/releases/latest)

See all releases: https://github.com/getmillipede/millipede-go/releases


## License

[MIT](https://github.com/getmillipede/millipede-go/blob/master/LICENSE)
