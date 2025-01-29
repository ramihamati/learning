const std = @import("std");
const cr = std.crypto;

pub const Address = struct {
    value : *const [65]u8,

    pub fn new() Address {
        const kp = cr.sign.ecdsa.EcdsaP256Sha256
        .KeyPair.generate();

        const bytes   = kp.public_key.toUncompressedSec1();

        return Address {
            .value = &bytes
        };
    }

    pub fn Value(self: Address) *const [65]u8   {
        return self.value;
    }
};