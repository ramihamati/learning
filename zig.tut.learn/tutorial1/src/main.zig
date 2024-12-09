const std = @import("std");
const print = std.debug.print;
const assert = std.debug.assert;

const Context = struct {
    Allocator: std.mem.Allocator,

    pub fn Init(allocator:  std.mem.Allocator) Context{
        return Context { 
            .Allocator = allocator
        };
    }
};

const NetAddr = struct {
    Address : []const u8,   

    pub fn Init(addr: []const u8) NetAddr {
        return NetAddr{
            .Address = addr
        };
    }

    pub fn Equals(self: NetAddr, other : NetAddr) bool{
        return std.mem.eql(u8, self.Address, other.Address);
    }
};

const RPC = struct {
    Payload : []u8
};


const Peer = struct{
    ConnectedPeers: std.ArrayList(*Peer),
    Transport: Transport,
    
    pub fn Init(transport: Transport, ctx: Context) Peer {
        return Peer{
            .ConnectedPeers= std.ArrayList(*Peer).init(ctx.Allocator),
            .Transport= transport
        };
    }
};

const Connection = struct{
    Peer : *Peer,
    Transport : *Transport,

};

const LocalTransport = struct{
    pub fn Init() LocalTransport{
        return LocalTransport{

        };
    }
};

const Transport = union(enum){
    local : LocalTransport
};

pub fn main() !void {

    std.debug.print("\ntest\n", .{});
        
    const firstAddr = NetAddr.Init("first");
    const secondAddr = NetAddr.Init("first");

    const context = Context.Init(std.heap.page_allocator);
    const transport = Transport {.local = LocalTransport.Init()};
    const peer1 = Peer.Init(transport, context);

    print("{}", .{firstAddr.Equals(secondAddr)});
    print ("{s}", .{firstAddr.Address});
    print("\n", .{});
    print("{}", @TypeOf(peer1));

}

pub fn values() void{
    const sum : i32 = 1 + 1;
    print("\n{}", .{sum});
    // interesting evaluation in an object
    print("\n{}", .{true or false});
    const optional_value : ?[]const u8 = null;
    assert(optional_value == null);

    const error_union : anyerror!i32 = error.ArgNotFound;
    print("\nerror union type : {} value : {!}", .{@TypeOf(error_union), error_union});
}

 // var biglist = std.ArrayList(u8).init(std.heap.page_allocator);
 //    defer biglist.deinit();
 //
 //    var index: u32 = 1;
 //    while (index < 101) : (index + 1) {
 //        if (index % 10 == 7 or index % 7 == 0) {
 //            try biglist.writer().print("SMAC", .{});
 //        } else {
 //            try biglist.writer().print("NO SMAC", .{});
 //        }
 //    }
 //
 //    try print_list(&biglist);

