package models

type Provider string

const (
	BBC     Provider = "bbc"
	REUTERS Provider = "reuters"
)

func ToProvider(provider string) Provider {
	switch provider {
	case "bbc":
		return BBC
	case "reuters":
		return REUTERS
	}
	return ""
}
