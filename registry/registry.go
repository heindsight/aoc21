package registry

import "errors"

type Solution func()

var registry = map[string]Solution{}

func RegisterSolution(name string, soln Solution) error {
	_, exists := registry[name]
	if exists {
		return errors.New("Duplicate solution: " + name)
	}

	registry[name] = soln
	return nil
}

func GetSolution(name string) (Solution, error) {
	soln, found := registry[name]
	if !found {
		return nil, errors.New("Solution not found: " + name)
	}

	return soln, nil
}
