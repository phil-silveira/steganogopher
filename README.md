# SteganoGopher

 <img src="examples/gopher-mustache.png" width="150" align="right"/>

![Go](https://img.shields.io/badge/Golang-gray?logo=go&labelColor=EEE)
[![LICENSE MIT](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://img.shields.io/badge/license-MIT-brightgreen.svg)
![Code Review](https://github.com/phil-silveira/steganogopher/actions/workflows/code-review.yml/badge.svg)


**SteganoGopher** is a steganography CLI written in pure go made for Gophers.
You can use this tool to encode/decode text messages into/from .PNG files.

## Installation

To properly install this tool, you need to clone this repository and run the command below:
```
    make install
```

## Demostration

| Original | Encoded |
| -------- | ------------ |
|<img src="examples/le-petit-prince.png" width="150" height="200" /> | <img src="examples/le-petit-prince.out.png" width="150" height="200" /> |
| no encoded data | encoded data: `"It is the time you have wasted for your rose that makes your rose so important."` |

### Encode message into a file
```
steganogopher encode screenshot.png -m "type your message right here"
```
encode should create a file `screenshot.out.png`

### Decode hidden messge from a file
```
steganogopher decode screenshot.out.png 
```
the message decoded should be stored inside file `decoded-message.txt`

### Get image infos
```
steganogopher info screenshot.png
```
should display info about the image `screenshot.out.png`
