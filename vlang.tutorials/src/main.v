module main
import persons

fn main() {
	println('Hello World!')
	a, b := foo();
	println(a)
	println(b)

	p:= persons.new_person(	'john', 'don');
	
	println(p.get_last_name())
}

fn foo() (int, int){
	return 1, 2
}