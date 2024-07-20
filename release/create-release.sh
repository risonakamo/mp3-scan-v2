# creates release. web dir needs to be npm i

set -exu
HERE=$(dirname $(realpath $BASH_SOURCE))
cd $HERE

# --- config
releaseName=mp3-scan-v2_1.0.0

workdir=$HERE/output/$releaseName
rm -rf output
mkdir -p $workdir

cd $HERE/..
bash dev/build-all.sh

cp -r mp3-scan.exe $workdir
cp -r config $workdir
cp version.md $workdir
mkdir -p $workdir/mp3-scan-v2-web
cp -r mp3-scan-v2-web/build $workdir/mp3-scan-v2-web

echo "done"