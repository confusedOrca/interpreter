package parser

var test_input = `
	let x = 5;
	let y = 10;
	let foobar = 838383;
	`
var expectedIdents = []struct {
	Name string
}{
	{"x"},
	{"y"},
	{"foobar"},
}
