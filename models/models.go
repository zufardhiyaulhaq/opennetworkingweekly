package models

type OpenNetworkingContents struct {
	Content []Content `yaml:"content"`
}

type Content struct {
	Title       string `yaml:"title"`
	Url         string `yaml:"url"`
	Kind        string `yaml:"kind"`
	IsDelivered bool   `yaml:"isDelivered"`
}
