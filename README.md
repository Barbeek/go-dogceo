# go-dogceo
Go wrapper for Dog CEO API

This library provides a connection to [Dog CEO API](https://dog.ceo/dog-api/). 
It will initialize an interface with different methods to interact with the API.

Note: This product uses the Dog CEO API but is not endorsed or certified by it.
## How to install

```
go get https://github.com/Barbeek/go-dogceo
```

## Init the interface
- `New() DogCEO` : set the api URL and returns a DogCEO interface

## DogCEO Methods
The methods provided for the interface are :

- `ListBreeds() (*MessageMap, error)`: returns all dog's breeds and their subbreeds
- `RandomImage() (*Message, error)`: returns a single random image from all dog collection
- `RandomImages(numberOfImages string) (*MessageArray, error)`: returns, randomly, the number of images passed in parameter from all dog collection
- `ImagesByBreed(breed string) (*MessageArray, error)`: returns an array of all the images from a breed
- `RandomImageByBreed(breed string) (*Message, error)`: returns a random dog image from a breed
- `RandomImagesByBreed(breed, numberOfImages string) (*MessageArray, error)`: returns, randomly, the number of images passed in parameter from a breed
- `ListSubBreeds(breed string) (*MessageArray, error)`: returns an array of all the sub-breeds from a breed
- `ImagesBySubBreed(breed, subbreed string) (*MessageArray, error)`: returns an array of all the images from the sub-breed
- `RandomImageBySubBreed(breed, subbreed string) (*Message, error)`: returns a random dog image from a breed's sub-bread.
- `RandomImagesBySubBreed(breed, subbreed, numberOfImages string) (*MessageArray, error)`: returns, randomly, the number of images passed in parameter from a breed's sub-breed

## How to use

Import the library

```
import "https://github.com/Barbeek/go-dogceo"
```

Create an instance of DogCEO interface

```
dogCEO, err := dogceo.New()
```

Call the API methods as you want. 

*Exemple, to get images url by bread*

```
breeds, err := dogCEO.ListBreeds()
if err != nil {
    return err
}

for key, _ := range breeds.Message {
    images, err := dogCEO.ImagesByBreed(key)
    if err != nil {
       return err
    }
}
```

## How to test
You can run unit tests by executing command `make test` and linter with `make lint`

*Enjoy!*