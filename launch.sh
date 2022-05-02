#!/usr/bin/env bash

# Add this script to your wm startup file.

cd $HOME/.config/polybar

# Terminate already running bar instances
# killall -q polybar

# Wait until the processes have been shut down
# while pgrep -u $UID -x polybar >/dev/null; do sleep 1; done

# Launch the bar
# polybar -q main -c config.ini &

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
   	MONITOR=$monitor polybar $1 -c config.ini &
	done

	echo "0" > .lock

	echo "Bars launched..."

fi
