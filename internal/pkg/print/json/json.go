package json

import (
	"encoding/json"

	"github.com/segmentio/terraform-docs/internal/pkg/doc"
	"github.com/segmentio/terraform-docs/internal/pkg/settings"
)

const (
	indent string = "  "
	prefix string = ""
)

// Print prints a document as json.
func Print(document *doc.Doc, settings *settings.Settings) (string, error) {
	if document.HasInputs() {
		if settings.SortByName {
			if settings.SortInputsByRequired {
				doc.SortInputsByRequired(document.Inputs)
			} else {
				doc.SortInputsByName(document.Inputs)
			}
		}
	}

	if document.HasOutputs() {
		if settings.SortByName {
			doc.SortOutputsByName(document.Outputs)
		}
	}

	buffer, err := json.MarshalIndent(document, prefix, indent)
	if err != nil {
		return "", err
	}

	return string(buffer), nil
}
