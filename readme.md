# Mp3 Scan v2
Tool for reviewing mp3 files.

# Functionality
Mp3 Scan detects mp3 files inside a target folder and presents them to the user via a web UI. The user can listen to each mp3 and make a decision on each one - "yes", "no", or "maybe". Mp3 Scan will then move the file into an accordingly named folder, **next to where the mp3 was originally located**. This is to keep all mp3s that belong together in an album in the same place, so they can be tagged later.

This is useful for when you want to listen to many mp3s and choose certain ones to keep.

Mp3 scan will ignore all mp3 files that already have a "decision" -  if it is in one of the decision folders. Thus, you can resume a review session without re-reviewing already reviewed items.

# Usage
## Configuration
Mp3 Scan must be pointed at a folder that has mp3 files. Edit `config/config.yml` and enter path to a folder.

The folder can have any internal folder structure - all mp3 files don't need to be at the top level.

When making a decision, Mp3 Scan will move the item into a folder named after the decision. Mp3 scan will ignore items that are already in a decision folder, except for "maybe", if "includeMaybe" is enabled in the config file.

## Doing Reviewing
Run `mp3-scan.exe` to open the web UI.

![](todo) picture of the web ui

Click "Open Item" to open the mp3 file with your default program for mp3 files.

To make a decision, click one of the decision options, then click "Next Item" to continue.

!! "Note" note
    It might fail to move the item if you still have it open in your default program.

Since items are moved into decision folders, they won't be detected again if you re-run the program later, so you can close the program without finishing reviewing all items.

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