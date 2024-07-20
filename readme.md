# Mp3 Scan v2
Tool for reviewing mp3 files.

# Dev
## Requirements
- golang 1.22+
- nodejs, last tested on v20.10.0
- pnpm
- cygwin to run bash scripts

## Initialisation
1. Submod update
2. Ensure node env activated
2. In `mp3-scan-v2-web`, `pnpm i`

## Deving
Need to run 2 components.

After building each component once, only need to run build/watch for the component that you are actively working on (i.e, don't need to rebuild exe if only working on web, and don't need to rebuild web if working on exe).

## Frontend
1. In `mp3-scan-v2-web`, `pnpm watch`

`pnpm build` to build without watching.

## Backend (main exe)
1. `bash dev/mp3-scan.sh run`
    - If run without `run`, builds without running

Program will run and open web ui

## Build all
`bash dev/build-all.sh` to build all components.