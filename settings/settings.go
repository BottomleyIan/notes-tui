package settings

import (
	"encoding/json"
	"os"
)

type Language struct {
	Name         string `json:"name"`
	MarkdownName string `json:"markdownName"`
}
type Note struct {
	Title          string
	Description    string
	Key            string
	HasBody        bool
	HasUrl         bool
	HasTitle       bool
	TagFirstLine   bool
	HasCodeSnippet bool
	Fields         []Field
}

type Field struct {
	Name         string
	DefaultValue string
	Type         string
	Options      []string
}

type Settings struct {
	Languages []Language
	Notes     []Note
}

func ParseSettingsFile() Settings {
	settings := Settings{}
	fileBytes, _ := os.ReadFile("settings.json")
	json.Unmarshal(fileBytes, &settings)
	return settings
}

func (s *Settings) LanguageNames() []string {
	if s.Languages == nil {
		return nil
	}
	names := make([]string, len(s.Languages))
	for i, l := range s.Languages {
		names[i] = l.Name
	}
	return names
}
