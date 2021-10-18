#!/usr/bin/env bash
WD=`dirname $0`
. ${WD}/include.sh
PK3="${WD}/CameramanPlayer.pk3"
ARGS="-file $PK3 "

case $1 in
  load)
    ARGS+="$(loadcam $2) +pukename Cman_PlayInPlayer"
    shift 2
    ;;
esac


$gzdoom $@ $ARGS
