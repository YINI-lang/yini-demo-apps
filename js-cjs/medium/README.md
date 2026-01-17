# JavaScript Medium Demo - CommonJS

This directory contains a simple demo showing how to use a [YINI config](https://github.com/YINI-lang) file with [yini-parser](https://github.com/YINI-lang/yini-parser-typescript) in a JavaScript project using `require` (CommonJS).

Demonstrates three ways to parse a YINI file:
- Variant A: Default parsing
- Variant B: Strict mode
- Variant C: Options object + metadata

This example shows how to:
- Load and parse settings.yini
- Read nested values
- Work with arrays and inline objects
- Access sections with spaces in their names
- Inspect both the JS object and JSON output

The example includes a sample [settings.yini](./settings.yini) file and a small JavaScript script that reads and prints its contents [out.txt](./out.txt).

## Usage
This example is located in `/js-cjs/medium/index.js`.

### Installation
Run in your terminal:
```bash 
npm install
```

### Run the Example
```bash
npm start
```

---
