#!/bin/sh
echo $0
dataFilePath=$1
dataFile=$2
qi=$3
sa=$4
centerCode=$5
tableId=$6
rScript=$7
rLog=$8

args="\"$dataFilePath\" \"$dataFile\" $qi $sa \"$centerCode\" \"$tableId\""
echo $args

R CMD BATCH --vanilla "--args \"$dataFilePath\" \"$dataFile\" $qi $sa \"$centerCode\" \"$tableId\"" "$rScript" "$rLog"

runR() {
  #R CMD BATCH --vanilla "--args \"$1\" \"F_BBP14_00006\" c(\"qi1\",\"qi2\") c(\"sa1\",\"sa2\") \"BBP14\" \"TBBP14_ID_06\"" /home/fasoo/R/script/r_script_fasoo.R.bak /home/fasoo/R/log/F_BBP14_00006.out
  R CMD BATCH --vanilla "--args $1" "$2" "$3"
}

#runR "$args" "$rScript" "$rLog"