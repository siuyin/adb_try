package main

import (
	"fmt"
	"log"

	"github.com/electricbubble/gadb"
)

func main() {
	adbClient, err := gadb.NewClient()
	checkErr(err, "fail to connect adb server")

	devices, err := adbClient.DeviceList()
	checkErr(err)

	if len(devices) == 0 {
		log.Fatalln("list of devices is empty")
	}

	dev := devices[0]

	shellOutput, err := dev.RunShellCommand("pm list packages")
	checkErr(err, "pm install")
	fmt.Println(shellOutput)

}

func checkErr(err error, msg ...string) {
	if err == nil {
		return
	}

	var output string
	if len(msg) != 0 {
		output = msg[0] + " "
	}
	output += err.Error()
	log.Fatalln(output)
}
