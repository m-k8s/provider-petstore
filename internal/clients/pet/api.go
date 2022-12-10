package pet

import (
	"encoding/json"
	"fmt"

	petstore "github.com/alexisries/provider-petstore/internal/clients"
)

const (
	PetStatusPending    PetStatus = "PENDING"
	PetStatusAvailable  PetStatus = "AVAILABLE"
	PetStatusInProgress PetStatus = "INPROGRESS"
	PetStatusInactive   PetStatus = "INACTIVE"
	PetStatusFailed     PetStatus = "FAILED"
)

type Category struct {
	Id   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type PetStatus string

type Tag struct {
	Id   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type Pet struct {
	Category  *Category `json:"category,omitempty"`
	Id        *int64    `json:"id,omitempty"`
	Name      string    `json:"name"`
	PhotoUrls []string  `json:"photoUrls"`
	Status    PetStatus `json:"status,omitempty"`
	Tags      *[]Tag    `json:"tags,omitempty"`
}

type PetClient struct {
	*petstore.Client
}

func New(cfg *petstore.Config) PetClient {
	return PetClient{
		petstore.New(cfg),
	}
}

func (c *PetClient) AddPet(pet *Pet) error {
	path := "/pet"
	pet.Status = PetStatusPending
	body, err := json.Marshal(*pet)
	if err != nil {
		return err
	}
	_, err = c.DoRequest(path, "POST", body)
	if err != nil {
		return err
	}
	return nil
}

func (c *PetClient) GetPetById(petId int64) (*Pet, error) {
	path := fmt.Sprintf("/pet/%d", petId)
	res, err := c.DoRequest(path, "GET", nil)
	if err != nil {
		return nil, err
	}
	var pet Pet
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(&pet)
	if err != nil {
		return &Pet{}, err
	}
	return &pet, nil
}

func (c *PetClient) UpdatePetById(petId int64, pet *Pet) error {
	path := fmt.Sprintf("/pet/%d", petId)

	body, err := json.Marshal(*pet)
	if err != nil {
		return err
	}
	_, err = c.DoRequest(path, "PUT", body)
	if err != nil {
		return err
	}
	return nil
}

func (c *PetClient) DeletePetById(petId int64) error {
	path := fmt.Sprintf("/pet/%d", petId)
	_, err := c.DoRequest(path, "DELETE", nil)
	if err != nil {
		return err
	}
	return nil
}
