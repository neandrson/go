package main

import (
	"testing"
)

func TestWorkerSort(t *testing.T) {

	type Worker struct {
		Name            string
		Position        string
		Salary          uint
		ExperienceYears uint
	}

	tests := []struct {
		workers  []Worker
		expected []string
	}{
		{
			workers: []Worker{
				{Name: "Михаил", Position: "директор", Salary: 200, ExperienceYears: 5},
				{Name: "Игорь", Position: "зам. директора", Salary: 180, ExperienceYears: 3},
				{Name: "Николай", Position: "начальник цеха", Salary: 120, ExperienceYears: 2},
				{Name: "Андрей", Position: "мастер", Salary: 90, ExperienceYears: 10},
				{Name: "Виктор", Position: "рабочий", Salary: 80, ExperienceYears: 3},
			},
			expected: []string{
				"Михаил - 12000 - директор",
				"Андрей - 10800 - мастер",
				"Игорь - 6480 - зам. директора",
				"Николай - 2880 - начальник цеха",
				"Виктор - 2880 - рабочий",
			},
		},
		{
			workers: []Worker{
				{Name: "Андрей", Position: "директор", Salary: 200, ExperienceYears: 1},
				{Name: "Максим", Position: "зам. директора", Salary: 180, ExperienceYears: 3},
				{Name: "Николай", Position: "начальник цеха", Salary: 120, ExperienceYears: 2},
				{Name: "Андрей", Position: "мастер", Salary: 90, ExperienceYears: 10},
				{Name: "Виктор", Position: "рабочий", Salary: 80, ExperienceYears: 3},
			},
			expected: []string{
				"Михаил - 12000 - директор",
				"Андрей - 10800 - мастер",
				"Игорь - 6480 - зам. директора",
				"Николай - 2880 - начальник цеха",
				"Виктор - 2880 - рабочий",
			},
		},
	}
	/*for _, tc := range tests {
		AddWorkerInfo(tc.workers)
		if slices.Compare(tc.expected, tc.workers) != 0 {
			t.Errorf("TestSortNames failed. Expected: %v, Got: %v", tc.expected, tc.names)
		}
	}*/
	for _, tc := range tests {
		t.Run(tc.workers, func(t *testing.T) {
			worker, err := AddWorkerInfo(tc.workers)
			if err != nil {
				t.Fatal(err)
			}

			if worker != tc.expected {
				t.Errorf("Expected status code %d, but got %d", tc.expected, worker)
			}
		})
	}
}
