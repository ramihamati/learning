module advanced_todo_list_addr::advance_to_do_list {
    use std::signer;
    use std::string::String;

    struct UserToDoListCounter has key {
        counter: u64
    }

    struct TodoList has key {
        owner: address,
        todos: vector<ToDo>
    }

    struct ToDo has key {
        content: String,
        completed: bool
    }

    // This function is only called once when the module is published for the first time.
    // init_module is optional, you can also have an entry function as the initializer.
    fun _init_module(_module_publisher: &signer) acquires UserToDoListCounter {
        // nothing to do here
    }

    public entry fun create_todo_list(sender: &signer) acquires UserToDoListCounter {
        let sender_address = signer::address_of(sender);
        let counter = if (exists<UserToDoListCounter>(sender_address)) {
            let counter = borrow_global<UserToDoListCounter>(sender_address);
            counter.counter
        }else {
            let counter = UserToDoListCounter { counter: 0 };
            move_to(sender, counter);
            0
        };
        
    }
}