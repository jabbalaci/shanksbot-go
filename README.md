# ShanksBot-Go

In the YouTube video [The Reciprocals of Primes](https://www.youtube.com/watch?v=DmfxIhmGPP4), Matt Parker explores the work of William Shanks, who was an English mathematician around 1870-1880. In the video, Matt shows ShanksBot, his Python program.

This is a Go implementation of ShanksBot.

The repository contains two binary files:

## cmd/shanksbot

    $ go run ./cmd/shanksbot

This is an interactive program that asks you to
enter a positive prime number.

It calculates the number of periods of the reciprocal
of the given prime number.

(Note that it also works with non-prime numbers.)

## cmd/generator

    $ go run ./cmd/generator

[William Shanks](https://en.wikipedia.org/wiki/William_Shanks) published a table of primes (and the periods of their reciprocals) up to 110,000.

This program prints this table to the screen. In the
first column we have a prime number, and next to it
you can see the number of periods of the reciprocal
of the given prime number.

In the folder `cmd/generator/`, you can also find
this table in a text file.

## Links
* [The Reciprocals of Primes](https://www.youtube.com/watch?v=DmfxIhmGPP4), on YouTube
* [William Shanks](https://en.wikipedia.org/wiki/William_Shanks), an English mathematician
* [ShanksBot-RS](https://github.com/babymotte/shanksbot-rs), a Rust implementation of ShanksBot (related work)
