module crypto

import crypto.ecdsa

@[noinit]
pub struct PublicKey {
	key ecdsa.PublicKey
}

pub fn PublicKey.new(key ecdsa.PublicKey) PublicKey {
	return PublicKey{
		key: key
	}
}

@[noinit]
pub struct PrivateKey {
	private_key ecdsa.PrivateKey
	public_key  ecdsa.PublicKey
}

pub fn PrivateKey.new() !PrivateKey {
	public_key, private_key := ecdsa.generate_key()!
	return PrivateKey{
		private_key: private_key
		public_key:  public_key
	}
}

pub fn (key &PrivateKey) public_key() PublicKey {
	return PublicKey.new(key.public_key)
}
