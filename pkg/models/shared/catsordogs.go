// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// CatsOrDogs - Case 2
type CatsOrDogs struct {
	Value []Cat
}

func (o *CatsOrDogs) GetValue() []Cat {
	if o == nil {
		return []Cat{}
	}
	return o.Value
}
