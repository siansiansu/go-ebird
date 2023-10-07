package ebird

import (
	"net/url"
	"strconv"
)

type ClientOption func(client *Client)

type RequestOption func(*requestOptions)

type requestOptions struct {
	urlParams url.Values
}

func processOptions(options ...RequestOption) requestOptions {
	o := requestOptions{
		urlParams: url.Values{},
	}
	for _, opt := range options {
		opt(&o)
	}
	return o
}

// Only fetch hotspots which have been visited up to 'back' days ago.
// Values: 1-30
func Back(amount int) RequestOption {
	return func(o *requestOptions) {
		o.urlParams.Set("back", strconv.Itoa(amount))
	}
}

// Only fetch records from these taxonomic categories.
// Values: any available category, must be lowercase
func Cat(code string) RequestOption {
	return func(o *requestOptions) {
		o.urlParams.Set("cat", code)
	}
}

// The characters used to separate elements in the name.
// Values: (any characters)
func Delim(code string) RequestOption {
	return func(o *requestOptions) {
		o.urlParams.Set("delim", code)
	}
}

// The search radius from the given position, in kilometers.
// Values: 0 - 500
func Dist(amount int) RequestOption {
	return func(o *requestOptions) {
		o.urlParams.Set("dist", strconv.Itoa(amount))
	}
}

// Fetch the records in CSV or JSON format.
// Values: csv, json
func Fmt(code string) RequestOption {
	return func(o *requestOptions) {
		o.urlParams.Set("fmt", code)
	}
}

// Locale for species group names. English names are returned for any
// non-listed locale or any non-translated group name
// Values: bg,cs,da,de,en,es,es_AR,es_CL,es_CU,es_ES,es_MX,es_PA,fr,he,
// is,nl,no,pt_BR,pt_PT,ru,sr,th,tr, or zh
func GroupNameLocale(code string) RequestOption {
	return func(o *requestOptions) {
		o.urlParams.Set("groupNameLocale", code)
	}
}

// Only fetch observations from hotspots.
// Values: true, false
func Hotspot(code string) RequestOption {
	return func(o *requestOptions) {
		o.urlParams.Set("hotspot", code)
	}
}

// Include observations which have not yet been reviewed.
// Values: true, false
func IncludeProvisional(code string) RequestOption {
	return func(o *requestOptions) {
		o.urlParams.Set("includeProvisional	", code)
	}
}

// Required. Latitude to 2 decimal places.
// Values: -90 - 90
func Lat(amount int) RequestOption {
	return func(o *requestOptions) {
		o.urlParams.Set("lat", strconv.Itoa(amount))
	}
}

// Required. Longitude to 2 decimal places.
// Values: -180 - 180
func Lng(amount int) RequestOption {
	return func(o *requestOptions) {
		o.urlParams.Set("lng", strconv.Itoa(amount))
	}
}

// Use this language for common names.
// Values: any available locale
func Locale(code string) RequestOption {
	return func(o *requestOptions) {
		o.urlParams.Set("locale", code)
	}
}

// Only fetch this number of contributors.
// Values: 1 - 100
func MaxResults(amount int) RequestOption {
	return func(o *requestOptions) {
		o.urlParams.Set("maxResults", strconv.Itoa(amount))
	}
}

// Fetch observations from up to 10 locations.
// Values: any location code
func R(code string) RequestOption {
	return func(o *requestOptions) {
		o.urlParams.Set("r", code)
	}
}

// Order by number of complete checklists (cl) or by number of species seen (spp).
// Values: spp, cl
func RankedBy(code string) RequestOption {
	return func(o *requestOptions) {
		o.urlParams.Set("rankedBy", code)
	}
}

// Control how the name is displayed.
// Values: detailed, detailednoqual, full, namequal, nameonly, revdetailed
func RegionNameFormat(code string) RequestOption {
	return func(o *requestOptions) {
		o.urlParams.Set("regionNameFormat", code)
	}
}

// Order the results by the date of the checklist or by the date it was submitted.
// Values: obs_dt, creation_dt
func SortKey(code string) RequestOption {
	return func(o *requestOptions) {
		o.urlParams.Set("sortKey", code)
	}
}

// Only fetch records for these species.
// Values: any species code
func Species(code string) RequestOption {
	return func(o *requestOptions) {
		o.urlParams.Set("species", code)
	}
}

// Use this language for species common names.
// Values: any available locale
func SppLocale(code string) RequestOption {
	return func(o *requestOptions) {
		o.urlParams.Set("sppLocale", code)
	}
}

// Fetch a specific version of the taxonomy.
// any available version
func Version(code string) RequestOption {
	return func(o *requestOptions) {
		o.urlParams.Set("version", code)
	}
}
