package geojsoncoords

import (
	"testing"
)

func TestUnmarshalFeatureCollection(t *testing.T) {
	cases := []struct {
		RawObject        []byte
		ExpectedFeatures int
		ExpectedType     string
	}{
		{
			[]byte(`
                {   "type": "FeatureCollection",
                    "features": []
                }
            `),
			0,
			"FeatureCollection",
		},
		{
			[]byte(`
                {   "type": "FeatureCollection",
                    "features": [
                        {
                            "type": "Point",
                            "coordinates": [101.0, 102.0]
                        },
                        {
                            "type": "LineString",
                            "coordinates": [
                                [1.01, 0.10203],
                                [1.93, -1.1]
                            ]
                        },
                        {   
                            "type": "Feature",
                            "geometry": {
                                "type": "Point",
                                "coordinates": [123.4, 5.6789]
                            }
                        }
                    ]
                }
            `),
			3,
			"FeatureCollection",
		},
		{
			[]byte(`
                {
                    "type": "Feature",
                    "geometry": {
                        "type": "Point",
                        "coordinates": [125.6, 10.1]
                    },
                    "properties": {
                        "name": "Dinagat Islands"
                    }
                }
            `),
			0,
			"Feature",
		},
	}

	for _, c := range cases {
		result, err := UnmarshalFeatureCollection(c.RawObject)
		if err != nil {
			t.Errorf("Error unmarshalling FeatureCollectionJSON %s", err)
		}

		if c.ExpectedFeatures != len(result.Features) {
			t.Errorf("Unexpected number of feature objects. Got %d, expected %d", len(result.Features), c.ExpectedFeatures)
		}

		if c.ExpectedType != result.Type {
			t.Errorf("Unexpected FeatureCollection type. Got %s, expected %s", result.Type, c.ExpectedType)
		}
	}
}

func TestUnmarshalGeometry(t *testing.T) {
	cases := []struct {
		RawObject          []byte
		ExpectedGeometries int
		ExpectedType       string
	}{
		{
			[]byte(`
                {
                    "type": "GeometryCollection",
                    "geometries": [
                        {
                            "type": "Point",
                            "coordinates": [1.0, 2.0]
                        },
                        {
                            "type": "MultiLineString", 
                            "coordinates": [[[1,2],[3,4]]]
                        }
                    ]
                }
            `),
			2,
			"GeometryCollection",
		},
		{
			[]byte(`
                {
                    "type": "GeometryCollection",
                    "geometries": []
                }
            `),
			0,
			"GeometryCollection",
		},
		{
			[]byte(`
                {
                   "type": "Feature",
                   "geometry": {
                       "type": "LineString",
                       "coordinates": [
                           [100.0, 0.0],
                           [101.0, 1.0]
                       ]
                   },
                   "properties": {
                       "prop0": "value0",
                       "prop1": "value1"
                   }
                }
            `),
			0,
			"Feature",
		},
	}

	for _, c := range cases {
		result, err := UnmarshalGeometry(c.RawObject)
		if err != nil {
			t.Errorf("Error unmarshalling FeatureCollectionJSON %s", err)
		}

		if c.ExpectedGeometries != len(result.Geometries) {
			t.Errorf("Unexpected number of feature objects. Got %d, expected %d", len(result.Geometries), c.ExpectedGeometries)
		}

		if c.ExpectedType != result.Type {
			t.Errorf("Unexpected FeatureCollection type. Got %s, expected %s", result.Type, c.ExpectedType)
		}
	}
}
