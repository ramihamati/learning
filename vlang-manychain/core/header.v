module core

import hash
import io
import encoding.binary

struct Header {
pub:
	version   u32
	prev_hash hash.Hash
	timestamp u64
	height    u64
	no_once   u64
}
