
function assert_equals_string  { # assert_equals_string("echo hej", "hej")
  testname="$1"
  cmd="$2"
  expected="$3"
  reference="$4"
#  echo ">>> cmd $cmd"
#  res=`$cmd`
  res=`bash -c "$cmd"`
#  echo ">>> res $res"
  if [ "$res" != "$expected" ] 
  then
    echo
    echo "Test FAILED: $testname"
    echo "cmd: $cmd"
    echo "expected: $expected"
    echo "found: $res"
    echo "in file $reference"
#    exit 1
    echo
  else 
    echo "Test OK: $testname"
  fi
#  echo ""
}

assert_equals_string "ls good" "ls -1 | head -1" "1" $0:$LINENO
#assert_equals_string "ls bad" "ls -1 | head -1" "xyz" $0:$LINENO
#exit 0

assert_equals_string "hej" "echo hej" "hej" $0:$LINENO
assert_equals_string "hej single quote" 'echo hej' 'hej' $0:$LINENO

assert_equals_string "hej echo -n" "echo -n hej; echo" "hej" $0:$LINENO
assert_equals_string "hej echo -n single quote" 'echo -n hej; echo' "hej" $0:$LINENO

