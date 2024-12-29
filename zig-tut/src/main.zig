const std = @import("std");
const crypto = @import("./crypto.zig");

const Person = struct {
    _name: []const u8,
    allocator: *const std.mem.Allocator,

    pub fn Init(name  : []const u8, allocator:  *const std.mem.Allocator) !*Person{
        const person = try allocator.create(Person);

        person.* = Person{
            ._name = try allocator.alloc(u8, name.len),
            .allocator = allocator
        };
        person._name = name;
        return person;
    }

    pub fn deinit(self: *Person) void{
        self.allocator.free(self.name);
        self.allocator.destroy(self);
    }

};

pub fn main() !void {
    const allocator = &std.heap.page_allocator;
    const person = try Person.Init("rami", allocator);
    const person3 = Person {
        .allocator = allocator,
        ._name = "test"
    };
    const privateKey = try crypto.Crypto.PrivateKey.init(allocator);
    const pubKey = privateKey.derivePublicKey(allocator);

    std.debug.print("All your {s} are belong to us.\n", .{person._name});
    std.debug.print("All your {s} are belong to us.\n", .{person3._name});
    std.debug.print("Public key {s} are belong to us.\n", .{pubKey.bytes()});
}

test "simple test" {
    var list = std.ArrayList(i32).init(std.testing.allocator);
    defer list.deinit(); // try commenting this out and see if zig detects the memory leak!
    try list.append(42);
    try std.testing.expectEqual(@as(i32, 42), list.pop());
}
