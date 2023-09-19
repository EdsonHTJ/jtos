Copy code
# [jtos](https://github.com/EdsonHTJ/jtos)

`jtos` is a command-line application for processing json input files and generating the struct code for the chosen language. Currently, only Go is supported as a language generator.

## Prerequisites

- Go (version 1.21 or newer)

## Installation

Before you can use `jtos`, you need to install it. The currently only installation possible is through go install:

```
go install github.com/EdsonHTJ/jtos/cmd/jtos@v0.1.1
```

## Usage

### Basic Usage

```
jtos [path/to/input/file]
```

With the basic usage, the input file is taken from the argument, and the output file defaults to a directory and file named after the input filename without its extension. The generator defaults to "go".

### Advanced Usage

```
jtos --input="path/to/input/file" --out="path/to/output/file" --gen="generator_type"
```

- `-g, --gen string`: The generator type to use. (default "go")
- `-h, --help`: Help for jtos.
- `-o, --out string`: Path to the output directory. (default "./")
- `-p, --package string`: Name of the output package.
- `-s, --struct string`: Name of the output structure.

For example, for an input file at `/etc/lib/test.txt`, the default output will be `test/test.go`.

### Examples

```bash
jtos example.json
# Output will be in example/example.go
```

```bash
jtos --input=sample.json --out=output/path --package=samplePkg --struct=SampleStruct
# Adjusted output, package name, and struct name.
```

## Contribution

Contributions are welcome! Please create an issue or open a pull request if you have changes or improvements.

## License

This project is licensed under the MIT License.
