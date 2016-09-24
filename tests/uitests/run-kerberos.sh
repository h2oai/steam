#!/bin/bash

scp .kerberos.sh patrick@mr-0xc7:~/
ssh root@mr-0xc7 /home/patrick/.kerberos.sh
ssh root@mr-0xc7 killall steam


