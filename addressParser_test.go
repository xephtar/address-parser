package addressParser

import "testing"

func TestParseStringAddress(t *testing.T) {
	address := "Haydarali mahallesi Dudullu cad 120.sok no:3 kat:5"
	want := "haydarali mahallesi dudullu cad 120.sok"
    parsedAddress, _ := ParseStringAddress(address)

	if parsedAddress.GetFullAddress() != want {
		t.Errorf("The parsed address is incorrect, got: %s, want: %s.", parsedAddress, want)
	}
}