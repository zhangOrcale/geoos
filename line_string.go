package geoos

// LineString represents a set of points to be thought of as a polyline.
type LineString []Point

// GeoJSONType returns the GeoJSON type for the linestring.
func (ls LineString) GeoJSONType() string {
	return TypeLineString
}

// Dimensions returns 1 because a LineString is a 1d object.
func (ls LineString) Dimensions() int {
	return 1
}

// Nums num of linestrings
func (ls LineString) Nums() int {
	return 1
}

// Boundary returns the closure of the combinatorial boundary of this linestring.
func (ls LineString) Boundary() (Geometry, error) {
	s := NormalStrategy()
	return s.Boundary(ls)
}
