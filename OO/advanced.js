(function() {
  function Vehicle() {}

  Vehicle.prototype = {
    wheels    : 1,
    type      : 'unicycle',
    fuel      : 100,
    consumes  : 0.75,
    distance  : 0,
    Travel : function Travel(distance) {
      if (this.fuel < 0) {
        console.log('Out of fuel!')
      } else {
        this.distance += distance * this.consumes
        this.fuel -= this.distance
      }

      return this.fuel
    }
  }

  function Bicycle() {
    this.wheels   = 2
    this.type     = 'bicycle'
    this.consumes = 0.5
    this.fuel     = 100
  }

  function Car() {
    this.wheels   = 4
    this.type     = 'car'
    this.consumes = 1.5
    this.fuel     = 100
  }

  Car.prototype = Bicycle.prototype = Vehicle.prototype

  var car = new Car()
  while (car.Travel(Math.random()) > 0)
    console.log('%s travelled %d, consuming %d', car.type, car.distance, car.consumes)

  var bike = new Bicycle()
  while (bike.Travel(Math.random()) > 0)
    console.log('%s travelled %d, consuming %d', bike.type, bike.distance, bike.consumes)
})()
