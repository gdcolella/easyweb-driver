package main

import "os/exec"
import "log"

func restart() {
	exec.Command("/bin/bash", "-c", "sudo shutdown -r -t now").Start()
	log.Print("Restarting.")
}

func take_hdmi() {
	exec.Command("/bin/bash", "-c", "echo \"as\" | cec-client -s -d 1").Start()
	log.Print("Took HDMI")
}

func release_hdmi() {
	exec.Command("/bin/bash", "-c", "echo \"tx 4f:82:20:00\" | cec-client -s -d 1").Start()
}

func open_browser() {
	exec.Command("/usr/bin/chromium-browser", "--kiosk", WEBSITE, "--auto-open-devtools-for-tabs", "--disable-infobars", "--use-fake-ui-for-media-stream").Start()
	log.Print("Opened browser.")
}
