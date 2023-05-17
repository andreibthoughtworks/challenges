package main

import (
	"fmt"
	. "polaris/internal/application/domain"
	. "polaris/internal/application/service"
	. "polaris/internal/infrastructure/repository"
)

var (
	adRepository    = ProvideInMemoryRepository()
	idGenerator     = ProvideIdGenerator()
	clock           = ProvideClock()
	createAdService = ProvideCreateAdService()
	findAdService   = ProvideFindAdService()
	findAdsService  = ProvideFindAdsService()
)

func main() {

	fmt.Println("Go to save new ad")
	fmt.Println("--------------------------------")

	createdAd := createAd("titulo1", "descripcion 1", 12)
	fmt.Printf("Your new Ad %v was created at %v\n", createdAd.Id, createdAd.CreatedAt)
	fmt.Println("--------------------------------")

	foundAd, adNotFoundError := findAdService.Execute(FindAdRequest{Id: createdAd.Id})
	if adNotFoundError != nil {
		fmt.Printf("Error, Ad %v not found\n", createdAd.Id)
	} else {
		fmt.Printf("Found Ad  %v\n", foundAd)
	}
	fmt.Println("--------------------------------")

	createAd("titulo2", "descripcion 2", 12)
	createAd("titulo3", "descripcion 3", 12)
	createAd("titulo4", "descripcion 4", 12)
	createAd("titulo5", "descripcion 5", 12)
	createAd("titulo6", "descripcion 6", 12)
	createAd("titulo7", "descripcion 7", 12)
	foundAdResponse := findAdsService.Execute()

	fmt.Println("--------------------------------")
	fmt.Printf("Found Ads  %v\n", foundAdResponse.Ads)

}

func createAd(
	title string,
	description string,
	price uint,
) CreateAdResponse {
	request := CreateAdRequest{
		Title:       title,
		Description: description,
		Price:       price,
	}

	createdAd := createAdService.Execute(request)
	return createdAd
}

func ProvideCreateAdService() CreateAdService {
	return CreateAdService{
		AdRepository: adRepository,
		IdGenerator:  idGenerator,
		Clock:        clock,
	}
}

func ProvideFindAdService() FindAdService {
	return FindAdService{
		AdRepository: adRepository,
	}
}

func ProvideFindAdsService() FindAdsService {
	return FindAdsService{
		AdRepository: adRepository,
	}
}

func ProvideInMemoryRepository() AdRepository {
	return &InMemoryAdRepository{}
}

func ProvideIdGenerator() IdGenerator {
	return &UUIDGenerator{}
}

func ProvideClock() Clock {
	return &ClockImpl{}
}