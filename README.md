# cryptopal-go

Command line utility to simplify key management.

## Motivation

Cryptocurrency at the time of this writing has one thing that I dislike: Once the private key is leaked, there is no way to know it's leaked until everything associated with that key is gone. Even if we know it's leaked, the only way to solve the problem is moving associated stuffs to another key. Generating vanity key using online tools sometimes pose a risk because the source code are minified, and performance is not their advantage. Performing key management online are even more risky. Therefore I decided to make a tool using Go to utilize my computer capability better.

## License

CryptoPal is licensed under MIT license. See LICENSE file and NOTICE file for more details.
