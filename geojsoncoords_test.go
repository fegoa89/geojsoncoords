package geojsoncoords

import (
	"reflect"
	"testing"
)

func TestExtractCoordinates(t *testing.T) {
	cases := []struct {
		RawObject        []byte
		ExpectedResponse [][]float64
	}{
		{
			[]byte(`
                { 	"type": "FeatureCollection",
				  	"features": [
				    	{ 	"type": "Feature",
				      		"geometry": {
				      			"type": "Point",
				      			"coordinates": [1.0, 0.5]
				      		}
				      	},
				    	{ 	"type": "Feature",
				      		"geometry": {
				        		"type": "LineString",
				        		"coordinates": [
				        			[1.0, 3.55], [1.03, 3.4533], [1.04, 3.545], [1.05, 3.455]
				        		]},
				      			"properties": {
				        			"prop": 0.0
				        		}
				      	},
				    	{ 	"type": "Feature",
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
            `),
			[][]float64{{1, 0.5}, {1, 3.55}, {1.03, 3.4533}, {1.04, 3.545}, {1.05, 3.455}, {1, 3}, {1.01, -2.56}, {1.01, -1.5646}, {1.9786, 1.5345}, {1.1, 0.4343}},
		},
		{
			[]byte(`
				{ 	"type": "FeatureCollection",
  					"features": [
    					{
    						"type": "Point", "coordinates": [1.20, 0.20]
    					}
  					]
  				}
  			`),
			[][]float64{{1.20, 0.20}},
		},
		{
			[]byte(`
				{ 	"type": "FeatureCollection",
  					"features": [
    					{ 	"type": "Feature",
				      		"geometry": {
				        		"type": "LineString",
				        		"coordinates": [
				        			[1.0, 3.55], [1.03, 3.4533]
				        		]}
				      	}
  					]
  				}
  			`),
			[][]float64{{1.0, 3.55}, {1.03, 3.4533}},
		},
		{
			[]byte(`
				{ 	"type": "FeatureCollection",
  					"features": [
    					{ 	"type": "Feature",
					      	"geometry": {
					      		"type": "Polygon",
					      		"coordinates": [
					        		[
					            		[1.0, 3.0], [1.01, -2.56], [1.01, -1.5646],
					            		[1.9786, 1.5345], [1.10, 0.43430]
					        		]
					      	]}
					    }
  					]
  				}
  			`),
			[][]float64{{1, 3}, {1.01, -2.56}, {1.01, -1.5646}, {1.9786, 1.5345}, {1.1, 0.4343}},
		},
	}

	for _, c := range cases {
		result, err := ExtractCoordinates(c.RawObject)
		if err != nil {
			t.Errorf("Error exctracting coordinates %s", err)
		}

		if reflect.DeepEqual(result, c.ExpectedResponse) != true {
			t.Errorf("Unexpected response. Got %v, want %v", result, c.ExpectedResponse)
		}
	}
}

func TestFlattenCoordsTo2DSlice(t *testing.T) {
	cases := []struct {
		Interface        interface{}
		ExpectedResponse [][]float64
	}{
		{
			[]interface{}{2.2, 0.5, 1.2, -2.4},
			[][]float64{{2.2, 0.5}, {1.2, -2.4}},
		},
		{
			[]interface{}{2.2, 0.5},
			[][]float64{{2.2, 0.5}},
		},
	}

	for _, c := range cases {
		result := FlattenCoordsTo2DSlice(c.Interface)
		if reflect.DeepEqual(result, c.ExpectedResponse) != true {
			t.Errorf("Unexpected response. Got %v, want %v", result, c.ExpectedResponse)
		}
	}
}

func TestFlattenSlice(t *testing.T) {
	cases := []struct {
		Interface        interface{}
		ExpectedResponse interface{}
	}{
		{
			[]interface{}{[]interface{}{2.2, 0.5}, []interface{}{1.7, -0.5}},
			[]interface{}{2.2, 0.5, 1.7, -0.5},
		},
		{
			[]interface{}{[]interface{}{[]interface{}{[]interface{}{0.9, -0.4}, []interface{}{0.1, 0.2}}}},
			[]interface{}{0.9, -0.4, 0.1, 0.2},
		},
	}

	for _, c := range cases {
		result := FlattenSlice(c.Interface)
		if reflect.DeepEqual(result, c.ExpectedResponse) != true {
			t.Errorf("Unexpected response. Got %v, want %v", result, c.ExpectedResponse)
		}
	}
}
