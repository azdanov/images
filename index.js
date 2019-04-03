#!/usr/bin/env node

const { existsSync, readFileSync } = require('fs')
const { extname } = require('path')

if (process.argv.length == 2) {
  process.stderr.write('No arguments specified.')
  process.exit(1)
}

const image = process.argv[2]

if (!existsSync(image)) {
  process.stderr.write('File not found.')
  process.exit(1)
}

const extension = extname(image).substr(1)
const supported = ['image/svg+xml', 'image/gif', 'image/jpeg', 'image/png']
const mime = supported.find(s => s.includes(extension))

if (!mime) {
  process.stderr.write('Image type not supported.')
  process.exit(1)
}

const data = readFileSync(image)
process.stdout.write(`data:${mime};base64,${data.toString('Base64')}`)
