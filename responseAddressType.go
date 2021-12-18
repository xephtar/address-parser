package addressParser

type addressData struct {
	Street string `json:"street"`
	District string `json:"district"`
	SubStreet string `json:"substreet"`
	FullAddress string `json:"full_address"`
}

func (a *addressData) GetStreet() string {
	return a.Street
}

func (a *addressData) GetDistrict() string {
	return a.District
}

func (a *addressData) GetSubStreet() string {
	return a.SubStreet
}

func (a *addressData) GetFullAddress() string {
	return a.FullAddress
}

func (a *addressData) String() string {
	return a.FullAddress
}