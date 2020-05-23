#!/bin/sh

mkdir $CERTSDIR
/app/generate.sh $DNS1 $DNS2 $CERTSDIR

/app/planbeer
