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
