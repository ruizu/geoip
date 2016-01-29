package maxminddb

// #cgo CFLAGS: -I/usr/local/include
// #cgo LDFLAGS: -L/usr/local/lib -lmaxminddb
/*
#include <maxminddb.h>
*/
import "C"

import (
	"fmt"
)

const (
	StatusSuccess                         int = C.MMDB_SUCCESS
	StatusFileOpenError                   int = C.MMDB_FILE_OPEN_ERROR
	StatusIOError                         int = C.MMDB_IO_ERROR
	StatusCorruptSearchTreeError          int = C.MMDB_CORRUPT_SEARCH_TREE_ERROR
	StatusInvalidMetadataError            int = C.MMDB_INVALID_METADATA_ERROR
	StatusUnknownDatabaseFormatError      int = C.MMDB_UNKNOWN_DATABASE_FORMAT_ERROR
	StatusOutOfMemoryError                int = C.MMDB_OUT_OF_MEMORY_ERROR
	StatusInvalidDataError                int = C.MMDB_INVALID_DATA_ERROR
	StatusInvalidLookupPathError          int = C.MMDB_INVALID_LOOKUP_PATH_ERROR
	StatusLookupPathDoesNotMatchDataError int = C.MMDB_LOOKUP_PATH_DOES_NOT_MATCH_DATA_ERROR

	ModeMMap uint32 = C.MMDB_MODE_MMAP

	DataTypeUTF8String int = C.MMDB_DATA_TYPE_UTF8_STRING
	DataTypeDouble     int = C.MMDB_DATA_TYPE_DOUBLE
	DataTypeBytes      int = C.MMDB_DATA_TYPE_BYTES
	DataTypeUInt16     int = C.MMDB_DATA_TYPE_UINT16
	DataTypeUInt32     int = C.MMDB_DATA_TYPE_UINT32
	DataTypeMap        int = C.MMDB_DATA_TYPE_MAP
	DataTypeInt32      int = C.MMDB_DATA_TYPE_INT32
	DataTypeUInt64     int = C.MMDB_DATA_TYPE_UINT64
	DataTypeUInt128    int = C.MMDB_DATA_TYPE_UINT128
	DataTypeArray      int = C.MMDB_DATA_TYPE_ARRAY
	DataTypeBoolean    int = C.MMDB_DATA_TYPE_BOOLEAN
	DataTypeFloat      int = C.MMDB_DATA_TYPE_FLOAT
)

type DB struct {
	mmdb C.MMDB_s
}

func Open(f string, m uint32) (*DB, error) {
	var mmdb C.MMDB_s
	s := int(C.MMDB_open(C.CString(f), C.uint32_t(m), &mmdb))
	if s != StatusSuccess {
		return nil, fmt.Errorf(ErrorString(s))
	}
	return &DB{mmdb}, nil
}

func (db *DB) Close() {
	C.MMDB_close(&db.mmdb)
}

func (db *DB) Lookup(ip string) (interface{}, error) {
	var gaiError, mmdbError C.int
	result := C.MMDB_lookup_string(&db.mmdb, C.CString(ip), &gaiError, &mmdbError)
	if gaiError != 0 {
		return nil, fmt.Errorf(C.GoString(C.gai_strerror(gaiError)))
	}
	if mmdbError != C.int(StatusSuccess) {
		return nil, fmt.Errorf(ErrorString(int(mmdbError)))
	}
	if result.found_entry != C._Bool(true) {
		return nil, fmt.Errorf("no entry for ip (%s) was found.", ip)
	}
	return nil, nil
}

func ErrorString(code int) string {
	return C.GoString(C.MMDB_strerror(C.int(code)))
}
