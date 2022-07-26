# SteganoGopher 

[![LICENSE MIT](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://img.shields.io/badge/license-MIT-brightgreen.svg)

SteganoGopher is a steganography CLI written in pure go make for Gophers.
You can use this tool to encode/decode text messages into/from .PNG files.

## Demostration

| Original | Edited Image |
| -------- | ------------ |
|<img src="examples/le-petit-prince.png" width="150" height="200" /> | <img src="examples/le-petit-prince.out.png" width="150" height="200" /> |

The second image contains the message `"It is the time you have wasted for your rose that makes your rose so important."`


### Encode message into a file
```
steganogopher -e screenshot.png -m "type your message right here"
```
encode should create a file `screenshot.out.png`
### Decode a file
```
steganogopher -d screenshot.out.png 
```
decode result should be something like `{"message": "type your message right here"}`
