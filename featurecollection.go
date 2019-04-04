package geojsoncoords

import (
	"encoding/json"
)

// FeatureCollection maps to GeoJSON FeatureCollection object
type FeatureCollection struct {
	Type        string     `json:"type"`
	BoundingBox []float64  `json:"bbox,omitempty"`
	Features    []*Feature `json:"features"`
}

// Feature maps to GeoJSON feature object
type Feature struct {
	ID          interface{}            `json:"id,omitempty"`
	Type        string                 `json:"type"`
	Coordinates interface{}            `json:"coordinates,omitempty"`
	BoundingBox []float64              `json:"bbox,omitempty"`
	Geometry    *Geometry              `json:"geometry"`
	Properties  map[string]interface{} `json:"properties"`
}

// Geometry maps to GeoJSON geometry object
type Geometry struct {
	Type        string      `json:"type"`
	BoundingBox []float64   `json:"bbox,omitempty"`
	Coordinates interface{} `json:"coordinates,omitempty"`
	Geometries  []*Geometry `json:"geometries,omitempty"`
}

// UnmarshalFeatureCollection converts a byte array that contains a GeoJSON object
// into FeatureCollection struct
func UnmarshalFeatureCollection(geojson []byte) (*FeatureCollection, error) {
	featureCollection := &FeatureCollection{}
	err := json.Unmarshal(geojson, featureCollection)
	if err != nil {
		return nil, err
	}

	return featureCollection, nil
}

// UnmarshalGeometry converts a byte array that contains a GeoJSON object
// into Geometry struct
func UnmarshalGeometry(geojson []byte) (*Geometry, error) {
	geometry := &Geometry{}
	err := json.Unmarshal(geojson, geometry)
	if err != nil {
		return nil, err
	}

	return geometry, nil
}
