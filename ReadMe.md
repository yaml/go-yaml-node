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
Kind: Document
More:
    Kind: Mapping
    More:
        Kind: Scalar
        Text: "scalar forms"

        Kind: Sequence
        Styl: Flow
        More:
            Kind: Scalar
            Tag : !!null
            Text: "null"

            Kind: Scalar
            Tag : !!bool
            Text: "true"

            Kind: Scalar
            Text: "yes"

            Kind: Scalar
            Tag : !!int
            Text: "0xcafe"

            Kind: Scalar
            Tag : !!int
            Text: "0o123"

            Kind: Scalar
            Text: "12:34:56"
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

This project is licensed under the MIT License.

See the [LICENSE](LICENSE) file for details.
