#!/usr/bin/env bash
# Filename: run_experiment.sh

# change this line below to be abs path to script
cd /mnt/c/code/52060_parallel/hw3-asclark109/hw3/experiment

### BOILERPLATE: borrowed from 
### https://betterdev.blog/minimal-safe-bash-script-template/
# set -Eeuo pipefail
trap cleanup SIGINT SIGTERM ERR EXIT

display_help() {
    echo
    echo "Usage: $(basename "${BASH_SOURCE[0]}") <minDelay> <maxDelay> <incrementAmount> <maxThreadCnt> <trials>" >&2
    echo "this script runs experiments for various locks"
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

# interpret args
if [ $# -eq 5 ]; then 
  minDelay=$1
  maxDelay=$2
  incrementAmount=$3
  maxThreadCnt=$4
  trials=$5

  if [[ "minDelay" =~ ^[0-9]+$ ]]
      then
          echo
          echo "[minDelay] argument must be positive integer"
          display_help
          return 0
  fi
  echo MIN_DELAY = $minDelay

  if [[ "maxDelay" =~ ^[0-9]+$ ]]
      then
          echo
          echo "[maxDelay] argument must be positive integer"
          display_help
          return 0
  fi
  echo MAX_DELAY = $maxDelay

  if [[ "incrementAmount" =~ ^[0-9]+$ ]]
      then
          echo
          echo "[incrementAmount] argument must be positive integer"
          display_help
          return 0
  fi
  echo INCREMENT_AMOUNT = $incrementAmount

  if [[ "maxThreadCnt" =~ ^[0-9]+$ ]]
      then
          echo
          echo "[maxThreadCnt] argument must be positive integer"
          display_help
          return 0
  fi
  echo MAX_THREAD_COUNT = $maxThreadCnt
  
  if [[ "trials" =~ ^[0-9]+$ ]]
      then
          echo
          echo "[trials] argument must be positive integer"
          display_help
          return 0
  fi
  echo TRIALS = $trials

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

msg "${RED}TAS TESTS:${NOFORMAT}"
for (( i=1; i<=$trials; i++ ))
do
    for (( threads=1; threads<=$maxThreadCnt; threads++ ))
    do
        echo Trial=$i threads=$threads
        NEWFILE="${EXP_DIR}/tas_${i}_threads_$threads"
        { echo EXP_TYPE=TAS; echo TRIAL=$i; echo THREADS=$threads;   } > $NEWFILE
        { echo MIN_DELAY=$minDelay; echo MAX_DELAY=$maxDelay; echo INCREMENT_AMOUNT=$incrementAmount;   } >> $NEWFILE
        { time go run experiment.go tas $incrementAmount $threads; } 2>> $NEWFILE
    done
done

msg "${RED}TTAS TESTS:${NOFORMAT}"
for (( i=1; i<=$trials; i++ ))
do
    for (( threads=1; threads<=$maxThreadCnt; threads++ ))
    do
        echo Trial=$i threads=$threads
        NEWFILE="${EXP_DIR}/ttas_${i}_threads_$threads"
        { echo EXP_TYPE=TTAS; echo TRIAL=$i; echo THREADS=$threads;   } > $NEWFILE
        { echo MIN_DELAY=$minDelay; echo MAX_DELAY=$maxDelay; echo INCREMENT_AMOUNT=$incrementAmount;   } >> $NEWFILE
        { time go run experiment.go ttas $incrementAmount $threads; } 2>> $NEWFILE
    done
done

msg "${RED}EB TESTS:${NOFORMAT}"
for (( i=1; i<=$trials; i++ ))
do
    for (( threads=1; threads<=$maxThreadCnt; threads++ ))
    do
        echo Trial=$i threads=$threads
        NEWFILE="${EXP_DIR}/eb_${i}_threads_$threads"
        { echo EXP_TYPE=EB; echo TRIAL=$i; echo THREADS=$threads;   } > $NEWFILE
        { echo MIN_DELAY=$minDelay; echo MAX_DELAY=$maxDelay; echo INCREMENT_AMOUNT=$incrementAmount;   } >> $NEWFILE
        { time go run experiment.go eb $minDelay $maxDelay $incrementAmount $threads; } 2>> $NEWFILE
    done
done

msg "${RED}EBF TESTS:${NOFORMAT}"
for (( i=1; i<=$trials; i++ ))
do
    for (( threads=1; threads<=$maxThreadCnt; threads++ ))
    do
        echo Trial=$i threads=$threads
        NEWFILE="${EXP_DIR}/ebf_min${minDelay}_max${maxDelay}_${i}_threads_$threads"
        { echo EXP_TYPE=EBF; echo TRIAL=$i; echo THREADS=$threads;   } > $NEWFILE
        { echo MIN_DELAY=$minDelay; echo MAX_DELAY=$maxDelay; echo INCREMENT_AMOUNT=$incrementAmount;   } >> $NEWFILE
        { time go run experiment.go ebf $minDelay $maxDelay $incrementAmount $threads; } 2>> $NEWFILE
    done
done

msg "${RED}SYNC TESTS:${NOFORMAT}"
for (( i=1; i<=$trials; i++ ))
do
    for (( threads=1; threads<=$maxThreadCnt; threads++ ))
    do
        echo Trial=$i threads=$threads
        NEWFILE="${EXP_DIR}/sync_${i}_threads_$threads"
        { echo EXP_TYPE=SYNC; echo TRIAL=$i; echo THREADS=$threads;   } > $NEWFILE
        { echo MIN_DELAY=$minDelay; echo MAX_DELAY=$maxDelay; echo INCREMENT_AMOUNT=$incrementAmount;   } >> $NEWFILE
        { time go run experiment.go sync $incrementAmount $threads; } 2>> $NEWFILE
    done
done

############ END OF EXPERIMENTS ############
