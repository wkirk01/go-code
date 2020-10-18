package lifx

import (
	"encoding/json"
)

//Light ...
type Light struct {
	ID                   string       `json:"id"`
	UUID                 string       `json:"uuid"`
	Label                string       `json:"label"`
	Connected            bool         `json:"connected,omitempty"`
	Power                string       `json:"power"`
	Color                Color        `json:"color"`
	Brightness           float32      `json:"brightness"`
	Group                group        `json:"group"`
	Location             location     `json:"location"`
	Capabilities         capabilities `json:"capabilities"`
	LastSeen             string       `json:"last_seen"`
	SecondsSinceLastSeen int          `json:"seconds_since_seen"`
}

//Color ...
type Color struct {
	Hue        int `json:"hue"`
	Saturation int `json:"saturation"`
	Kelvin     int `json:"kelvin"`
}

type group struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type location struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type product struct {
	Name       string `json:"name"`
	Identifier string `json:"identifier"`
	Company    string `json:"company"`
	VendorID   string `json:"vendor_id"`
	PendorID   string `json:"product_id"`
}

type capabilities struct {
	HasColor             bool `json:"has_color"`
	HasVariableColorTemp bool `json:"has_variable_color_temp"`
	HasIR                bool `json:"has_ir"`
	HasChain             bool `json:"has_chain"`
	HasMatrix            bool `json:"has_matrix"`
	HasMultiZone         bool `json:"has_multizone"`
	MinKelvin            int  `json:"min_kelvin"`
	MaxKelvin            int  `json:"max_kelvin"`
}

//GetLights ...
func GetLights() []Light {
	respBody := MakeRequest("GET", "lights/all", nil)
	lights := []Light{}
	err := json.Unmarshal(respBody, &lights)
	check(err)
	return lights
}
