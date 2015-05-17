package spring

import "errors"

var ErrDuplicateName = errors.New("name already exists in the assessembly")

var ErrDoesNotExist = errors.New("spring by that name does not exist")

var ErrCircularDependency = errors.New("circulare dependency detected")
