package main

import (
	"encoding/json"
	"example/consumo-api/models"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	res, err := http.Get("https://rickandmortyapi.com/api/character/45")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var fullObject models.Character
	json.Unmarshal(body, &fullObject)

	//Extract ID from URL and get full Location Origin object
	idOriginInt := extractIDFromURL(fullObject.OriginBase.URL)
	origin := GetLocation(idOriginInt)
	fullObject.Origin = origin

	//Extract ID from URL and get full Location object
	idLocationInt := extractIDFromURL(fullObject.LocationBase.URL)
	location := GetLocation(idLocationInt)
	fullObject.Location = location

	//Extract ID from URL and get full Episodes objects

	var episodes []models.Episode
	for _, episodeURL := range fullObject.EpisodesBase {
		idEpisodeInt := extractIDFromURL(episodeURL)
		episode := GetEpisodes(idEpisodeInt)
		episodes = append(episodes, episode...)
	}
	fullObject.Episodes = episodes

	printData(fullObject)
}

func extractIDFromURL(url string) int {

	parts := strings.Split(url, "/")
	id := parts[len(parts)-1]
	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
	}

	return idInt
}

func GetEpisodes(idEpisode int) []models.Episode {
	url := fmt.Sprintf("https://rickandmortyapi.com/api/episode/%d", idEpisode)
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var episodes models.Episode
	json.Unmarshal(body, &episodes)

	return []models.Episode{episodes}

}

func GetLocation(idLocation int) models.Location {
	url := fmt.Sprintf("https://rickandmortyapi.com/api/location/%d", idLocation)
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var location models.Location
	json.Unmarshal(body, &location)

	return location

}

func printData(data models.Character) {

	fmt.Println("-----------------------------")
	clean := models.CharacterClean{
		ID:       data.ID,
		Name:     data.Name,
		Status:   data.Status,
		Species:  data.Species,
		Type:     data.Type,
		Gender:   data.Gender,
		Origin:   data.Origin,
		Location: data.Location,
		Image:    data.Image,
		Episodes: data.Episodes,
	}
	jsonData, err := json.MarshalIndent(clean, "", "  ")
	if err != nil {
		fmt.Println("Error al serializar a JSON:", err)
		return
	}
	fmt.Println(string(jsonData))

}
