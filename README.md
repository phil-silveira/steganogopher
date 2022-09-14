# SteganoGopher

 <img src="examples/gopher-mustache.png" width="150" align="right"/>

![Go](https://img.shields.io/badge/Golang-gray?logo=go&labelColor=EEE)
[![LICENSE MIT](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://img.shields.io/badge/license-MIT-brightgreen.svg)
![Code Review](https://github.com/phil-silveira/steganogopher/actions/workflows/code-review.yml/badge.svg)


SteganoGopher is a steganography CLI written in pure go made for Gophers.
You can use this tool to encode/decode text messages into/from .PNG files.

## Installation

To proberly install this tool, you may clone this repository and run the command below:
```
    make install
```

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

### Decode hidden messge from a file
```
steganogopher -d screenshot.out.png 
```
decode result should be something like `{"message": "type your message right here"}`
