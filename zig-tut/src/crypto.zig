const std = @import("std");
const crypto = @import("./crypto.zig");

pub const Crypto = struct {
    // Represents a public key.
    pub const PublicKey = struct {
        key: ?*std.crypto.ecc.PublicKey,

        // Serialize the public key to bytes.
        pub fn bytes(self: PublicKey, allocator: *const std.mem.Allocator) ![]u8 {
            const key = self.key orelse return error.InvalidKey;
            return key.serializeCompressed(allocator);
        }
    };

    // Represents a private key.
    pub const PrivateKey = struct {
        key: ?*std.crypto.ecc.PrivateKey,

        // Create a new private key.
        pub fn init(allocator: *const std.mem.Allocator) !PrivateKey {
            std.crypto.sign.ecdsa.EcdsaP256Sha256.KeyPair;
            const key = try std.crypto.ecc.P256.generate(allocator, .{ .curve = std.crypto.ecc.Curve.p256 });
            return PrivateKey{ .key = key };
        }

        // Derive the corresponding public key.
        pub fn derivePublicKey(self: PrivateKey, allocator: *const std.mem.Allocator) !PublicKey {
            const privKey = self.key orelse return error.InvalidKey;
            const pubKey = try privKey.derivePublicKey(allocator);
            return PublicKey{ .key = pubKey };
        }
    };
};
