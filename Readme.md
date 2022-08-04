# Go-Config

Here is an example and a template how to fetch config data from a YAML file and use it. 

## Before use

Import the YAML package: `go get gopkg.in/yaml.v3`

## Note

It is important to note that the model is case sensitive. If you customize the Config-YAML, then you have to note that in the 3 column `yaml: "Config-Directory"` (e.g. `yaml: "Config-Directory"`), the first letter of the config directory must be case-sensitive and the same as in the config.yaml file.

Further one must adapt the model if the Config.yaml file is adapted. Furthermore you have to adjust the getters in the config.go file.

It is still important that the configfile is excluded from a gitignore file so that no data is pushed.

## Resources

[Github: go-yaml/yaml](https://github.com/go-yaml/yaml)
