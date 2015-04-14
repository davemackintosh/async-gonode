(function() {

  function Animal() {
    this.says = 'Nothing...'
  }

  Animal.prototype.Speak = function Speak() {
    return this.says
  }

  function Cat() {
    this.says = 'Meow'
  }

  Cat.prototype = Animal.prototype

  var a_cat  = new Cat()

  console.log(a_cat.Speak())

})()
