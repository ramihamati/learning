// trait Animal {
//     fn speak(&self);
//     fn name(&self) -> &str;
// }
//
// // Implement the trait for a concrete type `Dog`
// struct Dog {
//     name: String,
// }
//
// impl Animal for Dog {
//     fn speak(&self) {
//         println!("Woof! My name is {}", self.name);
//     }
//
//     fn name(&self) -> &str {
//         &self.name
//     }
// }
//
// // Implement the trait for a concrete type `Cat`
// struct Cat {
//     name: String,
// }
//
// impl Animal for Cat {
//     fn speak(&self) {
//         println!("Meow! My name is {}", self.name);
//     }
//
//     fn name(&self) -> &str {
//         &self.name
//     }
// }
//
// struct VTable {
//     speak: fn(&dyn Animal),
//     name: fn(&dyn Animal) -> &str,
// }
//
// // Example vtables for Dog and Cat
// static DOG_VTABLE: VTable = VTable {
//     speak: Dog::speak as fn(&dyn Animal),
//     name: Dog::name as fn(&dyn Animal) -> &str,
// };
//
// static CAT_VTABLE: VTable = VTable {
//     speak: Cat::speak as fn(&dyn Animal),
//     name: Cat::name as fn(&dyn Animal) -> &str,
// };
//
// // Example fat pointer (trait object)
// struct FatPointer {
//     data: *const u8, // Pointer to the actual data
//     vtable: &'static VTable, // Pointer to the vtable
// }
//
// fn test() {
//     let dog = Dog { name: String::from("Buddy") };
//     let cat = Cat { name: String::from("Whiskers") };
//
//     // Create a trait object for dog
//     let dog_trait_object = FatPointer {
//         data: &dog as *const Dog as *const u8,
//         vtable: &DOG_VTABLE,
//     };
//
//     // Create a trait object for cat
//     let cat_trait_object = FatPointer {
//         data: &cat as *const Cat as *const u8,
//         vtable: &CAT_VTABLE,
//     };
//
//     // Dynamic dispatch
//     (dog_trait_object.vtable.speak)(unsafe { &*(dog_trait_object.data as *const dyn Animal) });
//     (cat_trait_object.vtable.speak)(unsafe { &*(cat_trait_object.data as *const dyn Animal) });
// }