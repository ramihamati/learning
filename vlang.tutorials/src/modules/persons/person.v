module persons

struct Person{
	mut:
		first_name string	
		last_name string
	pub:
		this_is_public string
}

pub fn (p Person) get_last_name() string{
	return p.last_name
}

pub fn new_person(last_name string, first_name string) Person {
	return Person{
		first_name : first_name,
		last_name: last_name
	}
}