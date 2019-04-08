// Package geojsoncoords extract coordinates from GeoJSON objects
package geojsoncoords

import (
	"reflect"
)

// ExtractCoordinates return a single slice of coordinates that
// the GeoJSON object contains
func ExtractCoordinates(geojson []byte) ([][]float64, error) {
	fc, err := UnmarshalFeatureCollection(geojson)
	if err != nil {
		return nil, err
	}

	result := [][]float64{}
	for _, s := range fc.Features {
		if s.Coordinates != nil {
			coordsSlice := FlattenCoordsTo2DSlice(s.Coordinates)
			result = append(result, coordsSlice...)
		}

		if s.Geometry != nil {
			coordsSlice := FlattenCoordsTo2DSlice(s.Geometry.Coordinates)
			result = append(result, coordsSlice...)
		}
	}

	return result, nil
}

// FlattenCoordsTo2DSlice return a single slice of coordinates from a slice of slices
func FlattenCoordsTo2DSlice(coordinates ...interface{}) (coordsSlice [][]float64) {
	flattenCoords := FlattenSlice(coordinates)
	for i := 0; i < len(flattenCoords); i = i + 2 {
		coords := []float64{}
		coords = append(coords, flattenCoords[i].(float64))
		coords = append(coords, flattenCoords[i+1].(float64))
		coordsSlice = append(coordsSlice, coords)
	}
	return coordsSlice
}

// FlattenSlice recursively flattens a nested interface
func FlattenSlice(args ...interface{}) (list []interface{}) {
	for _, v := range args {
		if reflect.TypeOf(v).Kind() == reflect.Slice {
			for _, z := range FlattenSlice((v.([]interface{}))...) {
				list = append(list, z)
			}
		} else {
			list = append(list, v)
		}
	}
	return list
}
