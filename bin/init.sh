#!/bin/sh

#Set IP instead of DNS
#export PB_DBSERVER=`nslookup $PB_DBSERVER | awk '/^Address: / { print $2 ; exit }'`

mkdir -p $PB_CERTS_PATH
/app/generate.sh $DNS1 $DNS2 $PB_CERTS_PATH

/app/planbeer
