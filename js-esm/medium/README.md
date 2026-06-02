# JavaScript Medium Demo - ESM

This directory contains a medium-sized demo showing how to use a [YINI config](https://github.com/YINI-lang) file with [yini-parser](https://github.com/YINI-lang/yini-parser-typescript) in a JavaScript project using `import` (ESM).

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

The example includes [settings.yini](./settings.yini), [index.js](./index.js), and captured sample output in [out.txt](./out.txt).

## Usage

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
[YINI Homepage](https://yini-lang.org/?utm_source=github&utm_medium=referral&utm_campaign=yini_demos&utm_content=sub_readme_footer)  
