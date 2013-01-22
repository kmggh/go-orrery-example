// A Solar System Orrery in Go
// using object-oriented programming
// Author: kmguyton
// Thu 2011-03-24 17:40:05 -0400

package main

import "flag"
import "fmt"
import "math"

var (
	step_size  float64 = 0.5
	step_count int     = 5
)

func init() {
	flag.Float64Var(&step_size, "s", 0.5, "Size of each step in years.")
	flag.IntVar(&step_count, "c", 5, "Number of steps to compute.")
}

// All planets are initialized to the same zero-deg position.
// In other words, they are perfectly aligned to start.

var solarsystem_planets = []*Planet{
	newPlanet("Mercury", 0.39, 0.0),
	//      newPlanet("Venus", 0.72, 0.0),
	newPlanet("Earth", 1.0, 0.0),
	//      newPlanet("Mars", 1.52, 0.0),
	newPlanet("Jupiter", 5.2, 0.0),
	//      newPlanet("Saturn", 9.5, 0.0),
	//      newPlanet("Uranus", 19.2, 0.0),
	newPlanet("Neptune", 30.1, 0.0),
	//      newPlanet("Pluto", 39.5, 0.0),
}

type Planet struct {
	// A planet with a name, semi-major axis, and position.
	name string
	r    float64 // Semi-major axis in AU.
	pos  float64 // An orbit  position in radians.
}

func newPlanet(name string, r float64, pos float64) *Planet {
	// Create a new planet.
	return &Planet{name, r, pos}
}

func (planet *Planet) Period() (period float64) {
	// Return the planet’s period in years.
	// This is Kepler’s Third Law, the actual bit of astronomy in
	// this program.
	//
	// period^2 = r^3, period = r^(3/2)

	return math.Pow(planet.r, 1.5)
}

func (planet *Planet) Step(delta_t float64) float64 {
	// Advance the planet along the orbit for some time step (in years).
	var delta_theta float64 // The position step in degrees.

	delta_theta = (delta_t / planet.Period()) * 360.0

	planet.pos = planet.pos + delta_theta

	// Normalize the angle to the range (0, 359 deg).
	for planet.pos >= 360.0 {
		planet.pos -= 360.0
	}

	return planet.pos
}

func (planet *Planet) Str() string {
	return fmt.Sprintf("%s r=%.2f, period=%.2f, pos=%.2f deg",
		planet.name, planet.r, planet.Period(), planet.pos)
}

type Orrery struct {
	// An orrery with a list of planets and elapsed time in years.
	planets      []*Planet
	elapsed_time float64
}

func newOrrery(planets []*Planet) *Orrery {
	// Create a new orrery.
	var elapsed_time float64 = 0.0
	return &Orrery{planets, elapsed_time}
}

func (orrery *Orrery) Str() string {
	// Return a string for all planets in their current positions.
	var output_str string

	for i := 0; i < len(orrery.planets); i++ {
		output_str += orrery.planets[i].Str()
		output_str += "\n"
	}

	return output_str
}

func (orrery *Orrery) Step(delta_t float64) {
	// Step all planets in the orrery by delta_t years.

	orrery.elapsed_time += delta_t

	for i := 0; i < len(orrery.planets); i++ {
		orrery.planets[i].Step(delta_t)
	}
}

func main() {
	flag.Parse()
	fmt.Printf("\nSolar System Orrery\n\n")
	fmt.Printf("Stepping %.2f year for %d times.\n\n",
		step_size, step_count)

	var sol_sys = newOrrery(solarsystem_planets)
	fmt.Printf("%s\n", sol_sys.Str())

	for i := 0; i < step_count; i++ {
		sol_sys.Step(step_size)
		fmt.Printf("Elapsed time %.2f years\n\n",
			sol_sys.elapsed_time)
		fmt.Printf("%s\n", sol_sys.Str())
	}
}
