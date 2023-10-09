# set as default launcher 
cmd package set-home-activity "ae.axcapital.lockapp/.MainActivity" >> /data/local/tmp/logs.txt
# start app on top
am start -n ae.axcapital.lockapp/.MainActivity >> /data/local/tmp/logs.txt