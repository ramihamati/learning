module network

@[noinit]
struct LocalStream implements IStream {
	channel chan RPC
mut:
	is_closed bool
}

pub fn (ls &LocalStream) send_message(message RPC) {
	if ls.is_closed {
		panic('the channel is closed')
	}
	ls.channel <- message
}

pub fn (ls &LocalStream) consume(consumer fn (message RPC)) {
	go fn [ls, consumer] () {
		for {
			if ls.is_closed {
				return
			}
			select {
				message := <-ls.channel {
					consumer(message)
				}
			}
		}
	}()
}

pub fn (mut ls LocalStream) close() {
	if ls.is_closed {
		return
	}
	ls.channel.close()
	ls.is_closed = true
}

pub fn LocalStream.new() LocalStream {
	channel := chan RPC{}

	return LocalStream{
		channel:   channel
		is_closed: false
	}
}
