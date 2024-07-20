# eratosthenes

In this repo I'm trying to play around with Go and its concurrency features. I want to implement the Sieve of Eratosthenes, which is a famous and simple way of calculating prime numbers. In the first iteration I am doing this with a simple nested loop, but the plan is to use concurrency and channels to reduce the memory requirements and possibly speed up the process.

## Channels vs. Vector

Interestingly sieving the numbers with the use of channels is not faster. Actually, it is way faster to just use the vector version, which is a little bit surprising to me. The speed can be increased, when the channels are buffered, but the best buffer combination I could reach was still around 3 times slower than the vector version.