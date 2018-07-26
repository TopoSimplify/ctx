package ctx

import (
	"github.com/intdxdt/geom"
	"github.com/TopoSimplify/rng"
)

func linearCoords(wkt string) []geom.Point {
	return geom.NewLineStringFromWKT(wkt).Coordinates()
}

//create ctx geometries
func ctxGeoms(indxs [][]int, coords []geom.Point) []*ContextGeometry {
	var hulls []*ContextGeometry
	for _, o := range indxs {
		var r = rng.Range(o[0], o[1])
		var c = coords[r.I:r.J+1]
		var g = hullGeom(c)
		var cg = New(g, r.I, r.J)
		hulls = append(hulls, cg)
	}
	return hulls
}


//hull geom
func hullGeom(coords []geom.Point) geom.Geometry {
	var g geom.Geometry

	if len(coords) > 2 {
		g = geom.NewPolygon(coords)
	} else if len(coords) == 2 {
		g = geom.NewLineString(coords)
	} else {
		g = coords[0]
	}
	return g
}
