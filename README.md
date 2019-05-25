# Nagios-cli

Command Line Interface (CLI) to interact with Nagios API.

## Configuration

### Global

When starting, nagios-cli reads a config file and load variables from it.

By default, if the parameter `--config` is not specified, it looks for a `.nagios-cli.yaml` file under user's home directory.

You can see an example [here](examples/nagios-cli.yaml).

### Nagios objects

You can add and delete several Nagios objects by providing a specific file as a parameter.

As of now, these objects are supported :
- Hosts ([example](examples/hosts.yaml))
- Services ([example](examples/services.yaml))


## Usage

```
# To add services
nagios-cli --config nagios-cli.yaml service add -f services.yaml
```


## Author

<a href="https://romaindavaze.github.io/">Romain Davaze</a>