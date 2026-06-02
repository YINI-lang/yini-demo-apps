# YINI Demo Apps

This repository contains **example projects showing how to use YINI configuration files** with the official YINI parsers for Python and TypeScript/JavaScript.

Official parsers used by these demos:
- [yini-parser-python](https://github.com/YINI-lang/yini-parser-python)
- [yini-parser-typescript](https://github.com/YINI-lang/yini-parser-typescript)

These examples are intentionally kept **clear, minimal, and practical** so they are easy to read, learn from, and adapt.

Feel free to browse, copy, or adapt the code if you find it useful.

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

### Python
- **Python 3**:
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

Some JavaScript and TypeScript examples may also demonstrate additional parser options such as strict mode, metadata, or options objects.

### Python
- **Python 3**:
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

- Homepage: [https://yini-lang.org](https://yini-lang.org/?utm_source=github&utm_medium=referral&utm_campaign=yini_demos&utm_content=readme_footer)
- Spec: https://github.com/YINI-lang/YINI-spec
- TypeScript Parser: https://github.com/YINI-lang/yini-parser-typescript
- Python Parser: https://github.com/YINI-lang/yini-parser-python
