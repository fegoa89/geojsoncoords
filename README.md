# geojsoncoords

Extract coordinates from GeoJSON objects

## Install

Install with
```sh
go get github.com/fegoa89/geojsoncoords
```

### ExtractCoordinates
Returns a single slice of coordinates that the GeoJSON object contains

```golang
rawJSON := []byte(`
    {   "type": "FeatureCollection",
    "features": [
        { "type": "Feature",
        "geometry": {
            "type": "Polygon",
            "coordinates": [
                [
                    [1.0, 3.0], [1.01, -2.56], [1.01, -1.5646],
                    [1.9786, 1.5345], [1.10, 0.43430]
                ]
            ]}
        }]
    }
`)

result, err := geojsoncoords.ExtractCoordinates(rawJSON)
// result -> [[1.0, 3.0], [1.01, -2.56], [1.01, -1.5646] [1.9786, 1.5345], [1.10, 0.43430]]
```
