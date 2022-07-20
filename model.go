package dogceo

type api struct {
	baseURL string
}

type Message struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

type MessageArray struct {
	Message []string `json:"message"`
	Status  string   `json:"status"`
}

type MessageMap struct {
	Message map[string][]string `json:"message"`
	Status  string              `json:"status"`
}

type DogCEO interface {
	ListBreads() (*MessageMap, error)
	ImagesByBread(string) (*MessageArray, error)
	RandomImage() (*Message, error)
	RandomImages(string) (*MessageArray, error)
	RandomImageByBread(string) (*Message, error)
	RandomImagesByBread(string, string) (*MessageArray, error)
	ListSubBreads(string) (*MessageArray, error)
	ImagesBySubBreed(string, string) (*MessageArray, error)
	RandomImageBySubBreed(string, string) (*Message, error)
}
