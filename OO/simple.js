(function() {
  function Component(name) {
    this.name = name
  }

  function Bicycle(brand, model, size, components) {
    this.Brand = brand
    this.Model = model
    this.Size  = size
    this.Components = components
  }

  var myComponents = [
    new Component("frame"),
    new Component("forks"),
    new Component("shock"),
    new Component("wheels")
  ]

  Component.prototype.getName = function() {
    return this.name
  }

  Bicycle.prototype.Components = Component.prototype


  var bike = new Bicycle("Scott", "YPZ2", 15.5, myComponents)

  console.log(bike.Components[0].getName())
})()
