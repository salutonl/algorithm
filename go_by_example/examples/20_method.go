package examples

type Rect struct {
	Width, Height int
}

func (r Rect) Area() int {
	return r.Height * r.Width 
}

func (rp *Rect) Perim() int {
	return 2 * (rp.Height + rp.Width)
}