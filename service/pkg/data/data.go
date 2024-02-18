package data

type HealthStatus struct {
	Timestamp string `json:"timestamp"`
	Status    string `json:"status"`
}

type AboutInfo struct {
	Name       string              `json:"name"`
	Author     string              `json:"author"`
	Repository string              `json:"repository"`
	Version    string              `json:"version"`
	License    string              `json:"license"`
	Languages  []string            `json:"languages"`
	Algorithms map[string][]string `json:"algorithms"`
}

type BasicErrorInfo struct {
	StatusCode int
	Endpoint   string
	Message    string
}
