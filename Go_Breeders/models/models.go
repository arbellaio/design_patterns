package models

import (
	"database/sql"
	"time"
)

var db *sql.DB

type Models struct {
	DogBreed DogBreed
}

func New(conn *sql.DB) *Models {
	db = conn
	return &Models{
		DogBreed: DogBreed{},
	}
}

type DogBreed struct {
	ID               int    `json:"id"`
	Breed            string `json:"breed"`
	WeightLowLbs     int    `json:"weight_low_lbs"`
	WeightHighLbs    int    `json:"weight_high_lbs"`
	AverageWeight    int    `json:"average_weight"`
	Lifespan         int    `json:"average_lifespan"`
	Details          string `json:"details"`
	AlternateNames   string `json:"alternate_names"`
	GeographicOrigin string `json:"geographic_origin"`
}

func (d *DogBreed) All() ([]*DogBreed, error) {
	return d.AllDogBreeds()
}

type CatBreed struct {
	ID               int    `json:"id,omitempty"`
	Breed            string `json:"breed,omitempty"`
	WeightLowLbs     int    `json:"weight_low_lbs,omitempty"`
	WeightHighLbs    int    `json:"weight_high_lbs,omitempty"`
	AverageWeight    int    `json:"average_weight,omitempty"`
	Lifespan         int    `json:"lifespan,omitempty"`
	Details          string `json:"details,omitempty"`
	AlternateNames   string `json:"alternate_names,omitempty"`
	GeographicOrigin string `json:"geographic_origin,omitempty"`
}

type Cat struct {
	ID               int       `json:"id,omitempty"`
	CatName          string    `json:"cat_name,omitempty"`
	BreedID          int       `json:"breed_id,omitempty"`
	BreederID        int       `json:"breeder_id,omitempty"`
	Color            string    `json:"color,omitempty"`
	DateOfBirth      time.Time `json:"date_of_birth,omitempty"`
	SpayedOrNeutered int       `json:"spayed_neutered,omitempty"`
	Description      string    `json:"description"`
	Weight           int       `json:"weight"`
	Breed            CatBreed  `json:"breed"`
	Breeder          Breeder   `json:"breeder"`
}

type Dog struct {
	ID               int       `json:"id,omitempty"`
	DogName          string    `json:"dog_name,omitempty"`
	BreedID          int       `json:"breed_id,omitempty"`
	BreederID        int       `json:"breeder_id,omitempty"`
	Color            string    `json:"color,omitempty"`
	DateOfBirth      time.Time `json:"date_of_birth,omitempty"`
	SpayedOrNeutered int       `json:"spayed_neutered,omitempty"`
	Description      string    `json:"description"`
	Weight           int       `json:"weight"`
	Breed            DogBreed  `json:"breed"`
	Breeder          Breeder   `json:"breeder"`
}

type Breeder struct {
	ID          int         `json:"id,omitempty"`
	BreederName string      `json:"breeder_name,omitempty"`
	Address     string      `json:"address,omitempty"`
	City        string      `json:"city,omitempty"`
	ProvState   string      `json:"prov_state,omitempty"`
	Country     string      `json:"country,omitempty"`
	Zip         string      `json:"zip,omitempty"`
	Phone       string      `json:"phone,omitempty"`
	Email       string      `json:"email,omitempty"`
	Active      int         `json:"active,omitempty"`
	DogBreeds   []*DogBreed `json:"dog_breeds,omitempty"`
	CatBreeds   []*CatBreed `json:"cat_breeds,omitempty"`
}

type Pet struct {
	Species     string `json:"species,omitempty"`
	Breed       string `json:"breed,omitempty"`
	MinWeight   int    `json:"min_weight,omitempty"`
	MaxWeight   int    `json:"max_weight,omitempty"`
	Description string `json:"description,omitempty"`
	Lifespan    int    `json:"lifespan,omitempty"`
}
