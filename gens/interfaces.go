package gens

import "github.com/EdsonHTJ/jtos/domain"

type Generator interface {
	InsertObject(packageName string, object domain.Object) error
	Generate(packageName string) string
	GetOutPath(packageName string) string
}
