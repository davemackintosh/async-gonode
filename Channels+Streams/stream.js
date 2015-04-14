var stream = require('stream')
var util = require('util')
var Transform = stream.Transform
var Streamer = function(options) {
  Transform.call(this, options)
  this.text = []
}

util.inherits(Streamer, Transform)

Streamer.prototype._transform = function(chunk, encoding, callback) {
  this.text.push(chunk.toString('utf8'))
  callback()
}

Streamer.prototype._flush = function(callback) {
  this.text.forEach(function(text) {
    this.push(text, 'utf8')
  }.bind(this))
  callback()
}

//Here's how we use this stream
var outputStreamer = new Streamer()

outputStreamer.on('data', function(data) {
  console.log('out: ', data.toString('utf8'))
})

outputStreamer.write('Hello ')
outputStreamer.write('World ')
outputStreamer.end()
