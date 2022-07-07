package coverage

import (
	"os"
	"sort"
	"testing"
	"time"
	"reflect"
	"strconv"
	
)

// DO NOT EDIT THIS FUNCTION
func init() {
	content, err := os.ReadFile("students_test.go")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("autocode/students_test", content, 0644)
	if err != nil {
		panic(err)
	}
}

func TestLen(t *testing.T) {
	people := People{
		{firstName: "Serik",lastName: "Karbozov",birthDay: time.Date(1960,3,13,0,0,0,0,time.UTC)},
		{firstName: "Damir", lastName: "Karbozov", birthDay: time.Date(1989,9,4,0,0,0,0,time.UTC)},
		{firstName: "Timur", lastName: "Karbozov", birthDay: time.Date(1986,8,23,0,0,0,0,time.UTC)},
		{firstName: "Faiz", lastName: "Karbozov", birthDay: time.Date(1989,9,4,0,0,0,0,time.UTC)},
		{firstName: "Faiz", lastName: "Galymzhanov", birthDay: time.Date(1989,9,4,0,0,0,0,time.UTC)},

}
	if people.Len() != 5 {
		t.Errorf("not correct func Len()")
	} 
}


func TestSort(t *testing.T) {
	exp := []string{"DamirKarbozov", "FaizGalymzhanov","FaizKarbozov", "TimurKarbozov", "SerikKarbozov"}

	res := People{}
	res = append(res, Person{firstName: "Serik", lastName: "Karbozov", birthDay: time.Date(1960,3,13,0,0,0,0,time.UTC)})
	res = append(res, Person{firstName: "Damir", lastName: "Karbozov", birthDay: time.Date(1989,9,4,0,0,0,0,time.UTC)})
	res = append(res, Person{firstName: "Timur", lastName: "Karbozov", birthDay: time.Date(1986,8,23,0,0,0,0,time.UTC)})
	res = append(res, Person{firstName: "Faiz", lastName: "Karbozov", birthDay: time.Date(1989,9,4,0,0,0,0,time.UTC)})
	res = append(res, Person{firstName: "Faiz", lastName: "Galymzhanov", birthDay: time.Date(1989,9,4,0,0,0,0,time.UTC)})

	sort.Sort(People(res))

	for i, el := range exp{
		if el != res[i].firstName + res[i].lastName{
			t.Errorf("expected %s but got %s", el, res[i].firstName)
		}
	}
}
func TestNew(t *testing.T) {
	//t.Parallel()

	result1, err1 := New("cover")
	if result1 != nil && err1 != strconv.ErrSyntax {
		t.Errorf("Wrong parsing matrix")
	}

	result2, err2 := New("1 2 \n 3")
	if result2 != nil && err2.Error() != "Rows need to be the same length" {
		t.Errorf("Wrong Matrix")
	}

	result3, err3 := New("1 2 \n 3 4")
	expect3 := &Matrix{2, 2, []int{1,2,3,4}}
	
	if result3.cols != expect3.cols || result3.rows != expect3.rows || !reflect.DeepEqual(result3.data, expect3.data) || err3 != nil {
		t.Errorf("Wrong new Matrix")
	}
}


func TestRows(t *testing.T) {
	m := Matrix{
		rows: 2,
		cols: 5,
		data: []int{1, 2, 3, 4, 5, 
			6, 7, 8, 9, 10},
	}
	expMtrx := [][]int{
		{ 1, 2, 3, 4, 5},
		{ 6, 7, 8, 9, 10},
	}
	got := m.Rows()
	if !reflect.DeepEqual(got, expMtrx){
		t.Error("testing error: rows are not equal")
	}
	
}

func TestCols(t *testing.T) {
	m := Matrix{
		rows: 2,
		cols: 5,
		data: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	}

	expMtrx := [][]int{	
		{1, 6},	{2, 7},	{3, 8},	{4, 9},	{5, 10},
	}
	got := m.Cols()
	if !reflect.DeepEqual(got, expMtrx){
		t.Errorf("testing error: cols are not equal\n expMtrx: %v\ngot: %v", expMtrx, got)
	}
	
}


func TestSet(t *testing.T) {
	m := &Matrix{rows: 1, cols: 10, data: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}}
	matrix := map[string]map[string]int{
	"normal" : {
		"col" : 10,
		"row" : 1,
		"value" : 10,
		"res" : 10,
	},
	"err_col" : {
		"col" : 10,
		"row" : 10,
		"value" : 7,
		"res" : 10,
	},
	"err_row" : {
		"col" : 2,
		"row" : 2,
		"value" : 2,
		"res" : 0,
	},
	"err_minus" : {
		"col" : -4,
		"row" : 0,
		"value" : 7,
		"res" : 0,
	},
}

for k, v := range matrix {
	name, obj := k, v
	t.Run(name, func(t *testing.T) {
		t.Parallel()
		ans := obj["res"] 
		got := m.Set(obj["row"], obj["col"], obj["value"])
		if got && ans == 0 || !got && ans == 1 {
			t.Errorf("expected \"%t\" but got \"%t\"",!got, got)
		}
	})
}
}

