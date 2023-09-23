#!/system/bin/sh

# config/status variables
TERMINAL=false
LOGFILE="/data/local/tmp/monitor_remote_lock_daemon.log"
INTERVAL=60
MAX_RETRIES=5
SESSION_TIME=3072
MAGISKD_PID=""
APP="remote-lock"
APP_PATH="./remote-lock"

# count variables
DAYS_PASSED=0
RUNNING=0
CRASHED=0
RECOVERIES=0

function destroy_other_instances() {
	local PGREP_SCRIPT_STDOUT=$((pgrep -l $APP) 2>&1)
	PGREP_SCRIPT_STDOUT="${PGREP_SCRIPT_STDOUT//$'\n'/$'\n\t'}"
	echo "Processes:\n\t$PGREP_SCRIPT_STDOUT"
	
	PGREP_SCRIPT_STDOUT=$((pgrep $APP) 2>&1)
	
	local CURRENT_SCRIPT_PID=$$
	local CURRENT_SCRIPT_NAME=$(basename -- "$0")
	echo "Current Process: $CURRENT_SCRIPT_PID $CURRENT_SCRIPT_NAME"
	
	PGREP_SCRIPT_STDOUT="${PGREP_SCRIPT_STDOUT//$'\n'/' '}"
	PGREP_SCRIPT_STDOUT="${PGREP_SCRIPT_STDOUT//$CURRENT_SCRIPT_PID/''}"
	if [ "$PGREP_SCRIPT_STDOUT" ]
	then
		echo "Stripped: $PGREP_SCRIPT_STDOUT"
		for PROCESS_ID in "$PGREP_SCRIPT_STDOUT"
		do
			local PROCESS_KILL_RESULT=$((kill -KILL $PROCESS_ID) 2>&1)
			echo "Script instance killed : $PROCESS_ID > $PROCESS_KILL_RESULT"
		done
	else
		echo "No extra instance found"
	fi
}

function init() {
	while [ "$1" ]; do
	    case $1 in
	        -s  | --single )        destroy_other_instances
	                                ;;
	        -nl | --no-log )        LOGFILE=""
	                                ;;
	        -t  | --term   )		TERMINAL=true
	       						 ;;
	                  *    )		echo "Unknown parameter"
	                 			   exit 1
	                 			   ;;
	    esac
	    shift
	done
	if [ $LOGFILE ] && $TERMINAL; then log "LOGFILE: $LOGFILE"; else log "LOGFILE OFF"; fi
	log "__init__"
}

function log() {
	local LOGSTRING="$1 | $(date)"
	if $TERMINAL; then echo $LOGSTRING; fi
	if [ $LOGFILE ]; then echo $LOGSTRING >> $LOGFILE; fi
}

function start_logfile() {
	if [ $LOGFILE ]; then touch $LOGFILE && echo "" > $LOGFILE; fi
}

function reset() {
	((DAYS_PASSED++))
	RUNNING=0
	CRASHED=0
	RECOVERIES=0
	
	start_logfile
	log "DAY: $DAYS_PASSED" 
}

function check_remote_lock_running() {
	#pgrep magiskd > /dev/null 2>&1 # filename should not contain magiskd
	MAGISKD_PID=$((pgrep $APP) 2>&1)
}

function monitor_magiskd() {
	while true; do
		check_remote_lock_running
		if [ "$MAGISKD_PID" ]; then
			((RUNNING++))
			log "remote-lock daemon running : $RUNNING"
			log "remote-lock process id : $MAGISKD_PID"
		else
			((CRASHED++))
			log "remote-lock daemon crashed : $CRASHED"
			log "Attempting to restart : $CRASHED"
			
			local FAILURES=0
			
		    while true; do
			    local MAGISK_STDOUT=$((sh /data/local/tmp/run.sh &) 2>&1)
			    log "$MAGISK_STDOUT"
			    check_remote_lock_running
			    if [[ "$MAGISK_STDOUT" = *launching*new*process* ]] || [ "$MAGISKD_PID" ]
			    then
			    	((RECOVERIES++))
			    	log "Successfully recovered remote-lock : $RECOVERIES"
			    	break
			    fi  
			    	    
			    ((FAILURES++))
			    log "Failed to recover remote-lock : $FAILURES"
			    if [ $FAILURES -ge $MAX_RETRIES ]
			    then
			   	 log 'Max retries reached. Exiting'
			   	 exit
			   	 break
			    fi
			    sleep 30
		    done
		fi
		
		if [ $(expr "$RUNNING" + "$CRASHED") -ge "$SESSION_TIME" ]; then reset; fi
		
		sleep $INTERVAL
	done
}

init "$@"
start_logfile
monitor_magiskd

