package api

// SpringInitializer is a struct to hold spring initializer object
type SpringInitializer struct {
	Packagemanager       string `json:"Packagemanager"`
	Language             string `json:"Language"`
	SpringBootVersion    string `json:"SpringBootVersion"`
	Group                string `json:"Group"`
	Artifact             string `json:"Artifact"`
	Name                 string `json:"Name"`
	Description          string `json:"Description"`
	PackageName          string `json:"PackageName"`
	Packaging            string `json:"Packaging"`
	JavaVersion          string `json:"JavaVersion"`
	SelectedDependencies string `json:"SelectedDependencies"`
}
