package snapshot

import "io"

type DatabaseSnapshot interface {
	Load(r io.Reader) (map[string]any, error)
	Save(w io.Writer) error
}
