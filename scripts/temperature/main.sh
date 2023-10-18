#!/bin/bash
tmp=`sensors | awk '/^Package id 0:/ { print substr ( $4, 2 ) }' | grep -E '[0-9]+.[0-9]+'`
readarray -d . -t splitted<<< "$tmp"

if [ "$splitted" -lt "32" ];then
	echo -n "%{F#00cc00}"
elif [ "$splitted" -lt "48" ]; then
	echo -n "%{F#66cc00}"
elif [ "$splitted" -lt "64" ]; then
	echo -n "%{F#cccc00}"
elif [ "$splitted" -lt "80" ]; then
	echo -n "%{F#cc6600}"
else
	 
	if [ "$splitted" -gt "95" ]; then
		notify-send "computer is overheated"
	fi

	echo -n "%{F#cc0000}"
fi

echo " $splitted°C%{F-}"
