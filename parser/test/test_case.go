package parser

var letStmt_input = `
	let x = 5;
	let y = 10;
	let foobar = 838383;
	`
var letStmt_expIdents = []struct {
	Name string
}{
	{"x"},
	{"y"},
	{"foobar"},
}

var retStmt_input = `
	return 5;
	return 10;
	return 993322;
	`
