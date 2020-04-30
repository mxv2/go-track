package space

// EarthAgeSeconds is Earth orbital period in seconds
const EarthAgeSeconds = 31557600

// Planet is string representation of Sun system planets.
type Planet string

// Age returns someone age on other Sum system planet.
func Age(seconds float64, planet Planet) float64 {
	earthAge := earthAgeOnPlanet(planet)
	return seconds / (earthAge * EarthAgeSeconds)
}

func earthAgeOnPlanet(planet Planet) float64 {
	switch planet {
	case "Mercury":
		return 0.2408467
	case "Venus":
		return 0.61519726
	case "Mars":
		return 1.8808158
	case "Jupiter":
		return 11.862615
	case "Saturn":
		return 29.447498
	case "Uranus":
		return 84.016846
	case "Neptune":
		return 164.79132
	default:
		return 1.0
	}
}
