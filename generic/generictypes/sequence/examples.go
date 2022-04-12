package sequence

// type (
// 	 Vector(t) []t
// )

// type constraints can be defined *inline*
func FirstElem2[S interface{ ~[]E }, E any](s S) E {
	return s[0]
}
