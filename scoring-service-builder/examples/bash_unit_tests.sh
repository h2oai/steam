#!/usr/bin/env bash

# Unit test methods for bash

function assert_equals_string  { # assert_equals_string("echo hej", "hej")
  testname="$1"
  cmd="$2"
  expected="$3"
  reference="$4"
  res=`bash -c "$cmd"`
  if [ "$res" != "$expected" ] 
  then
    echo
    echo "Test FAILED: $testname"
    echo "cmd: $cmd"
    echo "expected: $expected"
    echo "found: $res"
    echo "in file $reference"
    echo
    have_failed=1
  else 
    echo "Test OK: $testname"
  fi
}

function pre_tests { # run this before test are done with "$0" as argument
  filename="$0"
  have_failed=0
  echo
  echo "Test suite starting: $filename" at `date`
  echo
}

function post_tests { # run this after tests are done with "$0" as argument
  filename="$0"
  echo
  echo "Test suite done: $filename" at `date`
  if [ $have_failed == 1 ]
  then
    echo "There was at least one failure"
    exit 1
  else
    echo "All tests passed"
  fi
  echo
}

# Write tests like this: TEST_NAME TEST_CMD EXPECTED_RESULT LINE_REFERENCE
# assert_equals_string "simple test that echoes hej" "echo hej" "hej" $0:$LINENO


