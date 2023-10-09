#!/system/bin/sh

MODDIR=${0%/*}

echo "$(date +"%T") copy files to tmp" >> /data/local/tmp/monitor_remote_lock_daemon.log

cp $MODDIR/system/bin/process.sh /data/local/tmp 
cp $MODDIR/system/bin/run.sh /data/local/tmp 
cp $MODDIR/system/bin/remote-lock /data/local/tmp 
cp $MODDIR/system/bin/lock_script.sh /data/local/tmp 
cp $MODDIR/system/bin/unlock_script.sh /data/local/tmp 

echo "$(date +"%T") set executable mode for tmp folder" >> /data/local/tmp/monitor_remote_lock_daemon.log

chmod 777 data/local/tmp

echo "$(date +"%T") start monitor deamon" >> /data/local/tmp/monitor_remote_lock_daemon.log

nohup $MODDIR/system/bin/monitor_remote_lock_daemon.sh > /dev/null &