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
            {   "type": "Feature",
                "geometry": {
                    "type": "Point",
                    "coordinates": [1.0, 0.5]
                }
            },
            {   "type": "Feature",
                "geometry": {
                    "type": "LineString",
                    "coordinates": [
                        [1.0, 3.55], [1.03, 3.4533], [1.04, 3.545], [1.05, 3.455]
                    ]},
                    "properties": {
                        "prop": 0.0
                    }
            },
            {   "type": "Feature",
                "geometry": {
                    "type": "Polygon",
                    "coordinates": [
                        [
                            [1.0, 3.0], [1.01, -2.56], [1.01, -1.5646],
                            [1.9786, 1.5345], [1.10, 0.43430]
                        ]
                    ]},
                    "properties": {
                        "prop": {"hello": "world"}
                    }
           }
        ]
    }
`)

result, err := geojsoncoords.ExtractCoordinates(rawJSON)
// result -> [
//              [1, 0.5], [1, 3.55], [1.03, 3.4533], [1.04, 3.545], [1.05, 3.455],
//              [1, 3], [1.01, -2.56], [1.01, -1.5646], [1.9786, 1.5345], [1.1, 0.4343]
//          ]
```
