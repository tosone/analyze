### Worklyzer
------

Colloction information from wakatime.com and analyze the information.

10 minutes travel all of the users and get all of the information work status since last request wakatime.

#### How to run

``` bash
worklyzer-linux --config=config.yaml service
```

``` bash
worklyzer-linux --help 
Usage:
  worklyzer [command]

Available Commands:
  help        Help about any command
  service     Fetch information from wakatime
  version     get version

Flags:
  -f, --config string   config file (default "/etc/analyze/config.yaml")
  -h, --help            help for worklyzer

Use "worklyzer [command] --help" for more information about a command.

```

Default config.yaml is `/etc/worklyzer/config.yaml`.

#### TODO

- Add fontend to display the information.
