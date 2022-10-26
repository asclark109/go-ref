#!/usr/bin/env bash
# Filename: run_experiment.sh

# change this line below to be abs path to script
cd /mnt/c/code/52060_parallel/hw3-asclark109/hw3/covid

### BOILERPLATE: borrowed from 
### https://betterdev.blog/minimal-safe-bash-script-template/
# set -Eeuo pipefail
trap cleanup SIGINT SIGTERM ERR EXIT

display_help() {
    echo
    echo "Usage: $(basename "${BASH_SOURCE[0]}") trials zipcode month year" >&2
    echo "this script runs experiments for the covid data wrangling problem"
    return 0
}

usage() {
  cat<<EOF #
$usage
EOF
  return
}

cleanup() {
  trap - SIGINT SIGTERM ERR EXIT
  # script cleanup here
}

setup_colors() {
  if [[ -t 2 ]] && [[ -z "${NO_COLOR-}" ]] && [[ "${TERM-}" != "dumb" ]]; then
    NOFORMAT='\033[0m' RED='\033[0;31m' GREEN='\033[0;32m' ORANGE='\033[0;33m' BLUE='\033[0;34m' PURPLE='\033[0;35m' CYAN='\033[0;36m' YELLOW='\033[1;33m'
  else
    NOFORMAT='' RED='' GREEN='' ORANGE='' BLUE='' PURPLE='' CYAN='' YELLOW=''
  fi
}

msg() {
  echo >&2 -e "${1-}"
}

die() {
  local msg=$1
  local code=${2-1} # default exit status 1
  msg "$msg"
  return "$code"
}

setup_colors

### END OF BOILERPLATE

############ START OF EXPERIMENTS ############

SCRIPT_DIR=$(cd "$(dirname "${BASH_SOURCE[0]}")" &>/dev/null && pwd -P)
DATA_DIR="$SCRIPT_DIR/data_exp"
EXP_DIR=$DATA_DIR/exp-$(date +"%F_%T")

echo SCRIPT LOCATION = $SCRIPT_DIR
echo DATA_DIR = $DATA_DIR
echo OUTPUT LOCATION = $EXP_DIR

if [ $# -eq 4 ]
then 
    trials=$1
    zipcode=$2
    month=$3
    year=$4
    if [[ "trials" =~ ^[0-9]+$ ]]
        then
            echo
            echo "[trials] argument must be positive integer"
            display_help
            return 0
    fi
    echo TRIALS = $trials
    echo ZIPCODE = $zipcode
    if [[ "month" =~ ^[0-9]+$ ]]
        then
            echo
            echo "[month] argument must be positive integer"
            display_help
            return 0
    fi
    echo MONTH = $month
    if [[ "year" =~ ^[0-9]+$ ]]
        then
            echo
            echo "[year] argument must be positive integer"
            display_help
            return 0
    fi
    echo YEAR = $year
else
    echo "incorrect number of arguments provided."
    display_help
    return 0
fi


# make directories if they dont exist
mkdir -p $DATA_DIR
mkdir -p $EXP_DIR

echo
echo RUNNING TEST...
echo 

msg "${RED}SEQUENTIAL TESTS:${NOFORMAT}"
for (( i=1; i<=$trials; i++ ))
do
    echo Trial=$i
    NEWFILE="${EXP_DIR}/seq_${i}"
    { echo EXP_TYPE=SEQUENTIAL; echo TRIAL=$i; echo THREADS=1;   } > $NEWFILE
    { echo ZIPCODE=$zipcode; echo MONTH=$month; echo YEAR=$year;   } >> $NEWFILE
    { time go run covid.go $zipcode $month $year; } 2>> $NEWFILE
done

echo ""
msg "${RED}STATIC TESTS:${NOFORMAT}"
for (( i=1; i<=$trials; i++ ))
do
    for threads in {2,4,8,12}
    do
        echo Trial=$i threads=$threads
        NEWFILE="${EXP_DIR}/static_${i}_threads_$threads"
        { echo EXP_TYPE=STATIC; echo TRIAL=$i; echo THREADS=$threads;   } > $NEWFILE
        { echo ZIPCODE=$zipcode; echo MONTH=$month; echo YEAR=$year;   } >> $NEWFILE
        { time go run covid.go -mode=static -threads=$threads $zipcode $month $year; } 2>> $NEWFILE
    done
done

echo ""
msg "${RED}MAP TESTS:${NOFORMAT}"
for (( i=1; i<=$trials; i++ ))
do
    echo Trial=$i
    NEWFILE="${EXP_DIR}/map_${i}"
    { echo EXP_TYPE=MAP; echo TRIAL=$i; echo THREADS=500;   } > $NEWFILE
    { echo ZIPCODE=$zipcode; echo MONTH=$month; echo YEAR=$year;   } >> $NEWFILE
    { time go run covid.go -mode=map $zipcode $month $year; } 2>> $NEWFILE
done
############ END OF EXPERIMENTS ############
