#!/bin/sh

export FLATNVIM_EDITOR="nvim"
export FLATNVIM_EXTRA_COMMAND=""

testdir="testing"
[ -d $testdir ] && ! rm -r $testdir && { echo "setup error"; exit 2; }
mkdir $testdir

testname="editor not in path"
FLATNVIM_EDITOR="command_not_in_path" ./bin/flatnvim 2>/dev/null &&
    { echo "FAIL - $testname"; exit 1; }
echo "pass - $testname"

testname="pass through"
testfile="$testdir/passthrough.txt"
echo "w $testfile" | ./bin/flatnvim -es 2>/dev/null
[ ! -e $testfile ] && { echo "FAIL - $testname"; exit 1; }
echo "pass - $testname"

testname="embedded"
testfile="$testdir/embedded.txt"
FLATNVIM_EXTRA_COMMAND="wq" ./bin/flatnvim --cmd "term ./bin/flatnvim $testfile" 2>/dev/null
[ ! -e $testfile ] && { echo "FAIL - $testname"; exit 1; }
echo "pass - $testname"

echo
echo "all tests passed"
