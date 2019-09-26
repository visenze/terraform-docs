package settings

// Settings represents all settings
type Settings struct {
	// AggregateTypeDefaults print default values of aggregate types (default: false)
	// scope: Global
	AggregateTypeDefaults bool

	// EscapeMarkdown escapes special Markdown characters (such as | _ * and etc) (default: true)
	// scope: Markdown
	EscapeMarkdown bool

	// MarkdownIndent control the indentation of Markdown headers [available: 1, 2, 3, 4, 5] (default: 2)
	// scope: Markdown
	MarkdownIndent int

	// ShowRequired show "Required" column when generating Markdown (default: true)
	// scope: Markdown
	ShowRequired bool

	// SortByName sorted rendering of inputs and outputs (default: true)
	// scope: Global
	SortByName bool

	// SortInputsByRequired sort inputs by name and prints required inputs first (default: false)
	// scope: Global
	SortInputsByRequired bool
}

//NewSettings returns new instance of Settings
func NewSettings() *Settings {
	return &Settings{
		AggregateTypeDefaults: false,
		EscapeMarkdown:        true,
		MarkdownIndent:        2,
		ShowRequired:          true,
		SortByName:            true,
		SortInputsByRequired:  false,
	}
}
