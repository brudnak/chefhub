package views

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
	Yield interface{}
}
