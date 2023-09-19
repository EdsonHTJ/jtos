package gens

import "github.com/EdsonHTJ/jtos/domain"

type Generator interface {
	ParseObject(packageName string, object domain.Object) string
	Generate(packageName string) string
	GetOutPath(packageName string) string
}
