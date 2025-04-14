package usecase

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type EnrichedData struct {
	Age         int
	Gender      string
	Nationality string
}

func EnrichName(name string) (EnrichedData, error) {
	var result EnrichedData

	ageResp, _ := http.Get(fmt.Sprintf("https://api.agify.io/?name=%s", name))
	genderResp, _ := http.Get(fmt.Sprintf("https://api.genderize.io/?name=%s", name))
	natResp, _ := http.Get(fmt.Sprintf("https://api.nationalize.io/?name=%s", name))

	defer ageResp.Body.Close()
	defer genderResp.Body.Close()
	defer natResp.Body.Close()

	var a struct{ Age int }
	var g struct{ Gender string }
	var n struct{ Country []struct{ CountryID string } }

	json.NewDecoder(ageResp.Body).Decode(&a)
	json.NewDecoder(genderResp.Body).Decode(&g)
	json.NewDecoder(natResp.Body).Decode(&n)

	result.Age = a.Age
	result.Gender = g.Gender
	if len(n.Country) > 0 {
		result.Nationality = n.Country[0].CountryID
	}

	return result, nil
}
