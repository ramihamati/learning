const std = @import("std");

pub fn main() !void {
    var biglist = std.ArrayList(u8).init(std.heap.page_allocator);
    defer biglist.deinit();

    var index: u32 = 1;
    while (index < 101) : (index + 1) {
        if (index % 10 == 7 or index % 7 == 0) {
            try biglist.writer().print("SMAC", .{});
        } else {
            try biglist.writer().print("NO SMAC", .{});
        }
    }

    try print_list(&biglist);
}

pub fn print_list(list: *std.ArrayList(u8)){

}
