#!/usr/bin/env bash

echo "dog unknown dog"

curl -X POST \
--form binary_C1=@dog.jpg \
--form abc=def \
--form data='{Dest: SFO, Orig: JFK}' \
http://localhost:55001/predictbinary

echo "test1 bird"

curl -X POST \
--form binary_C1=@test1.jpg \
http://localhost:55001/predictbinary

echo "test2 dog"

curl -X POST \
--form binary_C1=@test2.jpg \
http://localhost:55001/predictbinary

echo "url of dog"

#curl -X POST \
#--form C1="https://c1.staticflickr.com/1/225/515776742_bce2e6dbea_d.jpg" \
#http://localhost:55001/predictbinary

