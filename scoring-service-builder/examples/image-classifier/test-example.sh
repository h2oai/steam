#!/usr/bin/env bash

curl -X POST \
--form binary_image=@dog.jpg \
--form abc=def \
--form data='{Dest: SFO, Orig: JFK}' \
http://localhost:55001/predictbinary



