class Channel<T> {
	private queue: T[] = [];
	private resolvers: ((value: T) => void)[] = [];

	send(value: T) {
		console.log('Waiting for message...');
		console.log(this.queue);
		console.log(this.resolvers);

		if (this.resolvers.length > 0) {
			const resolve = this.resolvers.shift()!;
			resolve(value);
		} else {
			this.queue.push(value);
		}
	}

	receive(): Promise<T> {
		console.log('Waiting for messageSS SSS...');
		console.log(this.queue);
		console.log(this.resolvers);
		if (this.queue.length > 0) {
			return Promise.resolve(this.queue.shift()!);
		}
		return new Promise<T>((resolve) => {
			this.resolvers.push(resolve);
		});
	}
}

async function main() {
	const messages = new Channel<string>();
	// messages.receive();
	messages.send('hello world');
	messages.send('hello world 2');
	messages.send('ping!');

	// // Simulate goroutine
	// (async () => {

	// })();

	messages.receive();
	messages.receive();
	const msg = await messages.receive();
	console.log(msg);
	console.log(msg);
}

main();
