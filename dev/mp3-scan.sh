set -exu
HERE=$(dirname $(realpath $BASH_SOURCE))
cd $HERE/..

exeName="mp3-scan"

go build -o $exeName.exe bin/$exeName/$exeName.go

set +u
if [[ "$1" == "run" ]]; then
    ./$exeName.exe
fi