// module devaddr::advance_to_do_list {
//     use std::bcs;
//     use std::signer;
//     use std::string::String;
//     use std::vector;
//     use aptos_std::string_utils;
//     use aptos_framework::object;
//
//     struct UserToDoListCounter has key {
//         counter: u64
//     }
//
//     struct TodoList has key {
//         owner: address,
//         todos: vector<ToDo>
//     }
//
//     struct ToDo has store, drop, copy {
//         content: String,
//         completed: bool
//     }
//
//     // This function is only called once when the module is published for the first time.
//     // init_module is optional, you can also have an entry function as the initializer.
//     fun init_module(_module_publisher: &signer) {
//         // nothing to do here
//     }
//
//     // ======================== Write Functions ========================
//
//
//     public entry fun create_todo_list(sender: &signer) acquires UserToDoListCounter {
//         let sender_address = signer::address_of(sender);
//         let counter = if (exists<UserToDoListCounter>(sender_address)) {
//             let counter = borrow_global<UserToDoListCounter>(sender_address);
//             counter.counter
//         }else {
//             let counter = UserToDoListCounter { counter: 0 };
//             move_to(sender, counter);
//             0
//         };
//         // create a new object to hold the todolist
//         let obj_holds_todo_list = object::create_named_object(
//             sender,
//             construct_todo_list_object_seed(counter)
//         );
//         let obj_signer = object::generate_signer(&obj_holds_todo_list);
//         let todo_list = TodoList{
//             owner: sender_address,
//             todos: vector::empty()
//         };
//         // store the TodoList resource under the newly created object
//         move_to(&obj_signer, todo_list);
//         // increment the counter
//         let counter = borrow_global_mut<UserToDoListCounter>(sender_address);
//         counter.counter = counter.counter + 1;
//     }
//
//     // ======================== Read Functions ========================
//     #[view]
//     public fun get_todo_list_counter(sender: address) : u64 acquires UserToDoListCounter{
//         if (exists<UserToDoListCounter>(sender)){
//             let counter = borrow_global<UserToDoListCounter>(sender);
//             counter.counter
//         }else{
//             0
//         }
//     }
//
//     // =========== Helper ========= //
//     fun construct_todo_list_object_seed(counter: u64): vector<u8>{
//         // The seed must be unique per todo list creator
//         //We add contract address as part of the seed so seed from 2 todo list contract for same user would be different
//         bcs::to_bytes(&string_utils::format2(&b"{}_{}", @devaddr, counter))
//     }
// }