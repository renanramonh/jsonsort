[![Maintenance](https://img.shields.io/badge/Maintained%3F-yes-green.svg)](https://bitbucket.org/lbesson/ansi-colors)
[![Generic badge](https://img.shields.io/badge/Status-Alpha-red.svg)]()
[![GPLv3 license](https://img.shields.io/badge/License-Apache-purple.svg)](https://www.apache.org/licenses/LICENSE-2.0)
[![Ask Me Anything !](https://img.shields.io/badge/Ask%20me-anything-1abc9c.svg)](https://github.com/renanramonh)


# jsonsort

The jsonsort is a command-line tool written in Go that allows you to sort JSON files based on keys. It provides a simple and efficient way to sort JSON data.


## features
- Support for sorting json files in ascending order.
- Simple and intuitive command-line interface.

## Installation

To install `jsonsort`, follow these steps:

1. Make sure you have Go installed on your system. You can download and install it from the official website: https://golang.org/dl

2. Run the following command to install the JSON Sorter CLI:
```bash
go install github.com/renanramonh/jsonsort
```

3. The executable binary will be placed in your Go binary directory (`$GOPATH/bin`) after this command downloads the source code and compiles it..

4. (Optional) To facilitate access, add the Go binary directory to your system's PATH environment variable.
Edit your shell configuration file (`~/.bashrc`, `~/.bash_profile`, or `~/.zshrc`) and add the following command: ``export PATH=$PATH:$GOPATH/bin``.
Save the file and restart your terminal or run source `~/.bashrc` (or the corresponding command for your shell) to apply the changes.

## Usage
To sort a JSON file using `jsonsort`, use the following command:
```bash
jsonsort [file]
```

Replace `[file]` with the path to the JSON file you want to sort.

Here is an example command to sort a JSON file:
```bash
jsonsort ./input.json
```
the output file with the sorted json will be saved on `./input_sorted.json`

## Future Improvements
- Sort JSON files based on specific keys or fields. 
- Support for sorting in ascending or descending order.
- Format json output file with indentation
