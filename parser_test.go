package libreofficedrawparser

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestRead(t *testing.T) {
	testCases := []struct {
		filename      string
		expected      ParsingResult
		expectedError bool
	}{
		{
			filename: "process_diagram_1.fodg",
			expected: ParsingResult{
				Vertices: []Vertex{
					{
						ID:          "id1",
						Name:        "",
						Description: "Description of the process start symbol",
					},
					{
						ID:          "id2",
						Name:        "Activity 1",
						Description: "Description of activity 1",
					},
					{
						ID:          "id3",
						Name:        "Activity 2",
						Description: "Description of activity 2",
					},
					{
						ID:          "id4",
						Name:        "",
						Description: "Description of the opening exclusive gateway",
					},
					{
						ID:          "id5",
						Name:        "Activity 3",
						Description: "Description of activity 3, line 1.\nLine 2.\nLine 3.",
					},
					{
						ID:          "id6",
						Name:        "",
						Description: "Description of the closing exclusive gateway",
					},
					{
						ID:          "id7",
						Name:        "",
						Description: "Description of the process end symbol",
					},
				},
				Edges: []Edge{
					{
						Source: "id1", // start
						Target: "id2", // activity 1
						Label:  "",
					},
					{
						Source: "id2", // activity 1
						Target: "id3", // activity 2
						Label:  "",
					},
					{
						Source: "id3", // activity 2
						Target: "id4", // opening exclusive gateway
						Label:  "",
					},
					{
						Source: "id4", // opening exclusive gateway
						Target: "id5", // activity 3
						Label:  "Arrow from gateway to activity 3",
					},
					{
						Source: "id5", // activity 3
						Target: "id6", // closing gateway
						Label:  "",
					},
					{
						Source: "id4", // opening exclusive gateway
						Target: "id6", // closing exclusive gateway
						Label:  "Стрелка от шлюза к шлюзу",
					},
					{
						Source: "id6", // closing exclusive gateway
						Target: "id7", // end
						Label:  "",
					},
				},
			},
			expectedError: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.filename, func(t *testing.T) {
			//
			// Given
			//
			// 1. Compose the file path
			path := filepath.Join("testdata", tc.filename)

			// 2. Open the file
			file, err := os.Open(path)
			if err != nil {
				t.Fatal(err)
			}
			defer file.Close()

			// 3. Create a Parser
			parser := NewParser()

			//
			// When
			//
			// 4. Pass the file to Read
			// 5. Save the result and error
			actual, err := parser.Read(file)

			//
			// Then
			//
			// Assert error value is correct
			if (err != nil) != tc.expectedError {
				t.Errorf("expectedError = %v, got err = %v", tc.expectedError, err)
			}
			// Assert the graph data is correct
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("expected: %+v\ngot:      %+v", tc.expected, actual)
			}
		})
	}
}
