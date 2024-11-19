package main

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"
	"time"
)

func TestExtractLog(t *testing.T) {
	type test struct {
		fileContent string
		fileName    string
		outContent  []string
		wantError   bool
		start       time.Time
		end         time.Time
	}

	tests := []test{
		{
			fileContent: `12.12.2022 info
			13.12.2022 info
			14.12.2022 info
			15.12.2022 info
			16.12.2022 info`,
			start: time.Date(2022, 12, 14, 0, 0, 0, 0, time.UTC),
			end:   time.Date(2022, 12, 15, 0, 0, 0, 0, time.UTC),

			fileName: "file1.txt",
			outContent: []string{
				"14.12.2022 info",
				"15.12.2022 info",
			},
		},
		{
			fileContent: `12.12.2022 info
			13.12.2022 info
			14.12.2022 info
			15.12.2022 info
			16.12.2022 info`,
			start: time.Date(2022, 12, 19, 0, 0, 0, 0, time.UTC),
			end:   time.Date(2022, 12, 20, 0, 0, 0, 0, time.UTC),

			fileName:   "file2.txt",
			outContent: []string{},
			wantError:  true,
		},
		{
			fileContent: `12.12.2022 info
			13.12.2022 info
			14.12.2022 info
			15.12.2022 info
			16.12.2022 info`,
			fileName:   "file3.txt",
			outContent: []string{},
			wantError:  true,
		},
	}
	for _, tc := range tests {
		t.Run(tc.fileName, func(t *testing.T) {
			file, err := ioutil.TempFile("", tc.fileName)
			if tc.wantError {
				if err == nil {
					t.Errorf("Expected write error, but got none")
				}
				return
			}
			if err != nil {
				t.Fatalf("failed to create temporary file: %v", err)
			}
			defer os.Remove(file.Name())

			_, err = file.WriteString(tc.fileContent)
			if err != nil {
				t.Fatalf("failed to write to temporary file: %v", err)
			}

			gotContent, err := ExtractLog(file.Name(), tc.start, tc.end)
			if (err != nil) != tc.wantError {
				t.Fatalf("got error %v, want error %v", err, tc.wantError)
			}

			if !reflect.DeepEqual(gotContent, tc.outContent) {
				t.Errorf("got content %v, want content %v", gotContent, tc.outContent)
			}

		})
	}
}
