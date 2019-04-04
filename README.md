# geojsoncoords

Extract coordinates from GeoJSON objects

## Install

Install with
```sh
go get github.com/fegoa89/geojsoncoords
```


### UnmarshalFeatureCollection
Converts a byte array that contains a GeoJSON object into FeatureCollection struct.

```golang
rawJSON := []byte(`
{ "type": "FeatureCollection",
  "features": [
    {"type": "Point", "coordinates": [101.0, 102.0]},
    {"type": "LineString", "coordinates": [[1.01, 0.10203],[1.93, -1.1]]},
    {"type": "Feature", "geometry": {"type": "Point", "coordinates": [123.4, 5.6789]}}
  ]
}
`)

featureCollection, err := UnmarshalFeatureCollection(rawJSON)
```

### UnmarshalGeometry
Converts a byte array that contains a GeoJSON object into Geometry struct.

```golang
rawJSON := []byte(`
{ "type": "GeometryCollection",
  "geometries": [
    {"type": "Point", "coordinates": [1.0, 2.0]},
	{"type": "MultiLineString", "coordinates": [[[1,2],[3,4]]]}
  ]
}
`)

geometryCollection, err := UnmarshalGeometry(rawJSON)
```
