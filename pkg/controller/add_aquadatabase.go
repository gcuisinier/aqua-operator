package controller

import (
	"github.com/aquasecurity/aqua-operator/pkg/controller/aquadatabase"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, aquadatabase.Add)
}
