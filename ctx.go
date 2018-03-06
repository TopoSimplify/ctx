package ctx

import (
	"github.com/intdxdt/mbr"
	"github.com/intdxdt/cmp"
	"github.com/intdxdt/geom"
	"github.com/intdxdt/sset"
)

const (
	Self             = "self"
	SelfVertex       = "self_vertex"
	SelfSegment      = "self_segment"
	SelfSimple       = "self_simple"
	SelfNonVertex    = "self_non_vertex"
	ContextNeighbour = "context_neighbour"
)

type Meta struct {
	SelfVertices    *sset.SSet
	SelfNonVertices *sset.SSet
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
			SelfVertices:    sset.NewSSet(cmp.Int, 2),
			SelfNonVertices: sset.NewSSet(cmp.Int, 2),
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
func (o *ContextGeometry) AsSelfVertex() *ContextGeometry {
	o.Type = SelfVertex
	return o
}

func (o *ContextGeometry) IsSelfVertex() bool {
	return o.Type == SelfVertex
}

//--------------------------------------------------------------------
func (o *ContextGeometry) AsSelfNonVertex() *ContextGeometry {
	o.Type = SelfNonVertex
	return o
}

func (o *ContextGeometry) IsSelfNonVertex() bool {
	return o.Type == SelfNonVertex
}

//--------------------------------------------------------------------
func (o *ContextGeometry) AsSelfSegment() *ContextGeometry {
	o.Type = SelfSegment
	return o
}

func (o *ContextGeometry) IsSelfSegment() bool {
	return o.Type == SelfSegment
}

//--------------------------------------------------------------------
func (o *ContextGeometry) AsSelfSimple() *ContextGeometry {
	o.Type = SelfSimple
	return o
}

func (o *ContextGeometry) IsSelfSimple() bool {
	return o.Type == SelfSimple
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
