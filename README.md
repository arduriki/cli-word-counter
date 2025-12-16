# CLI Word Counter

A simple command-line tool written in Go that counts words or lines from standard input.

## What I Learned

Building this CLI tool helped me understand several key Go concepts:

- **Flag Parsing**: Using the `flag` package to handle command-line arguments
- **I/O Operations**: Working with `io.Reader` interface for flexible input handling
- **Buffered Scanning**: Using `bufio.Scanner` to efficiently read and process text
- **Scanner Split Functions**: Customizing how the scanner splits input (words vs lines)
- **Standard Input**: Reading from `os.Stdin` to create a pipe-friendly CLI tool

## Features

- Count words from standard input (default behavior)
- Count lines with the `-l` flag
- Works with pipes and file redirection

## Installation

```bash
go build -o wordcount
```

## Usage

### Count words:
```bash
echo "Hello world from Go" | ./wordcount
# Output: 4
```

### Count lines:
```bash
cat file.txt | ./wordcount -l
```

### Using with echo and newlines:
```bash
echo -e "Line 1\nLine 2\nLine 3" | ./wordcount -l
# Output: 3
```

## Flags

- `-l`: Count lines instead of words

## How It Works

1. The program accepts input from `stdin`
2. It uses `bufio.Scanner` to read the input
3. By default, it splits by words using `bufio.ScanWords`
4. When `-l` flag is provided, it splits by lines (default scanner behavior)
5. Returns the total count