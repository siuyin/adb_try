#!/bin/bash
if adb shell dumpsys display | grep -q "mScreenState=OFF"; then
    echo "Waking up device..."
    adb shell input keyevent 26 && adb shell input swipe 500 1800 500 500 500 && adb shell input text $PASSWD
    echo "Screen on"
else
   echo "Screen is already on"
   exit 0
fi


