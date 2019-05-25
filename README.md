# NagiosXI CLI

Command Line Interface (CLI) to interact with NagiosXI API.

[![Build Status](https://travis-ci.com/romainDavaze/nagiosxi-cli.svg?branch=master)](https://travis-ci.com/romainDavaze/nagiosxi-cli)

## Configuration

### Global

When starting, nagiosxi-cli reads a config file and load variables from it.

By default, if the parameter `--config` is not specified, it looks for a `.nagiosxi-cli.yaml` file under user's home directory.

You can see an example [here](examples/nagiosxi-cli.yaml).

### NagiosXI objects

You can add and delete several NagiosXI objects by providing a specific file as a parameter.

As of now, these objects are supported :
- Hosts ([example](examples/hosts.yaml))
- Services ([example](examples/services.yaml))


## Usage

```
# To add services
nagiosxi-cli --config nagiosxi-cli.yaml service add -f services.yaml
```


## Author

<a href="https://romaindavaze.github.io/">Romain Davaze</a>