package views

import (
	"log"

	"chefhub.pw/models"
)

const (
	// AlertLvlError is the Bootstrap danger formatting.
	AlertLvlError = "danger"

	// AlertLvlWarning is the Bootstrap warning formatting.
	AlertLvlWarning = "warning"

	// AlertLvlInfo is the Bootstrap warning formatting.
	AlertLvlInfo = "info"

	// AlertLvlSuccess is the Bootstrap success formatting.
	AlertLvlSuccess = "success"

	// AlertMsgGeneric is displayed when any random error
	// is encountered by our backend.
	AlertMsgGeneric = "Something went wrong. Please try again, and contact us if the problem persists."
)

// Alert is used to render Bootstrap Alert messages in templates.
type Alert struct {
	Level   string
	Message string
}

// Data is the top level structure that views expect data
// to come in.
type Data struct {
	Alert *Alert
	User  *models.User
	Yield interface{}
}

// SetAlert only lets users see errors that are
// approved for them to receive.
func (d *Data) SetAlert(err error) {
	if pErr, ok := err.(PublicError); ok {
		d.Alert = &Alert{
			Level:   AlertLvlError,
			Message: pErr.Public(),
		}
	} else {
		log.Println(err)
		d.Alert = &Alert{
			Level:   AlertLvlError,
			Message: AlertMsgGeneric,
		}
	}
}

// AlertError is our helper function for showing errors.
func (d *Data) AlertError(msg string) {
	d.Alert = &Alert{
		Level:   AlertLvlError,
		Message: msg,
	}
}

// PublicError is our interface.
type PublicError interface {
	error
	Public() string
}
