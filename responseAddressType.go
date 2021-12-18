package addressParser

type addressData struct {
	Street string `json:"street"`
	District string `json:"district"`
	SubStreet string `json:"substreet"`
	FullAddress string `json:"full_address"`
}

func (a *addressData) getStreet() string {
	return a.Street
}

func (a *addressData) getDistrict() string {
	return a.District
}

func (a *addressData) getSubStreet() string {
	return a.SubStreet
}

func (a *addressData) getFullAddress() string {
	return a.FullAddress
}

func (a *addressData) String() string {
	return a.FullAddress
}