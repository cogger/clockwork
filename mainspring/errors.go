package mainspring

import "errors"

var ErrCircularDependency = errors.New("circulare dependency detected")
