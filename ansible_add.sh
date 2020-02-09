#!/bin/bash
if [ -z "$1" ]
  then
    logger "No argument supplied"
    exit 1
fi
logger "Added device $1"

if [ ! -f "/tmp/device" ] ; then
  logger "No devicefile exists, creating..."
  device=$1
  echo "${device}" > /tmp/device
fi

FILE=/tmp/device
OLDTIME=120
CURTIME=$(date +%s)
FILETIME=$(stat $FILE -c %Y)
TIMEDIFF=$(expr $CURTIME - $FILETIME)

logger "Last devicefile change: $TIMEDIFF seconds ago"

if [ $TIMEDIFF -gt $OLDTIME ] || [ $TIMEDIFF -eq 0 ]; then
    logger "Updating active device..."
    echo "${device}" > /tmp/device
    logger "Running pucker playbook for $1"
    /bin/su - pi -c "ssh pi@comms.wotlemons.com \"ansible-playbook /home/pi/playbooks/set_up_pucker.yml --extra-vars \"device=$1\"\""
  else
    logger "Device added too recently, try again later."
    exit 0
fi
