package main

import (
	"github.com/0xAX/notificator" //for desktop
	"github.com/therecipe/qt/androidextras"
	"github.com/therecipe/qt/core"
	"runtime"
	"fmt"
)

type NotificationHandler struct {
	AndroidNotifier *NotificationClient
	DesktopNotifier *notificator.Notificator
}

type NotificationClient struct {
	core.QObject

	_ func() `constructor:"init"`

	_ string `property:"notification"`

	_ func(string) `slot:"updateAndroidNotification"`
}
func (n *NotificationHandler) Initialise() {

	n.AndroidNotifier = NewNotificationClient(nil)
	n.AndroidNotifier.ConnectNotificationChanged(n.AndroidNotifier.updateAndroidNotification)
	//initialise the notifiers
	n.DesktopNotifier = notificator.New(notificator.Options{
	    DefaultIcon: "",
	    AppName: "WingIt",
	})
}
func (n *NotificationHandler) Push(title, message string) {

	switch runtime.GOOS {
		case "android":
			//in this case, we are notifying on android
		// notificationClient = NewNotificationClient(view)
		// view.Engine().RootContext().SetContextProperty("notificationClient", notificationClient)
		fmt.Println("setting android notification")
		n.AndroidNotifier.SetNotification(message)
		default:
		//notifying on desktop
		n.DesktopNotifier.Push(title, message, "", notificator.UR_CRITICAL)
	}
}

func (c *NotificationClient) init() {
	c.ConnectNotificationChanged(c.updateAndroidNotification)
}

func (c *NotificationClient) updateAndroidNotification(n string) {
	fmt.Println("string is ", n)
	var err = androidextras.QAndroidJniObject_CallStaticMethodVoid2Caught(
		"org/qtproject/example/notification/NotificationClient",
		"notify",
		"(Ljava/lang/String;)V",
		n,
	)

	if err != nil {
		println(err.Error())
	}
}
