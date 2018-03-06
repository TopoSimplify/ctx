package ctx

func ConvertToContextsGeoms(objects []interface{}) []*ContextGeometry {
	var n = len(objects)
	var a = make([]*ContextGeometry, n, n)
	for i, o := range objects {
		a[i] = o.(*ContextGeometry)
	}
	return a
}
