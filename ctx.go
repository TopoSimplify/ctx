package ctx

import (
	"github.com/intdxdt/cmp"
	"github.com/intdxdt/geom"
	"github.com/intdxdt/mbr"
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


type ContextGeoms []*CtxGeom

func (s ContextGeoms) Len() int {
	return len(s)
}

func (s ContextGeoms) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ContextGeoms) Less(i, j int) bool {
	return s[i].I < s[j].I
}


type ctxMeta struct {
	SelfVertices    *sset.SSet
	SelfNonVertices *sset.SSet
}

type CtxGeom struct {
	Geom geom.Geometry
	Type string
	I    int
	J    int
	Meta *ctxMeta
}

func NewCtxGeom(g geom.Geometry, i, j int) *CtxGeom {
	return &CtxGeom{
		Geom: g,
		Type: Self,
		I:    i,
		J:    j,
		Meta: &ctxMeta{
			SelfVertices:    sset.NewSSet(cmp.Int, 2),
			SelfNonVertices: sset.NewSSet(cmp.Int, 2),
		},
	}
}

//implements stringer interface
func (o *CtxGeom) String() string {
	return o.Geom.WKT()
}

//implements igeom interface
func (o *CtxGeom) Geometry() geom.Geometry{
	return o.Geom
}

//implements bbox interface
func (o *CtxGeom) BBox() *mbr.MBR {
	return o.Geom.BBox()
}

// --------------------------------------------------------------------
func (o *CtxGeom) AsSelf() *CtxGeom {
	o.Type = Self
	return o
}

func (o *CtxGeom) IsSelf() bool {
	return o.Type == Self
}

// --------------------------------------------------------------------
func (o *CtxGeom) AsSelfVertex() *CtxGeom {
	o.Type = SelfVertex
	return o
}

func (o *CtxGeom) IsSelfVertex() bool {
	return o.Type == SelfVertex
}

// --------------------------------------------------------------------
func (o *CtxGeom) AsSelfNonVertex() *CtxGeom {
	o.Type = SelfNonVertex
	return o
}

func (o *CtxGeom) IsSelfNonVertex() bool {
	return o.Type == SelfNonVertex
}

// --------------------------------------------------------------------
func (o *CtxGeom) AsSelfSegment() *CtxGeom {
	o.Type = SelfSegment
	return o
}

func (o *CtxGeom) IsSelfSegment() bool {
	return o.Type == SelfSegment
}

// --------------------------------------------------------------------
func (o *CtxGeom) AsSelfSimple() *CtxGeom {
	o.Type = SelfSimple
	return o
}

func (o *CtxGeom) IsSelfSimple() bool {
	return o.Type == SelfSimple
}

// --------------------------------------------------------------------
func (o *CtxGeom) AsContextNeighbour() *CtxGeom {
	o.Type = ContextNeighbour
	return o
}

func (o *CtxGeom) IsContextNeighbour() bool {
	return o.Type == ContextNeighbour
}

// --------------------------------------------------------------------
func (o *CtxGeom) Intersection(other geom.Geometry) []*geom.Point {
	return o.Geom.Intersection(other)
}

