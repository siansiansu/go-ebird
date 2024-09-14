package ebird

import (
	"net/url"
	"strconv"
	"strings"
)

type ClientOption func(*Client)

type RequestOption func(*RequestOptions)

type RequestOptions struct {
	URLParams url.Values
}

func processOptions(options ...RequestOption) RequestOptions {
	o := RequestOptions{
		URLParams: url.Values{},
	}
	for _, opt := range options {
		opt(&o)
	}
	return o
}

func Back(days int) RequestOption {
	return func(o *RequestOptions) {
		if days > 0 && days <= 30 {
			o.URLParams.Set("back", strconv.Itoa(days))
		}
	}
}

func Cat(category string) RequestOption {
	return func(o *RequestOptions) {
		o.URLParams.Set("cat", category)
	}
}

func Delim(delimiter string) RequestOption {
	return func(o *RequestOptions) {
		o.URLParams.Set("delim", delimiter)
	}
}

func Dist(distance int) RequestOption {
	return func(o *RequestOptions) {
		if distance >= 0 && distance <= 500 {
			o.URLParams.Set("dist", strconv.Itoa(distance))
		}
	}
}

func Fmt(format string) RequestOption {
	return func(o *RequestOptions) {
		if format == "csv" || format == "json" {
			o.URLParams.Set("fmt", format)
		}
	}
}

func GroupNameLocale(locale string) RequestOption {
	return func(o *RequestOptions) {
		o.URLParams.Set("groupNameLocale", locale)
	}
}

func Hotspot(isHotspot bool) RequestOption {
	return func(o *RequestOptions) {
		o.URLParams.Set("hotspot", strconv.FormatBool(isHotspot))
	}
}

func IncludeProvisional(include bool) RequestOption {
	return func(o *RequestOptions) {
		o.URLParams.Set("includeProvisional", strconv.FormatBool(include))
	}
}

func Lat(latitude float64) RequestOption {
	return func(o *RequestOptions) {
		if latitude >= -90 && latitude <= 90 {
			o.URLParams.Set("lat", strconv.FormatFloat(latitude, 'f', 2, 64))
		}
	}
}

func Lng(longitude float64) RequestOption {
	return func(o *RequestOptions) {
		if longitude >= -180 && longitude <= 180 {
			o.URLParams.Set("lng", strconv.FormatFloat(longitude, 'f', 2, 64))
		}
	}
}

func Locale(locale string) RequestOption {
	return func(o *RequestOptions) {
		o.URLParams.Set("locale", locale)
	}
}

func MaxResults(max int) RequestOption {
	return func(o *RequestOptions) {
		if max > 0 && max <= 100 {
			o.URLParams.Set("maxResults", strconv.Itoa(max))
		}
	}
}

func R(locations ...string) RequestOption {
	return func(o *RequestOptions) {
		if len(locations) > 0 && len(locations) <= 10 {
			o.URLParams.Set("r", strings.Join(locations, ","))
		}
	}
}

func RankedBy(rankMethod string) RequestOption {
	return func(o *RequestOptions) {
		if rankMethod == "spp" || rankMethod == "cl" {
			o.URLParams.Set("rankedBy", rankMethod)
		}
	}
}

func RegionNameFormat(format string) RequestOption {
	return func(o *RequestOptions) {
		validFormats := []string{"detailed", "detailednoqual", "full", "namequal", "nameonly", "revdetailed"}
		if contains(validFormats, format) {
			o.URLParams.Set("regionNameFormat", format)
		}
	}
}

func SortKey(key string) RequestOption {
	return func(o *RequestOptions) {
		if key == "obs_dt" || key == "creation_dt" {
			o.URLParams.Set("sortKey", key)
		}
	}
}

func Species(speciesCodes ...string) RequestOption {
	return func(o *RequestOptions) {
		if len(speciesCodes) > 0 {
			o.URLParams.Set("species", strings.Join(speciesCodes, ","))
		}
	}
}

func SppLocale(locale string) RequestOption {
	return func(o *RequestOptions) {
		o.URLParams.Set("sppLocale", locale)
	}
}

func Version(version string) RequestOption {
	return func(o *RequestOptions) {
		o.URLParams.Set("version", version)
	}
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
