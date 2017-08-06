package nessie

type (
	PageMeta struct {
		Previous string `json:"previous"`
		Next     string `json:"next"`
	}
)