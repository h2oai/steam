#!/bin/bash

scp .yarn.sh patrick@mr-0xd7:~/
ssh root@mr-0xd7 /home/patrick/.yarn.sh
ssh root@mr-0xd7 killall steam

