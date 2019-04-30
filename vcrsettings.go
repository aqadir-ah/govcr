package govcr

import (
	"log"
	"net/http"
)

// Setting defines an optional functional parameter as received by NewVCR()
type Setting func(vcrConfig *VCRSettings)

// WithClient is an optional functional parameter to provide a VCR with
// a custom HTTP client.
func WithClient(httpClient *http.Client) Setting {
	return func(vcrConfig *VCRSettings) {
		vcrConfig.client = httpClient
	}
}

// WithCassette is an optional functional parameter to provide a VCR with
// a cassette to load.
func WithCassette(cassetteName string) Setting {
	return func(vcrConfig *VCRSettings) {
		k7, err := loadCassette(cassetteName)
		if err != nil {
			log.Printf("failed loading cassette %s': %s\n", cassetteName, err.Error())
			return
		}
		vcrConfig.cassette = k7
	}
}

// VCRSettings holds a set of options for the VCR.
type VCRSettings struct {
	client   *http.Client
	cassette *cassette

	// trackRecordingMutator mutatorOfSomeSortThatTakesATrack // only the exported fields of the track will be mutable, the others will be invisible
	// trackReplayingMutator mutatorOfSomeSortThatTakesATrack // only the exported fields of the track will be mutable, the others will be invisible
}