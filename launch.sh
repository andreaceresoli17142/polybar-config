#!/usr/bin/env bash

cd $HOME/.config/polybar

isLocked=`cat .lock`

# echo $isLocked >> .deb

if [ $isLocked == "0" ];
then
	echo "1" > .lock

	# Terminate already running bar instances
	killall -q polybar

	# Wait until the processes have been shut down
	while pgrep -x polybar >/dev/null; do sleep 1; done

	for monitor in $(xrandr --query | grep " connected" | cut -d" " -f1); do
   	echo "Starting top bar on monitor '$monitor'"
   	MONITOR=$monitor polybar stdBar &
	done

	echo "0" > .lock

	echo "Bars launched..."

fi
