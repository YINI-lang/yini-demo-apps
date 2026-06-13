# AGENTS.md

> AI agent instructions for this repository.
> Read this before making changes.
> If any instruction in this file is unclear, ambiguous, or conflicts with the repository state, stop and ask the human maintainer before proceeding.

## Project Overview

- **Name:** YINI Demo Apps
- **Purpose:** Small example projects that show how to read YINI configuration files with the official Python and JavaScript/TypeScript parsers.
- **Languages/runtimes:** Python 3, JavaScript on Node.js, TypeScript 5.9.x on Node.js.
- **Package managers:** npm for JavaScript/TypeScript demos; pip with `requirements.txt` for Python demos.
- **Framework/stack:** No application framework. These are minimal console demos.
- **Repository type:** Multi-demo repository. There is no root package.json, no root build, and no shared test runner.

## Repository Structure

```text
.
|-- js-cjs/
|   |-- basic/      # JavaScript CommonJS basic demo
|   `-- medium/     # JavaScript CommonJS medium demo
|-- js-esm/
|   |-- basic/      # JavaScript ESM basic demo
|   `-- medium/     # JavaScript ESM medium demo
|-- ts-cjs/
|   |-- basic/      # TypeScript compiled to CommonJS
|   `-- medium/     # TypeScript compiled to CommonJS
|-- ts-esm/
|   |-- basic/      # TypeScript compiled to ESM
|   `-- medium/     # TypeScript compiled to ESM
|-- python/
|   |-- basic/      # Python basic demo
|   `-- medium/     # Python medium demo
|-- README.md       # Repository overview
`-- LICENSE
```

Each demo directory is intentionally self-contained. Prefer changing only the matching demo when a task is specific to one language, module format, or difficulty level.

## Demo Contents

For the exact contents, purpose, setup, and usage of a demo, read that demo directory's `README.md` before making changes or running commands. The notes below are only a repository-level summary.

- `basic` demos use `config.yini` and show the simplest parse/read flow.
- `medium` demos use `settings.yini` and show nested sections, arrays, inline objects, sections with spaces, and parser options where supported.
- JavaScript demos use `index.js`.
- TypeScript demos use `index.ts`, `tsconfig.json`, and compile to `dist/`.
- Python demos use `main.py`.
- Medium JavaScript and TypeScript demos include `out.txt` with captured sample output.

## Commands

Run commands from the demo directory you are working in, and use that demo directory's `README.md` as the source of truth for installation, build, and run instructions.

Do not assume commands from one demo apply to another demo. The JavaScript, TypeScript, Python, CommonJS, ESM, `basic`, and `medium` examples may have different instructions.

## Required Checks

There is currently no automated test suite, lint command, or root build command.

Before considering a change complete, run the smallest relevant check:
- For code changes: run the build or run command documented in the affected demo directory's `README.md`.
- For documentation-only changes: proofread the changed files and verify links or commands when they were changed.

When changing shared behavior across multiple equivalent demos, run the relevant README-documented checks in each affected demo directory.

If a check cannot be run because dependencies are missing or the environment is unavailable, say exactly what blocked it and what was validated instead.

## Code Style

Follow the existing style of the demo being edited.

General rules:
- Prefer clear, simple, maintainable code.
- Keep examples easy to read for someone learning YINI.
- Prefer explicit names and straightforward control flow over compact or clever code.
- Do not rewrite unrelated demos.
- Do not reformat files unnecessarily.
- Keep public parser usage examples stable unless the task explicitly asks for a change.
- Preserve CommonJS vs ESM differences in the matching directories.
- Preserve TypeScript CommonJS vs ESM compiler settings unless the task is specifically about module behavior.
- Keep YINI examples human-readable and deterministic.

## YINI Demo Guidance

When editing `.yini` files or code that reads them:
- Preserve the distinction between `basic` and `medium` examples.
- Keep syntax examples aligned with the current YINI specification.
- Keep parser behavior explicit; avoid hidden defaults or magical helper layers in these demos.
- Preserve strict-mode and lenient/default-mode examples where they are demonstrated.
- Prefer precise error messages and comments when a demo intentionally shows parser options.
- Update the matching README when behavior, commands, file names, or demonstrated parser options change.
- Update `out.txt` only when the corresponding medium demo output intentionally changes.

## Dependencies

- JavaScript and TypeScript demos use `yini-parser` from npm.
- TypeScript demos also use `typescript`.
- Python demos use `yini-parser` from `requirements.txt`.

Do not add new runtime dependencies unless they are clearly necessary for the demo. Before adding a dependency, prefer the existing parser APIs, standard library functionality, or a small local helper. If dependency changes are required, update the corresponding lockfile or requirements file in the affected demo only.

## Safety and Scope Boundaries

### Always Do

- Run tests before submitting any change.
- Match the code patterns in the file you are editing.
- Keep changes focused — one concern per PR.
- When editing Markdown files, if a line introduces a bulleted list and ends with a colon (`:`), place the first bullet immediately on the next line. Do not insert a blank line between the introductory line and the first bullet.

### Ask First

- Before adding a new dependency.
- Before changing the public API or exported types.
- Before modifying CI/CD configuration.
- Before refactoring shared utilities used across multiple modules.

### Never Do

Do not modify the files or dirs specified in `.gitignore`, unless the task specifically requires it.

Do not perform destructive operations such as resetting history, deleting large parts of the repository, force-pushing, creating releases, or publishing packages.

Do not create commits, tags, branches, or releases unless explicitly requested.
