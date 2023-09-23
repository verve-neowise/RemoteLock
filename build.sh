cd service
sh ./bundle.sh linux 
cd ../
cp service/remote-lock module/system/bin

cd module

zip -vr ../lock-module.zip ./* -x "*.DS_Store"

cd ..

if [ -z ${1} ] 
then
    echo "without push to device"
else
    adb push lock-module.zip /sdcard
fi