# enable app
pm enable ae.axcapital.lockapp
# set as default launcher 
cmd package set-home-activity "ae.axcapital.lockapp/.MainActivity"
# start app on top
am start -n ae.axcapital.lockapp/.MainActivity --el "ae.axcapital.lockapp.LOCK_ACTION" "lock_action"