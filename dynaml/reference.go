package dynaml

type ReferenceExpr struct {
	Path []string
}

func (e ReferenceExpr) Evaluate(context Context) Node {
	reference := context.FindReference(e.Path)

	switch reference.(type) {
	case Expression:
		return nil
	default:
		return reference
	}
}
