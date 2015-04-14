var stream = require('stream'),
  util = require('util')
var Transform = stream.Transform
var Reverse = function(options) {
  Transform.call(this, options)
  this.text = []
}
util.inherits(Reverse, Transform)
Reverse.prototype._transform = function(chunk, encoding, callback) {
  this.text.push(chunk.toString('utf8'))
  callback()
}
Reverse.prototype._flush = function(callback) {
  this.text.forEach(function(text) {
    this.push(text, 'utf8')
  }.bind(this))
  callback()
}

//Here's how we use this stream
var reverseStream = new Reverse()
reverseStream.on('data', function(data) {
    console.log('out: ', data.toString('utf8'))
})

reverseStream.write('Hello ')
reverseStream.write('World ')
reverseStream.end()
