#!/usr/bin/env bash
WD=`dirname $0`
. ${WD}/include.sh
PK3="${WD}/CameramanEditor.pk3"
LOG="${WD}/editor.log"
ARGS="+freelook 1 +noclip 2 +notarget -file $PK3 +logfile $LOG "

case $1 in
  load)
    ARGS+="$(loadcam $2) +pukename Cman_WarpToPath"
    shift 2
    ;;
esac

$gzdoom $@ $ARGS && expo $LOG
