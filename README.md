# Go Concurrency Problem

## Thomas Carriero - 730510525

### Overview
This Go program implements a 3-stage concurrent pipeline using goroutines and channels. The pipeline consists of two producers generating numbers, two consumers squaring them, and a final filter ensuring the output is strictly increasing.

### Approach
1. Buffered Channels
Buffered channels with a capacity of 5 are used to store data between stages, allowing for asynchronous communication and preventing deadlocks by decoupling the stages.

2. Producers
I implemented 2 producer goroutines to generate odd and even numbers, respectively. They send these numbers into the inCh channel, introducing a random delay between sends to simulate real-world data generation.

3. Consumers
Two consumer goroutines receive numbers from inCh, square them, and send the results to the outCh channel. Each consumer introduces a random delay to simulate processing time, similar to the producers.

4. Final Filter
I had a single goroutine read from outCh and print each number only if it is strictly greater than the previously printed number, ensuring a strictly increasing sequence.

5. Deadlock Prevention
Deadlocks are prevented by ensuring that each channel has at least one sender and one receiver ready to operate. Buffered channels allow goroutines to send data without blocking, and this proper synchronization ensures that all data is processed without deadlocks.

### Questions for Discussion
1. Deadlocks are avoided by using buffered channels and ensuring that each stage has at least one sender and one receiver ready. The use of time.Sleep() introduces randomness, but the buffered channels provide enough capacity to prevent blocking, allowing the pipeline to continue processing data even if some stages are temporarily delayed.

2. In Elixir, I would use separate processes for each stage of the pipeline. Each producer would spawn a process that sends data to the next stage's process. Consumers would be separate processes that receive data from the previous stage and send the results to the next stage. This approach leverages Elixir's actor model, where each process has its own mailbox and communicates with other processes through message passing. To handle multiple producers sending data to multiple consumers, I would probably use a combination of GenServer processes (if allowed) and message routing techniques, ensuring that each message is processed by the appropriate consumer.