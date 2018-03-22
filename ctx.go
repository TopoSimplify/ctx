package ctx

import (
	"github.com/intdxdt/mbr"
	"github.com/intdxdt/cmp"
	"github.com/intdxdt/geom"
	"github.com/intdxdt/sset"
)

const (
	Self             = "self"
	PlanarVertex     = "planar_vertex"
	NonPlanarVertex  = "non_planar_vertex"
	PlanarSegment    = "planar_segment"
	LinearSimple     = "linear_simple"
	ContextNeighbour = "context_neighbour"
)

type Meta struct {
	PlanarVertices    *sset.SSet
	NonPlanarVertices *sset.SSet
}

type ContextGeometry struct {
	Geom geom.Geometry
	Type string
	I    int
	J    int
	Meta *Meta
}

func New(g geom.Geometry, i, j int) *ContextGeometry {
	return &ContextGeometry{
		Geom: g,
		Type: Self,
		I:    i,
		J:    j,
		Meta: &Meta{
			PlanarVertices:    sset.NewSSet(cmp.Int, 2),
			NonPlanarVertices: sset.NewSSet(cmp.Int, 2),
		},
	}
}

//implements stringer interface
func (o *ContextGeometry) String() string {
	return o.Geom.WKT()
}

//implements igeom interface
func (o *ContextGeometry) Geometry() geom.Geometry {
	return o.Geom
}

//implements bbox interface
func (o *ContextGeometry) BBox() *mbr.MBR {
	return o.Geom.BBox()
}

//--------------------------------------------------------------------
func (o *ContextGeometry) AsSelf() *ContextGeometry {
	o.Type = Self
	return o
}

func (o *ContextGeometry) IsSelf() bool {
	return o.Type == Self
}

//--------------------------------------------------------------------
func (o *ContextGeometry) AsPlanarVertex() *ContextGeometry {
	o.Type = PlanarVertex
	return o
}

func (o *ContextGeometry) IsPlanarVertex() bool {
	return o.Type == PlanarVertex
}

//--------------------------------------------------------------------
func (o *ContextGeometry) AsNonPlanarVertex() *ContextGeometry {
	o.Type = NonPlanarVertex
	return o
}

func (o *ContextGeometry) IsNonPlanarVertex() bool {
	return o.Type == NonPlanarVertex
}

//--------------------------------------------------------------------
func (o *ContextGeometry) AsPlanarSegment() *ContextGeometry {
	o.Type = PlanarSegment
	return o
}

func (o *ContextGeometry) IsPlanarSegment() bool {
	return o.Type == PlanarSegment
}

//--------------------------------------------------------------------
func (o *ContextGeometry) AsLinearSimple() *ContextGeometry {
	o.Type = LinearSimple
	return o
}

func (o *ContextGeometry) IsLinearSimple() bool {
	return o.Type == LinearSimple
}

//--------------------------------------------------------------------
func (o *ContextGeometry) AsContextGeometries(objects ...*ContextGeometry) *ContextGeometries {
	return NewContexts(o).Extend(objects)
}

func (o *ContextGeometry) AsContextNeighbour() *ContextGeometry {
	o.Type = ContextNeighbour
	return o
}

func (o *ContextGeometry) IsContextNeighbour() bool {
	return o.Type == ContextNeighbour
}

//--------------------------------------------------------------------
func (o *ContextGeometry) Intersection(other geom.Geometry) []*geom.Point {
	return o.Geom.Intersection(other)
}
