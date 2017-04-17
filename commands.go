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

func open_browser() {
	exec.Command("/usr/bin/chromium-browser", "--kiosk", WEBSITE, "--incognito", "--disable-infobars").Start()
	log.Print("Opened browser.")
}
