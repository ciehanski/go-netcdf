// Copyright 2014 The Go-NetCDF Authors. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// These files are autogenerated from nc_double.go using generate.sh

package netcdf

import (
	"unsafe"
)

// #include <stdlib.h>
// #include <netcdf.h>
import "C"

// WriteByte writes data as the entire data for variable v.
func (v Var) WriteByte(data []int8) error {
	if err := okData(v, NC_BYTE, len(data)); err != nil {
		return err
	}
	return newError(C.nc_put_var_schar(C.int(v.f), C.int(v.id), (*C.schar)(unsafe.Pointer(&data[0]))))
}

// ReadByte reads the entire variable v into data, which must have enough
// space for all the values (i.e. len(data) must be at least v.Len()).
func (v Var) ReadByte(data []int8) error {
	if err := okData(v, NC_BYTE, len(data)); err != nil {
		return err
	}
	return newError(C.nc_get_var_schar(C.int(v.f), C.int(v.id), (*C.schar)(unsafe.Pointer(&data[0]))))
}

// WriteByte sets the value of attribute a to val.
func (a Attr) WriteByte(val []int8) error {
	// We don't need okData here because netcdf library doesn't know
	// the length or type of the attribute yet.
	cname := C.CString(a.name)
	defer C.free(unsafe.Pointer(cname))
	return newError(C.nc_put_att_schar(C.int(a.v.f), C.int(a.v.id), cname,
		C.nc_type(NC_BYTE), C.size_t(len(val)), (*C.schar)(unsafe.Pointer(&val[0]))))
}

// ReadByte reads the entire attribute value into val.
func (a Attr) ReadByte(val []int8) (err error) {
	if err := okData(a, NC_BYTE, len(val)); err != nil {
		return err
	}
	cname := C.CString(a.name)
	defer C.free(unsafe.Pointer(cname))
	err = newError(C.nc_get_att_schar(C.int(a.v.f), C.int(a.v.id), cname,
		(*C.schar)(unsafe.Pointer(&val[0]))))
	return
}

// ByteReader is a interface that allows reading a sequence of values of fixed length.
type ByteReader interface {
	Len() (n uint64, err error)
	ReadByte(val []int8) (err error)
}

// GetByte reads the entire data in r and returns it.
func GetByte(r ByteReader) (data []int8, err error) {
	n, err := r.Len()
	if err != nil {
		return
	}
	data = make([]int8, n)
	err = r.ReadByte(data)
	return
}
