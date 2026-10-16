# Go Medium Demo

This directory contains a medium-sized demo showing how to read a [YINI configuration file](./settings.yini) with the [official Go parser](https://github.com/YINI-lang/yini-parser-go).

The example in [main.go](./main.go) reads nested sections, arrays, an inline object, and a section with a space in its name. It parses the same file in the default lenient mode and in strict mode, then prints selected values and formatted JSON. [out.txt](./out.txt) contains sample output.

The file has one top-level section and ends with `/END`, as required by the Go parser's strict mode. Strict mode is selected explicitly in `main.go`.

## Usage

Run all commands below from this demo directory. Go 1.22 or newer is required.

### Installation

```bash
go mod download
```

The module dependency was added with `go get github.com/YINI-lang/yini-parser-go@latest`. Its resolved version is recorded in `go.mod` and `go.sum`.

### Run the Example

```bash
go run .
```

### Optional: Save an Executable

`go run .` does not save an executable in this directory. To build one and run it from PowerShell:

```powershell
go build -o demo.exe .
.\demo.exe
```

On macOS or Linux, use `go build -o demo .` and `./demo` instead. Run the executable from this directory so it can find `settings.yini`.

---

[YINI Homepage](https://yini-lang.org/?utm_source=github&utm_medium=referral&utm_campaign=yini_demos&utm_content=sub_readme_footer)
