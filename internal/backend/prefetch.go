package backend

import (
	"os"
)

/*
	Deletes all files in `Prefetch` in %windir%
	It may still log the name of this exe in the prefetch folder
	I have tried disabling prefetch via registry but it still logs
*/

// Notification represents a user notification that can be sent to the operating system.
type Notification struct {
	Title, Content string
}

// NewNotification creates a notification that can be passed to App.SendNotification.
func NewNotification(title, content string) *Notification {
	return &Notification{Title: title, Content: content}
}

func Prefetch() error {
	prefetch := "C:\\Windows\\Prefetch"
	os.RemoveAll(prefetch)
	os.Mkdir(prefetch, 0755)

	return nil

	//fyne.CurrentApp().SendNotification(&fyne.Notification{
	//	Title:   "HeadshotHider",
	//	Content: "Successfully cleared prefetch",
	//})
}