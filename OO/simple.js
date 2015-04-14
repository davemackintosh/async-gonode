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
  var animal = new Animal()

  console.log(a_cat)
  console.log(animal)

  console.log(a_cat.Speak())

})()
