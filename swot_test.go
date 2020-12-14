package swot

import (
	"testing"
)

func TestIsAcademic(t *testing.T) {
	var tests = []struct {
		address string
		result  bool
	}{
		{"lreilly@stanford.edu", true},
		{"LREILLY@STANFORD.EDU", true},
		{"Lreilly@Stanford.Edu", true},
		{"lreilly@slac.stanford.edu", true},
		{"lreilly@strath.ac.uk", true},
		{"lreilly@soft-eng.strath.ac.uk", true},
		{"lee@ugr.es", true},
		{"lee@uottawa.ca", true},
		{"lee@mother.edu.ru", true},
		{"lee@ucy.ac.cy", true},
		{"stanford.edu", true},
		{"slac.stanford.edu", true},
		{"www.stanford.edu", true},
		{"http://www.stanford.edu", true},
		{"http://www.stanford.edu:9393", true},
		{"strath.ac.uk", true},
		{"soft-eng.strath.ac.uk", true},
		{"ugr.es", true},
		{"uottawa.ca", true},
		{"mother.edu.ru", true},
		{"ucy.ac.cy", true},
		{" stanford.edu", true},
		{"lee@strath.ac.uk ", true},
		{"lee@stud.uni-corvinus.hu", true},
		{"lee@harvard.edu", true},
		{"lee@mail.harvard.edu", true},
		{"lee@leerilly.net", false},
		{"lee@gmail.com", false},
		{"lee@stanford.edu.com", false},
		{"lee@strath.ac.uk.com", false},
		{"leerilly.net", false},
		{"gmail.com", false},
		{"stanford.edu.com", false},
		{"strath.ac.uk.com", false},
		{"", false},
		{"the", false},
		{" gmail.com ", false},
		{"si.edu", false},
		{" si.edu ", false},
		{"imposter@si.edu", false},
		{"foo.si.edu", false},
		{"america.edu", false},
		{"folger.edu", false},
		{"foo@fit.ac.cy", true},
		{"foo@stud.fit.ac.cy", true},
		{"foo@stud.frederick.ac.cy", true},
		{"foo@fit.ac.cy", true},
		{"foo@frederick.ac.cy", true},
		{"foo@uds.berlin", true},
	}

	for _, tt := range tests {
		address := tt.address
		result := IsAcademic(address)
		if result != tt.result {
			t.Errorf("%s got %t, want %t", address, result, tt.result)
		}
	}
}

func TestGetSchoolName(t *testing.T) {
	var tests = []struct {
		address string
		result  string
	}{
		{"lreilly@cs.strath.ac.uk", "University of Strathclyde"},
		{"lreilly@fadi.at", "BRG Fadingerstraße Linz, Austria"},
		{"foo@shop.com", ""},
		{"bar@gmail.com", ""},
		{"gmail.com", ""},
		{"harvard.edu", "Harvard University"},
		{"stanford.edu", "Stanford University"},
		{"http://www.stanford.edu:9393", "Stanford University"},
	}

	for _, tt := range tests {
		address := tt.address
		result := GetSchoolName(address)
		if result != tt.result {
			t.Errorf("%s got %s, want %s", address, result, tt.result)
		}
	}
}
