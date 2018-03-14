// Copyright 2014 The Go-NetCDF Authors. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// These files are autogenerated from nc_double.go using generate.go
// DO NOT EDIT (except nc_double.go).

package netcdf

import (
	"fmt"
	"unsafe"
)

// #include <stdlib.h>
// #include <netcdf.h>
import "C"

// WriteUint32s writes data as the entire data for variable v.
func (v Var) WriteUint32s(data []uint32) error {
	if err := okData(v, UINT, len(data)); err != nil {
		return err
	}
	return newError(C.nc_put_var_uint(C.int(v.ds), C.int(v.id), (*C.uint)(unsafe.Pointer(&data[0]))))
}

// ReadUint32s reads the entire variable v into data, which must have enough
// space for all the values (i.e. len(data) must be at least v.Len()).
func (v Var) ReadUint32s(data []uint32) error {
	if err := okData(v, UINT, len(data)); err != nil {
		return err
	}
	return newError(C.nc_get_var_uint(C.int(v.ds), C.int(v.id), (*C.uint)(unsafe.Pointer(&data[0]))))
}

// WriteUint32s sets the value of attribute a to val.
func (a Attr) WriteUint32s(val []uint32) error {
	// We don't need okData here because netcdf library doesn't know
	// the length or type of the attribute yet.
	cname := C.CString(a.name)
	defer C.free(unsafe.Pointer(cname))
	return newError(C.nc_put_att_uint(C.int(a.v.ds), C.int(a.v.id), cname,
		C.nc_type(UINT), C.size_t(len(val)), (*C.uint)(unsafe.Pointer(&val[0]))))
}

// ReadUint32s reads the entire attribute value into val.
func (a Attr) ReadUint32s(val []uint32) (err error) {
	if err := okData(a, UINT, len(val)); err != nil {
		return err
	}
	cname := C.CString(a.name)
	defer C.free(unsafe.Pointer(cname))
	err = newError(C.nc_get_att_uint(C.int(a.v.ds), C.int(a.v.id), cname,
		(*C.uint)(unsafe.Pointer(&val[0]))))
	return
}

// ReadUint32At returns a value via index position
func (v Var) ReadUint32At(idx []uint64) (val uint32, err error) {
	err = newError(C.nc_get_var1_uint(C.int(v.ds), C.int(v.id),
		(*C.size_t)(unsafe.Pointer(&idx[0])), (*C.uint)(unsafe.Pointer(&val))))
	return
}

// WriteUint32At sets a value via its index position
func (v Var) WriteUint32At(idx []uint64, val uint32) (err error) {
	err = newError(C.nc_put_var1_uint(C.int(v.ds), C.int(v.id),
		(*C.size_t)(unsafe.Pointer(&idx[0])), (*C.uint)(unsafe.Pointer(&val))))
	return
}

// Uint32sReader is a interface that allows reading a sequence of values of fixed length.
type Uint32sReader interface {
	Len() (n uint64, err error)
	ReadUint32s(val []uint32) (err error)
}

// GetUint32s reads the entire data in r and returns it.
func GetUint32s(r Uint32sReader) (data []uint32, err error) {
	n, err := r.Len()
	if err != nil {
		return
	}
	data = make([]uint32, n)
	err = r.ReadUint32s(data)
	return
}

// testReadUint32s writes somes data to v. N is v.Len().
// This function is only used for testing.
func testWriteUint32s(v Var, n uint64) error {
	data := make([]uint32, n)
	for i := 0; i < int(n); i++ {
		data[i] = uint32(i + 10)
	}
	return v.WriteUint32s(data)
}

// testReadUint32s reads data from v and checks that it's the same as what
// was written by testWriteDouble. N is v.Len().
// This function is only used for testing.
func testReadUint32s(v Var, n uint64) error {
	data := make([]uint32, n)
	if err := v.ReadUint32s(data); err != nil {
		return err
	}
	for i := 0; i < int(n); i++ {
		if val := uint32(i + 10); data[i] != val {
			return fmt.Errorf("data at position %d is %v; expected %v", i, data[i], val)
		}
	}
	return nil
}

func testReadUint32At(v Var, n uint64) error {
	data := make([]uint32, n)
	if err := v.ReadUint32s(data); err != nil {
		return err
	}
	for i := 0; i < int(n); i++ {
		shape, _ := v.LenDims()
		coords, _ := UnravelIndex(uint64(i), shape)
		expected := uint32(i + 10)
		val, _ := v.ReadUint32At(coords)
		if val != data[i] {
			return fmt.Errorf("data at position %v is %v; expected %v", i, val, expected)
		}
	}
	return nil
}

func testWriteUint32At(v Var, n uint64) error {
	shape, _ := v.LenDims()
	ndim := len(shape)
	coord := make([]uint64, ndim)
	for i := 0; i < ndim; i++ {
		for k := 0; k < ndim; k++ {
			coord[k] = uint64(i)
		}
		v.WriteUint32At(coord, uint32(i))
	}
	for i := 0; i < ndim; i++ {
		for k := 0; k < ndim; k++ {
			coord[k] = uint64(i)
		}
		val, _ := v.ReadUint32At(coord)
		if val != uint32(i) {
			return fmt.Errorf("data at position %v is %v; expected %v", coord, val, int(i))
		}
	}
	return nil
}
