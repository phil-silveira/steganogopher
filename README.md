# SteganoGopher 

[![LICENSE MIT](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://img.shields.io/badge/license-MIT-brightgreen.svg)

SteganoGopher is a steganography CLI written in pure go make for Gophers.
You can use this tool to encode/decode text messages into/from .PNG files.

## Demostration
| Original | Edited Image |
| -------- | -------------|
|![original image](examples/avatar.png)| ![edited image](examples/avatar_out.png)|

The second image contains a hidden message ;-)



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
