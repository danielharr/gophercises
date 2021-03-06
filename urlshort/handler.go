package urlshort

import (
	"net/http"

	"gopkg.in/yaml.v2"
)

type mapping struct {
	Path string
	Url  string
}

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if url, exists := pathsToUrls[r.URL.Path]; exists {
			http.Redirect(rw, r, url, http.StatusPermanentRedirect)
		} else {
			fallback.ServeHTTP(rw, r)
		}
	}
}

func parseYAML(yml []byte) ([]mapping, error) {
	var mappings []mapping
	err := yaml.UnmarshalStrict(yml, &mappings)
	return mappings, err
}

func buildMap(mappings []mapping) map[string]string {
	result := make(map[string]string)
	for _, m := range mappings {
		result[m.Path] = m.Url
	}
	return result
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yaml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	mappings, err := parseYAML(yaml)
	if err != nil {
		return nil, err
	}
	return MapHandler(buildMap(mappings), fallback), nil
}
