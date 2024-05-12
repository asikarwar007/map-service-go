package models

type PlaceRequestQuery struct {
	Location string `json:"location"`
	Category string `json:"category"`
	Language string `json:"lang"`
}

type Place struct {
	Title        string        `json:"title"`
	ID           string        `json:"id"`
	Language     string        `json:"language"`
	OntologyID   string        `json:"ontologyId"`
	ResultType   string        `json:"resultType"`
	Address      Address       `json:"address"`
	Position     Position      `json:"position"`
	Access       []Access      `json:"access"`
	Distance     int           `json:"distance"`
	Categories   []Category    `json:"categories"`
	Chains       []Chain       `json:"chains"`
	References   []Reference   `json:"references"`
	Contacts     []Contact     `json:"contacts"`
	OpeningHours []OpeningHour `json:"openingHours"`
}

type Address struct {
	Label       string `json:"label"`
	CountryCode string `json:"countryCode"`
	CountryName string `json:"countryName"`
	StateCode   string `json:"stateCode"`
	State       string `json:"state"`
	County      string `json:"county"`
	City        string `json:"city"`
	District    string `json:"district"`
	Subdistrict string `json:"subdistrict"`
	Street      string `json:"street"`
	PostalCode  string `json:"postalCode"`
}

type Position struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type Access struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type Category struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Primary bool   `json:"primary"`
}

type Chain struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Reference struct {
	Supplier Supplier `json:"supplier"`
	ID       string   `json:"id"`
}

type Supplier struct {
	ID string `json:"id"`
}

type Contact struct {
	Phone    []Phone `json:"phone"`
	TollFree []Phone `json:"tollFree"`
	WWW      []Phone `json:"www"`
}

type Phone struct {
	Value string `json:"value"`
}

type OpeningHour struct {
	Text       []string     `json:"text"`
	IsOpen     bool         `json:"isOpen"`
	Structured []Structured `json:"structured"`
}

type Structured struct {
	Start      string `json:"start"`
	Duration   string `json:"duration"`
	Recurrence string `json:"recurrence"`
}

type PlaceResponse struct {
	Items []Place `json:"items"`
}
