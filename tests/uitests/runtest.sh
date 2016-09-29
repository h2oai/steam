#!/bin/bash

i=0

export WD=`pwd`

cnt=`ls -d *-test | wc -l`
touch /root/.failures

num=$(($RANDOM % 9))
if [ "$#" -ge 1 ] 
then
	num=$1
fi

./steam/steam login localhost:9000 --username superuser --password superuser



for dir in `ls -d *-test`; do
	if [ $i -ne $num ] 
	then
		i=`expr $i + 1`
		continue
	fi
	i=`expr $i + 1`
	cp testutil.py $dir/
	python $dir/run.py > /root/.tmp
	pass=$?
	if [ $pass -ne 0 ] 
	then
		echo $dir >> /root/.failures
		cat /root/.tmp >> /root/.failures
		i=`expr $i + 1`
		echo "$dir FAILED"
	else
		echo "$dir PASSED"
	fi
	rm $dir/testutil.*
	break
done

cat /root/.failures

