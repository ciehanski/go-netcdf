// Copyright 2014 The Go-NetCDF Authors. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// These files are autogenerated from nc_double.go using generate.sh

package netcdf

import (
	"fmt"
	"unsafe"
)

// #include <stdlib.h>
// #include <netcdf.h>
import "C"

// WriteShort writes data as the entire data for variable v.
func (v Var) WriteShort(data []int16) error {
	if err := okData(v, NC_SHORT, len(data)); err != nil {
		return err
	}
	return newError(C.nc_put_var_short(C.int(v.f), C.int(v.id), (*C.short)(unsafe.Pointer(&data[0]))))
}

// ReadShort reads the entire variable v into data, which must have enough
// space for all the values (i.e. len(data) must be at least v.Len()).
func (v Var) ReadShort(data []int16) error {
	if err := okData(v, NC_SHORT, len(data)); err != nil {
		return err
	}
	return newError(C.nc_get_var_short(C.int(v.f), C.int(v.id), (*C.short)(unsafe.Pointer(&data[0]))))
}

// WriteShort sets the value of attribute a to val.
func (a Attr) WriteShort(val []int16) error {
	// We don't need okData here because netcdf library doesn't know
	// the length or type of the attribute yet.
	cname := C.CString(a.name)
	defer C.free(unsafe.Pointer(cname))
	return newError(C.nc_put_att_short(C.int(a.v.f), C.int(a.v.id), cname,
		C.nc_type(NC_SHORT), C.size_t(len(val)), (*C.short)(unsafe.Pointer(&val[0]))))
}

// ReadShort reads the entire attribute value into val.
func (a Attr) ReadShort(val []int16) (err error) {
	if err := okData(a, NC_SHORT, len(val)); err != nil {
		return err
	}
	cname := C.CString(a.name)
	defer C.free(unsafe.Pointer(cname))
	err = newError(C.nc_get_att_short(C.int(a.v.f), C.int(a.v.id), cname,
		(*C.short)(unsafe.Pointer(&val[0]))))
	return
}

// ShortReader is a interface that allows reading a sequence of values of fixed length.
type ShortReader interface {
	Len() (n uint64, err error)
	ReadShort(val []int16) (err error)
}

// GetShort reads the entire data in r and returns it.
func GetShort(r ShortReader) (data []int16, err error) {
	n, err := r.Len()
	if err != nil {
		return
	}
	data = make([]int16, n)
	err = r.ReadShort(data)
	return
}

// TestWriteShort writes somes data to v. N is v.Len().
// This function is only used for testing.
func testWriteShort(v Var, n uint64) error {
	data := make([]int16, n)
	for i := 0; i < int(n); i++ {
		data[i] = int16(i + 10)
	}
	return v.WriteShort(data)
}

// TestReadShort reads data from v and checks that it's the same as what
// was written by testWriteShort. N is v.Len().
// This function is only used for testing.
func testReadShort(v Var, n uint64) error {
	data := make([]int16, n)
	if err := v.ReadShort(data); err != nil {
		return err
	}
	for i := 0; i < int(n); i++ {
		if val := int16(i + 10); data[i] != val {
			return fmt.Errorf("data at position %d is %f; expected %f\n", i, data[i], val)
		}
	}
	return nil
}
