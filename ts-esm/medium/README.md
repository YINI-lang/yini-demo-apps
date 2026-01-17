# TypeScript Medium Demo - ESM

This directory contains a simple demo showing how to use a [YINI config](https://github.com/YINI-lang) file with [yini-parser](https://github.com/YINI-lang/yini-parser-typescript) in a TypeScript project using `import` (ESM).

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

The example includes a sample `settings.yini` file and a small TypeScript script that reads and prints its contents.

## Usage
This example is located in `/ts-esm/medium/index.ts`.

### Installation
Run in your terminal:
```bash 
npm install
```

### Run the Example
```bash
npm run build    # compiles to dist/index.js
npm start        # runs the compiled JS
```

---
[YINI Homepage](https://yini-lang.org)  
