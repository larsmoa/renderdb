package geometry

import "github.com/ungerik/go3d/float64/vec3"

type Object interface {
	// ID returns an unique ID of the object
	ID() int64
	// Bounds returns the bounding box of the object.
	Bounds() *vec3.Box
	// GeometryData returns raw geometry data of the object
	GeometryData() []byte
	// Metadata returns arbitrary JSON-convertible metadata for
	// the object.
	Metadata() interface{}
}

type SimpleObject struct {
	id           int64
	bounds       *vec3.Box
	geometryData []byte
	metadata     interface{}
}

func (o *SimpleObject) ID() int64 {
	return o.id
}

func (o *SimpleObject) Bounds() *vec3.Box {
	return o.bounds
}

func (o *SimpleObject) GeometryData() []byte {
	return o.geometryData
}

func (o *SimpleObject) Metadata() interface{} {
	return o.metadata
}
