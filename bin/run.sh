#!/bin/sh

docker run -v ~/planbeer/test/recipes:/recipes -e DNS1=picobrew.com -e DNS2=www.picobrew.com -e CERTSDIR=/certs -p "8443:443" planbeer
