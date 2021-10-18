#!/usr/bin/env bash
gzdoom="/home/morb/games/zdoom/gzdoom" #"/path/to/gzdoom"

loadcam() {
awk '{a=a"+cman_"$1." "$3." "}END{print a}' $1
}

expo() {
i=$(startnum);o=0
cat $1 | while read c 
do 
if [[ $c =~ END\ CAM ]]
    then o=0;i=$(printf "%04d" $(($i+1)))
fi
if [ $o -gt 0 ]
    then printf -- "$c\n" >> ${WD}/export-${i}.cman 
fi
if [[ $c =~ BEGIN\ CAM ]]
    then o=1 
fi
done
}

startnum() {
i=$(ls ${WD}/*.cman|sort -n|tail -1|sed "s/${WD}.*-//;s/\..*//" 2>/dev/null)
[ -z $i ] && i=0000 || i=$(printf "%04d" $((${i}+1)))
printf "$i"
}
