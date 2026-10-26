# YINI Config Demo Apps

This repository contains **example projects showing how to use YINI configuration files** with the official YINI parsers for Go, Python, and TypeScript/JavaScript.

Official parser versions used by these demos:
- [yini-parser](https://pypi.org/project/yini-parser/1.0.0rc1/) for Python: `1.0.0rc1`
- [yini-parser](https://www.npmjs.com/package/yini-parser/v/1.7.0) for JavaScript and TypeScript: `1.7.0`
- [yini-parser-go](https://github.com/YINI-lang/yini-parser-go) for Go: `v1.0.0-alpha.3`

These examples are intentionally kept **clear, minimal, and practical** so they are easy to read, learn from, and adapt.

Feel free to browse, copy, or adapt the code if you find it useful.

---

## Getting started

Choose a demo below, then follow the installation and run instructions in that demo's `README.md`. Run its commands from the demo directory. The Go demos require Go 1.22+, the JavaScript and TypeScript demos require Node.js 20+, and the Python demos require Python 3.10+.

If you use VS Code, install the [YINI Syntax Highlighting extension](https://marketplace.visualstudio.com/items?itemName=yini-lang.yini-syntax-highlighting) to highlight `.yini` files.

---

## What is YINI?

YINI is a human-readable text format for representing structured information. It is designed to be **clear, predictable, and easy for humans to read and write**, while still being precise for machines to process.

It is suitable for:
- Configuration files.
- Application settings.
- Structured data files.
- Developer-friendly alternatives to JSON, XML, YAML, or TOML.

- Learn more: [https://yini-lang.org](https://yini-lang.org/?utm_source=github&utm_medium=referral&utm_campaign=yini_demos&utm_content=readme)
- Specification: https://github.com/YINI-lang/YINI-spec

---

## Basic demos

These examples show the **simplest possible way** to load and read a YINI file.

### Go
- **Go 1.22+**:
  [/go/basic](./go/basic/)

### Python
- **Python 3.10+**:
  [/python/basic](./python/basic/)

### JavaScript
- **JavaScript (CommonJS)** (using `require`):
  [/js-cjs/basic](./js-cjs/basic/)
- **JavaScript (ESM)** (using `import`):
  [/js-esm/basic](./js-esm/basic/)

### TypeScript
- **TypeScript (CommonJS)** (compiled to CommonJS):
  [/ts-cjs/basic](./ts-cjs/basic/)
- **TypeScript (ESM)** (using `import`):
  [/ts-esm/basic](./ts-esm/basic/)

---

## Medium demos

These examples demonstrate:
- Nested sections.
- Inline objects.
- Arrays.
- Sections with spaces in their names.

The Go example also demonstrates default lenient parsing and strict mode. Some JavaScript and TypeScript examples demonstrate additional parser options such as metadata or options objects.

### Go
- **Go 1.22+**:
  [/go/medium](./go/medium/)

### Python
- **Python 3.10+**:
  [/python/medium](./python/medium/)

### JavaScript
- **JavaScript (CommonJS)** (using `require`):
  [/js-cjs/medium](./js-cjs/medium/)
- **JavaScript (ESM)** (using `import`):
  [/js-esm/medium](./js-esm/medium/)

### TypeScript
- **TypeScript (CommonJS)** (compiled to CommonJS):
  [/ts-cjs/medium](./ts-cjs/medium/)
- **TypeScript (ESM)** (using `import`):
  [/ts-esm/medium](./ts-esm/medium/)

---

**^YINI = A simple, structured, and human-friendly configuration format.**

- Homepage: [https://yini-lang.org](https://yini-lang.org/?utm_source=github&utm_campaign=yini_demos&utm_medium=referral&utm_content=readme_footer)
- Spec: https://github.com/YINI-lang/YINI-spec
- Go Parser: https://github.com/YINI-lang/yini-parser-go
- TypeScript Parser: https://github.com/YINI-lang/yini-parser-typescript
- Python Parser: https://github.com/YINI-lang/yini-parser-python
