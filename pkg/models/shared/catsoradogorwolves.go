// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"errors"
	"github.com/speakeasy-sdks/oneof-v1/pkg/utils"
)

type CatsOrADogOrWolvesValueType string

const (
	CatsOrADogOrWolvesValueTypeArrayOfCat  CatsOrADogOrWolvesValueType = "arrayOfCat"
	CatsOrADogOrWolvesValueTypeDog         CatsOrADogOrWolvesValueType = "Dog"
	CatsOrADogOrWolvesValueTypeArrayOfWolf CatsOrADogOrWolvesValueType = "arrayOfWolf"
)

type CatsOrADogOrWolvesValue struct {
	ArrayOfCat  []Cat
	Dog         *Dog
	ArrayOfWolf []Wolf

	Type CatsOrADogOrWolvesValueType
}

func CreateCatsOrADogOrWolvesValueArrayOfCat(arrayOfCat []Cat) CatsOrADogOrWolvesValue {
	typ := CatsOrADogOrWolvesValueTypeArrayOfCat

	return CatsOrADogOrWolvesValue{
		ArrayOfCat: arrayOfCat,
		Type:       typ,
	}
}

func CreateCatsOrADogOrWolvesValueDog(dog Dog) CatsOrADogOrWolvesValue {
	typ := CatsOrADogOrWolvesValueTypeDog

	return CatsOrADogOrWolvesValue{
		Dog:  &dog,
		Type: typ,
	}
}

func CreateCatsOrADogOrWolvesValueArrayOfWolf(arrayOfWolf []Wolf) CatsOrADogOrWolvesValue {
	typ := CatsOrADogOrWolvesValueTypeArrayOfWolf

	return CatsOrADogOrWolvesValue{
		ArrayOfWolf: arrayOfWolf,
		Type:        typ,
	}
}

func (u *CatsOrADogOrWolvesValue) UnmarshalJSON(data []byte) error {

	dog := Dog{}
	if err := utils.UnmarshalJSON(data, &dog, "", true, true); err == nil {
		u.Dog = &dog
		u.Type = CatsOrADogOrWolvesValueTypeDog
		return nil
	}

	arrayOfCat := []Cat{}
	if err := utils.UnmarshalJSON(data, &arrayOfCat, "", true, true); err == nil {
		u.ArrayOfCat = arrayOfCat
		u.Type = CatsOrADogOrWolvesValueTypeArrayOfCat
		return nil
	}

	arrayOfWolf := []Wolf{}
	if err := utils.UnmarshalJSON(data, &arrayOfWolf, "", true, true); err == nil {
		u.ArrayOfWolf = arrayOfWolf
		u.Type = CatsOrADogOrWolvesValueTypeArrayOfWolf
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u CatsOrADogOrWolvesValue) MarshalJSON() ([]byte, error) {
	if u.ArrayOfCat != nil {
		return utils.MarshalJSON(u.ArrayOfCat, "", true)
	}

	if u.Dog != nil {
		return utils.MarshalJSON(u.Dog, "", true)
	}

	if u.ArrayOfWolf != nil {
		return utils.MarshalJSON(u.ArrayOfWolf, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

// CatsOrADogOrWolves - Case 3
type CatsOrADogOrWolves struct {
	Value CatsOrADogOrWolvesValue
}

func (o *CatsOrADogOrWolves) GetValue() CatsOrADogOrWolvesValue {
	if o == nil {
		return CatsOrADogOrWolvesValue{}
	}
	return o.Value
}
