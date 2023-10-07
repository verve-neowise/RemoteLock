#!/system/bin/sh

MODDIR=${0%/*}

cp $MODDIR/system/bin/process.sh /data/local/tmp 
cp $MODDIR/system/bin/run.sh /data/local/tmp 
cp $MODDIR/system/bin/remote-lock /data/local/tmp 
cp $MODDIR/system/bin/lock_script.sh /data/local/tmp 
cp $MODDIR/system/bin/unlock_script.sh /data/local/tmp 

chmod 777 data/local/tmp

nohup $MODDIR/system/bin/monitor_remote_lock_daemon.sh > /dev/null &