(function() {

  function Animal() {
    this.says = 'zzzz'
  }

  Animal.prototype.speak = function Speak() {
    return this.says
  }

  function Cat() {
    this.says = 'Meow'
  }

  Cat.prototype = Animal.prototype

  var a_cat  = new Cat
  var animal = new Animal

  console.log(a_cat)
  console.log(animal)

  console.log(a_cat.Speak()) // Meow
  console.log(animal.Speak()) // zzzz

  console.log('Is Cat an Animal?', typeof Cat === 'Animal')

})()