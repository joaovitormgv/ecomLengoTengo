package models

type User struct {
	ID                 int    `json:"id"`
	Email              string `json:"email"`
	Username           string `json:"username"`
	Password           string `json:"password"`
	Phone              string `json:"phone"`
	Gender             string `json:"gender"`
	AddressGeolocation struct {
		Lat  float64 `json:"lat"`
		Long float64 `json:"long"`
	} `json:"address_geolocation"`
	AddressCity    string `json:"address_city"`
	AddressStreet  string `json:"address_street"`
	AddressNumber  string `json:"address_number"`
	AddressZipcode string `json:"address_zipcode"`
	NameFirstname  string `json:"name_firstname"`
	NameLastname   string `json:"name_lastname"`
}
