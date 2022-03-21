package registry

import "fmt"

func RepositoryNameInvalidError(name string) error {
	return fmt.Errorf("invalid repository name: %v", name)
}

func RepositoryNameNotFoundError(name string) error {
	return fmt.Errorf("unable to find a repository with the name: %v", name)
}
