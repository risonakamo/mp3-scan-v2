set -exu
HERE=$(dirname $(realpath $BASH_SOURCE))
cd $HERE

bash mp3-scan.sh

cd ../mp3-scan-v2-web
rm -rf build
pnpm build