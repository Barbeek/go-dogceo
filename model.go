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
	ListBreeds() (*MessageMap, error)
	ImagesByBreed(breed string) (*MessageArray, error)
	RandomImage() (*Message, error)
	RandomImages(numberOfImages string) (*MessageArray, error)
	RandomImageByBreed(breed string) (*Message, error)
	RandomImagesByBreed(breed, numberOfImages string) (*MessageArray, error)
	ListSubBreeds(breed string) (*MessageArray, error)
	ImagesBySubBreed(breed, subbreed string) (*MessageArray, error)
	RandomImageBySubBreed(breed, subbreed string) (*Message, error)
	RandomImagesBySubBreed(breed, subbreed, numberOfImages string) (*MessageArray, error)
}
