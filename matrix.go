package gp

type Matrix struct {
	data map[int64]*Vector
}

func NewMatrix() *Matrix {
	m := Matrix{}
	m.data = make(map[int64]*Vector)
	return &m
}

func (m *Matrix) AddValue(k1, k2 int64, v float64){
	_, ok := m.data[k1]
	if !ok {
		m.data[k1] = NewVector()
	}
	m.data[k1].AddValue(k2, v)
}

func (m *Matrix) SetValue(k1, k2 int64, v float64){
	_, ok := m.data[k1]
	if !ok {
		m.data[k1] = NewVector()
	}
	m.data[k1].SetValue(k2, v)
}

func (m *Matrix) GetValue(k1, k2 int64) float64 {
	row := m.GetRow(k1)
	if row == nil {
		return 0.0
	} else {
		return row.GetValue(k2)
	}
}

func (m *Matrix) GetRow(k1 int64) *Vector {
	row, ok := m.data[k1]
	if !ok {
		return nil
	} else {
		return row
	}
	return row
}

func (m *Matrix) Scale(scale float64) *Matrix {
	ret := NewMatrix()
	for id, vi := range m.data {
		ret.data[id] = vi.Scale(scale)
	}
	return ret
}

func (m *Matrix) MultiplyVector(v *Vector) *Vector {
	// This is intended for l-by-m * m-by-1
	// For m-by-1 * 1-by-n, use OuterProduct in vector.go
	// Probably should just have a MatrixMultiply for everything
	ret := NewVector()
	for id, vi := range m.data {
		ret.SetValue(id, v.Dot(vi))
	}
	return ret
}

func (m *Matrix) Trans() *Matrix {
	ret := NewMatrix()
	for rid, vi := range m.data {
		for cid, w := range vi.data {
			ret.SetValue(cid, rid, w)
		}
	}
	return ret
}

func (m *Matrix) ElemWiseAddMatrix(n *Matrix) *Matrix {
	ret := NewMatrix()
	for key, mi := range m.data{
		ret.data[key] = mi
	}
	for key, ni := range n.data{
		if ret.GetRow(key) == nil{
			ret.data[key] = ni
		} else {
			ret.data[key] = ni.ElemWiseAddVector(ret.GetRow(key))
		}
	}
	return ret
}
