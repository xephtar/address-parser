package addressParser

import (
    "strings"
    "fmt"
)

func ParseStringAddress(address string) *addressData {
    dist_re := getDistrictParser();
    street_re := getStreetParser();
    subStreet_re := getSubStreetParser();

    tmp := address

    districtName := ""
    streetName := ""
    subStreetName := ""

    if mah_result := dist_re.FindAllString(tmp, -1); len(mah_result) > 0 {
        tmp = strings.Replace(tmp, mah_result[0], "", 1)
        districtName = mah_result[0]
    }

    if cad_result := street_re.FindAllString(tmp, -1); len(cad_result) > 0 {
        tmp = strings.Replace(tmp, cad_result[0], "", 1)
        streetName = cad_result[0]
    }

    if sok_result := subStreet_re.FindAllString(tmp, -1); len(sok_result) > 0 {
        subStreetName = sok_result[0]
    }

    full_address := fmt.Sprintf("%s %s %s", districtName, streetName, subStreetName)

    return &addressData {
        Street: streetName,
        District: districtName,
        SubStreet: subStreetName,
        FullAddress: full_address,
    }
}