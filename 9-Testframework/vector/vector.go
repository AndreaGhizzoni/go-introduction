//Basic Vector module
package vector

type Vector []interface{}

//create a new instance of Vector returning the pointer
func New() *Vector {
	return new(Vector)
}

//return the length of the Vector
func (p *Vector) Len() int {
	return len(*p)
}

//add a element
func (p *Vector) Add(elem interface{}) {
	(*p) = append((*p), elem)
}

//add a element at given position
func (p *Vector) AddAt(elem interface{}, i int) {
	if i < 0 || i > len(*p) {
		return
	}
	(*p)[i] = elem
}

//delete the last element
func (p *Vector) DeleteLast() {
	(*p) = (*p)[:len(*p)]
}

//delete the first element
func (p *Vector) DeleteFirst() {
	(*p) = (*p)[1:len(*p)]
}

//return the i'th element
func (p *Vector) At(i int) interface{} {
	if i < 0 || i > len(*p) {
		return nil
	}
	return (*p)[i]
}

//check if element is present
func (p *Vector) Contains(elem interface{}) bool {
	for _, x := range *p {
		if x == elem {
			return true
		}
	}
	return false
}
