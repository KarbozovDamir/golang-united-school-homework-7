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
    var strMatrix string ="test"
	result1, err1 := New(strMatrix)
	if result1 != nil && err1 != strconv.ErrSyntax {
		t.Errorf("wrong parsing matrix")
	}

	result2, err2 := New("1 2 \n 3")
	if result2 != nil && err2.Error() != "rows have to be same length" {
		t.Errorf("wrong the Matrix")
	}

	result3, err3 := New("1 2 \n 3 4")
	expect3 := &Matrix{2, 2, []int{1,2,3,4}}
	
	if result3.cols != expect3.cols || result3.rows != expect3.rows || !reflect.DeepEqual(result3.data, expect3.data) || err3 != nil {
		t.Errorf("wrong the new added Matrix")
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

	var m = &Matrix{
		rows: 1, 
		cols: 10, 
		data: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9,10},
	}
	
	rows, cols := 1, 10
	data := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	if len(m.Rows()) != rows {
		t.Errorf("return incorrect rows : get %v, should %v", len(m.Rows()), rows)
	}

	if len(m.Cols()) != cols {
		t.Errorf("return incorrect cols: get %v, should %v", len(m.Rows()), cols)
	}

	if !reflect.DeepEqual(data, m.data) {
		t.Errorf("return  incorrect data: get %v, should %v", m.data, data)
	}

	if !m.Set(0,1,10) {
		t.Errorf("return incorrect data")
	}

	if m.Set(1,0,10) {
		t.Errorf("set wrong col or row")
	}
}
