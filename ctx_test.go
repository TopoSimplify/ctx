package ctx

import (
	"testing"
	"github.com/intdxdt/geom"
	"github.com/franela/goblin"
)

func TestCtx(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("context neighbours", func() {
		g.It("should test context neighbours", func() {
	        lnGeom := geom.NewLineString([]*geom.Point{{0,0}, {5, 5}})
	        ctxGeom := geom.NewPointXY(2.5, 2.5)
			ctxG := New(ctxGeom, 0 , -1)

			inters := ctxG.Intersection(lnGeom)
			g.Assert(len(inters)==1).IsTrue()
			g.Assert(inters[0].Equals2D(ctxGeom)).IsTrue()
	        g.Assert(ctxG.BBox().IsPoint()).IsTrue()
	        g.Assert(ctxG.String() == ctxGeom.WKT()).IsTrue()

	        g.Assert(ctxG.IsSelf()).IsTrue()
	        g.Assert(ctxG.AsSelf().IsSelf()).IsTrue()
	        g.Assert(ctxG.AsSelfVertex().IsSelfVertex())
	        g.Assert(ctxG.AsSelfNonVertex().IsSelfNonVertex())
	        g.Assert(ctxG.AsSelfSegment().IsSelfSegment())
	        g.Assert(ctxG.AsSelfSimple().IsSelfSimple())
	        g.Assert(ctxG.AsSelfNonVertex().IsSelfNonVertex())
	        g.Assert(ctxG.AsContextNeighbour().IsContextNeighbour())
		})
	})
}
