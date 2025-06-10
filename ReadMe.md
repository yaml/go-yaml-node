go-yaml-node
============

Load a YAML file and print the nodes


## Description

This program is extremely useful for understanding / debugging how
[go-yaml](https://github.com/yaml/go-yaml) loads YAML documents.

The `Unmarshal` function loads YAML documents into a `Node` data type.

This program will read YAML from stdin and print the nodes in an indented tree
format, showing all the (non-empty) field values.


## Usage

Example:

```
$ go-yaml-node <<< 'scalar forms: [null, true, yes, 0xcafe, 0o123, 12:34:56]'
kind: Document
content:
    - kind: Mapping
      content:
        - kind: Scalar
          text: scalar forms
        - kind: Sequence
          style: Flow
          content:
            - kind: Scalar
              tag: '!!null'
              text: "null"
            - kind: Scalar
              tag: '!!bool'
              text: "true"
            - kind: Scalar
              text: "yes"
            - kind: Scalar
              tag: '!!int'
              text: "0xcafe"
            - kind: Scalar
              tag: '!!int'
              text: "0o123"
            - kind: Scalar
              text: "12:34:56"
```


## Installation

```
make install PREFIX=<prefix-dir-path>
```


## Testing

```
make test
make test file=test/file-002.yaml
```

No deps required.
On first run this will install Go locally.

This should install all deps including a local installation of Go.


## Installation

```
go install github.com/ingydotnet/go-yaml-node@latest
```


## License

Copyright 2025 - Ingy döt Net

This project is licensed under the Apache License, Version 2.0.

See the [LICENSE](LICENSE) file for details.
