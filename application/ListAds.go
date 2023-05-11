package application

import (
	"learning-go-challenges/application/response"
	"learning-go-challenges/domain/ad"
)

type ListAdsRequest struct {
	Limit uint
}

type ListAdsResponse struct {
	Ads []response.AdResponse
}

type ListAdsService struct {
	AdRepository ad.AdRepository
}

func NewListAdsService(adRepository ad.AdRepository) *ListAdsService {
	return &ListAdsService{AdRepository: adRepository}
}

const maxListedAds = 5

func (dependencies ListAdsService) Execute(request ListAdsRequest) ListAdsResponse {
	ads := dependencies.AdRepository.FindAll()

	var responseAds []response.AdResponse
	for _, domainAd := range ads {
		responseAds = append(responseAds, response.FromDomain(domainAd))
	}
	if request.Limit < maxListedAds {
		return ListAdsResponse{Ads: responseAds[:request.Limit]}
	} else {
		return ListAdsResponse{Ads: responseAds[:maxListedAds]}
	}
}
