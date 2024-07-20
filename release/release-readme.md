1. Update `version.md` with changelog
2. Edit `create-release.sh`. Update version var in config section
3. Web dir needs to already by `pnpm i`
4. Activate node env
5. `bash create-release.sh`
6. Check the dir created in `output`
7. Zip as zip file and place somewhere
8. Git commit, tag, push
9. Create release. Use same text as in change log
10. Upload zip