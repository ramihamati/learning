const std = @import("std");
const address = @import("./core/address.zig");

const Employee = struct {
    name: []const u8,
};

const Leader = struct {
    name: []const u8,
};

const Human = union(enum) {
    employee: Employee,
    leader: Leader,
};

pub fn GetName (self: Human) []const u8 {
    switch (self) {
        .employee => |employee| return employee.name,
        .leader => |leader| return leader.name
    }
}
pub fn main() !void {
    const a1 = address.Address.new();

    std.debug.print("address is {x}\n", .{a1.Value()});
    std.debug.print("address is {}\n", .{a1.Value()[0]});
    try await Huh();
    //
    // const p1 = Leader{
    //     .name = "ramis"
    // };
    //
    // const human = Human{.leader = p1};
    //
    // const p2 = GetName(human);
    //
    // try stdout.print("person name {s}\n", .{p2} );
    // try stdout.print("Run `zig build test` to run the tests.\n", .{});
}

pub fn Huh() !void {
    std.log.debug("hey {s}", .{"chief"});
}

test "simple test" {
    var list = std.ArrayList(i32).init(std.testing.allocator);
    defer list.deinit(); // try commenting this out and see if zig detects the memory leak!
    try list.append(42);
    try std.testing.expectEqual(@as(i32, 42), list.pop());
}
