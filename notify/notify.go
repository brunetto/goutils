package notify

import (
	"log"
	"github.com/godbus/dbus"
)

func notify (title, msg string) () {
	if os.Getenv("DISPLAY") == "" {
		log.Println("Not in graphical session, ignore notification sending and send a beep. \x07")
	}
	conn, err := dbus.SessionBus()
	if err != nil {
		log.Fatal(err)
	}
	obj := conn.Object("org.freedesktop.Notifications", "/org/freedesktop/Notifications")
	call := obj.Call("org.freedesktop.Notifications.Notify", 
					 0, 
					"", 
					uint32(0),
					"", 
					title, 
					msg, 
					[]string{},
					map[string]dbus.Variant{}, 
					int32(5000))
	if call.Err != nil {
		log.Fatal(call.Err)
	}
}

