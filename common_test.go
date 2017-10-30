package ctx

import (
	"simplex/rng"
	"github.com/intdxdt/geom"
)

//
//import (
//	"fmt"
//	"simplex/pln"
//	"simplex/rng"
//	"simplex/node"
//	"github.com/intdxdt/geom"
//	"github.com/intdxdt/rtree"
//)
//
//func DebugPrintNodes(ns []*node.Node) {
//	for _, n := range ns {
//		fmt.Println(n.Geom.WKT())
//	}
//}
//
//func ctxGeom(wkt string) *ContextGeometry {
//	return New(geom.NewGeometry(wkt), 0, -1)
//}
//
func linearCoords(wkt string) []*geom.Point {
	return geom.NewLineStringFromWKT(wkt).Coordinates()
}

//create ctx geometries
func ctxGeoms(indxs [][]int, coords []*geom.Point) []*ContextGeometry {

	var hulls = []*ContextGeometry{}
	for _, o := range indxs {
		var r = rng.NewRange(o[0], o[1])
		var c = coords[r.I():r.J()+1]
		var g = hullGeom(c)
		var cg = New(g, r.I(), r.J())
		hulls = append(hulls, cg)
	}
	return hulls
}

////hull db
//func hullsDB(ns []*node.Node) *rtree.RTree {
//	db := rtree.NewRTree(8)
//	for _, n := range ns {
//		db.Insert(n)
//	}
//	return db
//}
//
////hull geom
func hullGeom(coords []*geom.Point) geom.Geometry {
	var g geom.Geometry

	if len(coords) > 2 {
		g = geom.NewPolygon(coords)
	} else if len(coords) == 2 {
		g = geom.NewLineString(coords)
	} else {
		g = coords[0].Clone()
	}
	return g
}
