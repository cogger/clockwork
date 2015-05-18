package spring

import "errors"

//ErrDuplicateName is returned when the name is already in the assessembly
var ErrDuplicateName = errors.New("name already exists in the assessembly")

//ErrDoesNotExist is returned when you try to get a spring from the assessembly and it does not exist
var ErrDoesNotExist = errors.New("spring by that name does not exist")

//ErrCircularDependency is returned when a circular dependency is created
var ErrCircularDependency = errors.New("circulare dependency detected")
