#!/bin/bash
tmp=`sensors | awk '/^Package id 0:/ { print substr ( $4, 2 ) }' | egrep -o '[0-9]+.[0-9]+'`
readarray -d . -t splitted<<< "$tmp"
echo "" $splitted°C
