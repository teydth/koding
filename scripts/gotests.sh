#!/bin/bash

set -o errexit

#  make relative paths work.
cd $(dirname $0)/..
COVERAGE_FOLDER="./go/src"
COVMERGE="./go/bin/gocovmerge"
KD=$(pwd)

action="sh -c"
# action="echo"

# compile uses -coverpkg flag to test with coverage
# this satisfies to cover the subpackages of tests
function compile () {
    # list the packages
    export PKGS=$(go list $1 | grep -v /vendor/)

    # make comma-separated
    export PKGS_DELIM=$(echo "$PKGS" | paste -sd "," -)

    go list -f '{{if len .TestGoFiles}}"go test -v -cover -c {{.ImportPath}} -o={{.Dir}}/{{.Name}}.test -coverpkg='$PKGS_DELIM' "{{end}}' $PKGS | grep -v vendor | xargs -L 1 -I{} $action "{}$COMPILE_FLAGS"
}

# run runs the binary file that created before (with compile function)
# this function runs the binary with given RUN_FLAGS
# RUN_FLAGS values might be config file or any value if required like
# ...('-kite-init' required for collaboration tests)
function run () {
    go list -f '{{if len .TestGoFiles}}"{{.Dir}}/{{.Name}}.test -test.coverprofile={{.Dir}}/coverage.txt "{{end}}' $1 | grep -v vendor | xargs -L 1 -I{} $action "{}$RUN_FLAGS"
}

function merge () {
    folder=$(createFolder $1)

    go list -f '{{if len .TestGoFiles}}"{{.Dir}}/coverage.txt"{{end}}' $1 | grep -v vendor | xargs $COVMERGE > "$folder"
}

# runWithCD runs like run function.
# But this function changes the directory while running.
function runWithCD () {
    go list -f '{{if len .TestGoFiles}}"cd {{.Dir}} && ./{{.Name}}.test -test.coverprofile={{.Dir}}/coverage.txt "{{end}}' $1 | grep -v vendor | xargs -L 1 -I{} $action "{}$RUN_FLAGS"
}

# clean removes the .tests files after creation
function clean () {
    go list -f '{{if len .TestGoFiles}}"rm {{.Dir}}/{{.Name}}.test "{{end}}' $1 | xargs -L 1 $action
}

function runAll () {
    compile $@
    run $@
    merge $@
    removeCoverProfiles $@
    clean $@
}

# removeCoverProfiles deletes the coverage.txt files after merging all
# coverages into 1 coverage.txt(merged coverage file won't be deleted)
function removeCoverProfiles () {
    # folder=$(createFolder $1)
    name=$(generateFolderName $1)
    mkdir -p "$KD/go/src/coverages"
    mkdir -p "$KD/go/src/coverages/$name"
    touch "$KD/go/src/coverages/$name/coverage.txt"
    folder="$KD/go/src/coverages/$name/coverage.txt"

    go list -f '{{if len .TestGoFiles}}"{{.Dir}}/coverage.txt"{{end}}' $1 | grep -v vendor | xargs rm
}

# commandExists checks if the given parameter command is exits or not
function commandExists() {
    command -v $1 >/dev/null 2>&1
}

# generateFolderName generates a random string with given parameter
# according to given parameter & command checking
function generateFolderName() {
    if commandExists md5 ; then
        echo $1 | md5
    elif commandExists md5sum ; then
        echo $1 | md5sum
    else
        date +%s%N
    fi
}

# createFolder creates a '/coverage' folder if doesn't exists
# then appends the generated random string as folder name
# e.g:
# assume random string : $generatedName
# path will be -> koding/go/src/coverage/$generatedName
function createFolder() {
    name=$(generateFolderName $1)
    mkdir -p "$KD/go/src/coverages"
    mkdir -p "$KD/go/src/coverages/$name"
    touch "$KD/go/src/coverages/$name/coverage.txt"
    folder="$KD/go/src/coverages/$name/coverage.txt"
    echo $folder
}

if  [ "$1" == "kites" ]; then
    shift
    export COMPILE_FLAGS=${COMPILE_FLAGS:-"-race"}

    compile  $@
    runWithCD $@
    clean $@

elif [ "$1" == "socialapi" ]; then
    shift
    export RUN_FLAGS=${RUN_FLAGS:-"-c=./go/src/socialapi/config/dev.toml"}

    runAll $@

elif [ "$1" == "generate" ]; then
    shift

    generateFolderName $@
else
    runAll $@
fi
