# YINI Demo Apps

This repository contains **example projects showing how to use YINI configuration files** with the official [yini-parser](https://github.com/YINI-lang/yini-parser-typescript).

The demos illustrate how to:
- Load and parse `.yini` files.  
- Access nested data.  
- Work with lists and inline objects.  
- Handle sections with spaces in their names.  
- Use both simple and advanced parsing options.  

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

- Learn more: https://yini-lang.org  
- Specification: https://github.com/YINI-lang/YINI-spec

---

## Basic demos

These examples show the **simplest possible way** to load and read a YINI file.

### JavaScript
- **JavaScript (CommonJS)** (using `require`):  
  [/js-cjs/basic](./js-cjs/basic/)
- **JavaScript (ESM)** (using `import`):  
  [/js-esm/basic](./js-esm/basic/)

### TypeScript
- **TypeScript (CommonJS)** (using `import`):  
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
- Three parsing variants:
  - Default parsing.  
  - Strict mode.  
  - Options object + metadata.

### JavaScript
- **JavaScript (CommonJS)** (using `require`):  
  [/js-cjs/medium](./js-cjs/medium/)
- **JavaScript (ESM)** (using `import`):  
  [/js-esm/medium](./js-esm/medium/)

### TypeScript
- **TypeScript (CommonJS)** (using `import`):  
  [/ts-cjs/medium](./ts-cjs/medium/)
- **TypeScript (ESM)** (using `import`):  
  [/ts-esm/medium](./ts-esm/medium/)

---

**^YINI ≡ A simple, structured, and human-friendly configuration format.**

- 🌐 Homepage: https://yini-lang.org  
- 🧠 Spec: https://github.com/YINI-lang/YINI-spec  
- 🧰 Parser: https://github.com/YINI-lang/yini-parser-typescript  
