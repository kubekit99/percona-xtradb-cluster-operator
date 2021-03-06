package controller

import (
	"github.com/Percona-Lab/percona-xtradb-cluster-operator/pkg/controller/perconaxtradbcluster"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, perconaxtradbcluster.Add)
}
