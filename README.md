# gituser
[![Build Status](https://travis-ci.org/WindomZ/gituser.svg?branch=master)](https://travis-ci.org/WindomZ/gituser)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](https://opensource.org/licenses/MIT)
[![Go Report Card](https://goreportcard.com/badge/github.com/WindomZ/gituser)](https://goreportcard.com/report/github.com/WindomZ/gituser)

> Easily switch git users.

![v1.2.0](https://img.shields.io/badge/version-v1.2.0-blue.svg)
![status](https://img.shields.io/badge/status-stable-green.svg)

## Installation

To get the package, execute:

```bash
go get github.com/WindomZ/gituser
```

## Usage
```bash
$ gituser -h

  Usage:
    gituser add <user> <name> <email> [--private-github]
    gituser (remove|rm) <user>
    gituser (list|ls)
    gituser set <user> [--private-github]
    gituser unset
    gituser -h|--help
    gituser -v|--version

  Options:
    --private-github
                  private email address for GitHub
    -h --help     output usage information
    -v --version  output the version number

  Commands:
    add           add user configuration
    remove|rm     remove user configuration
    list|ls       list user configuration
    set           set local git user configuration from <user>
    unset         unset local git user configuration

  Argument:
    <user>        the name of the configure user information
    <name>        the name of the git username
    <email>       the address of the git email
```

## Example

```bash
gituser add windomz WindomZ WindomZ@users.noreply.github.com  # Save 'WindomZ' into configuration file
gituser add windomz WindomZ '' --private-github               # Ibid, but email is GitHub privacy address
gituser rm windomz                                            # Delete 'WindomZ' from configuration file
gituser list                                                  # List all saved users
gituser set windomz                                           # Set local git user and email
gituser unset                                                 # Unset local git user and email
```

## PowerBy

[go-commander](https://github.com/WindomZ/go-commander) - The solution for building command shell programs

## Related

[WindomZ/gituser.js](https://github.com/WindomZ/gituser.js) - Written in **Node.js**

## License

The [MIT License](https://github.com/WindomZ/gituser/blob/dev/LICENSE)
