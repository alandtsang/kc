# kc
Simplified kubectl implemented by golang

## Usage
```
kc controls the Kubernetes cluster manager.

Find more information at: https://github.com/alandtsang/kc

Usage:
  kc [command]

Available Commands:
  apply       Apply a configuration to a resource by filename or stdin.
  delete      Delete resources by filenames, stdin, resources and names, or by resources and label selector
  get         Display one or many resources
  help        Help about any command
  logs        Print the logs for a container in a pod or specified resource
  version     A brief description of your command

Flags:
      --config string   config file (default is $HOME/.kc.yaml)
  -h, --help            help for kc
  -t, --toggle          Help message for toggle

Use "kc [command] --help" for more information about a command.
```

```
kc k8s-test2 get ns
kc k8s-test2 get -n logging pod
```