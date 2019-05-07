package mockbuilder

import "io"

type BodyComparer interface {
	Match(io.Reader, string) (bool, error)
}
