# NagiosXI CLI

Command Line Interface (CLI) to interact with NagiosXI API.

[![Build Status](https://travis-ci.com/romainDavaze/nagiosxi-cli.svg?branch=master)](https://travis-ci.com/romainDavaze/nagiosxi-cli)

## Configuration

### Global

When starting, nagiosxi-cli reads a config file and load variables from it.

By default, if the parameter `--config` is not specified, it looks for a `.nagiosxi-cli.yaml` file under user's home directory.
You can see an example [here](examples/nagiosxi-cli.yaml).

**The API token must belong to a NagiosXI admin.**


### NagiosXI objects

You can add and delete several NagiosXI objects by providing a specific file as a parameter.

As of now, these objects are supported :
- Commands ([example](examples/commands.yaml))
- Hosts ([example](examples/hosts.yaml))
- Services ([example](examples/services.yaml))



## Usage

```
# To add services
nagiosxi-cli add services -f services.yaml
```


## Author

<a href="https://romaindavaze.github.io/">Romain Davaze</a>