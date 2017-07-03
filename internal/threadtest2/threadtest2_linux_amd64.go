// Code generated by ccgo. DO NOT EDIT.

// threadtest2.c

// 2004 January 13
//
// The author disclaims copyright to this source code.  In place of
// a legal notice, here is a blessing:
//
// May you do good and not evil.
// May you find forgiveness for yourself and forgive others.
// May you share freely, never taking more than you give.
//
// *************************************************************************
// This file implements a simple standalone program used to test whether
// or not the SQLite library is threadsafe.
//
// This file is NOT part of the standard SQLite library.  It is used for
// testing only.

package main

import (
	"math"
	"os"
	"unsafe"

	"github.com/cznic/crt"
	"github.com/cznic/sqlite/internal/bin"
)

var argv []*int8

func main() {
	for _, v := range os.Args {
		argv = append(argv, (*int8)(crt.CString(v)))
	}
	argv = append(argv, nil)
	X_start(crt.NewTLS(), int32(len(os.Args)), &argv[0])
}

func X_start(tls *crt.TLS, _argc int32, _argv **int8) {
	crt.X__register_stdfiles(tls, Xstdin, Xstdout, Xstderr)
	crt.X__builtin_exit(tls, Xmain(tls, _argc, _argv))
}

var Xstdin unsafe.Pointer

func init() {
	Xstdin = unsafe.Pointer(&X__stdfiles)
}

var X__stdfiles [3]unsafe.Pointer

var Xstdout unsafe.Pointer

func init() {
	Xstdout = (unsafe.Pointer)(uintptr(unsafe.Pointer(&X__stdfiles)) + 8)
}

var Xstderr unsafe.Pointer

func init() {
	Xstderr = (unsafe.Pointer)(uintptr(unsafe.Pointer(&X__stdfiles)) + 16)
}

// Initialize the database and start the threads
func Xmain(tls *crt.TLS, _argc int32, _argv **int8) (r0 int32) {
	var _i, _rc int32
	var _1_zJournal *int8
	var _db unsafe.Pointer
	var _aThread [5]uint64
	r0 = i32(0)
	if crt.Xstrcmp(tls, str(0), str(8)) != 0 {
		_1_zJournal = bin.Xsqlite3_mprintf(tls, str(17), unsafe.Pointer(str(0)))
		crt.Xunlink(tls, str(0))
		crt.Xunlink(tls, _1_zJournal)
		bin.Xsqlite3_free(tls, (unsafe.Pointer)(_1_zJournal))
	}
	bin.Xsqlite3_open(tls, str(0), (**bin.Xsqlite3)(unsafe.Pointer(&_db)))
	if _db == nil {
		crt.Xfprintf(tls, (*crt.XFILE)(Xstderr), str(28))
		crt.Xexit(tls, i32(1))
	}
	_rc = bin.Xsqlite3_exec(tls, (*bin.Xsqlite3)(_db), str(59), nil, nil, nil)
	if _rc != 0 {
		crt.Xfprintf(tls, (*crt.XFILE)(Xstderr), str(79), _rc)
		crt.Xexit(tls, i32(1))
	}
	bin.Xsqlite3_close(tls, (*bin.Xsqlite3)(_db))
	_i = i32(0)
_3:
	if uint64(_i) >= u64(5) {
		goto _6
	}
	crt.Xpthread_create(tls, (*uint64)(unsafe.Pointer(uintptr((unsafe.Pointer)(&_aThread))+8*uintptr(_i))), nil, Xworker, (unsafe.Pointer)(uintptr(_i)))
	_i += 1
	goto _3
_6:
	_i = i32(0)
_7:
	if uint64(_i) >= u64(5) {
		goto _10
	}
	crt.Xpthread_join(tls, *(*uint64)(unsafe.Pointer(uintptr((unsafe.Pointer)(&_aThread)) + 8*uintptr(_i))), nil)
	_i += 1
	goto _7
_10:
	if Xall_stop == 0 {
		crt.Xprintf(tls, str(107))
		return i32(0)
	}
	crt.Xprintf(tls, str(129))
	return i32(1)

	_ = _aThread
	panic(0)
}

// This is the worker thread
func Xworker(tls *crt.TLS, _workerArg unsafe.Pointer) (r0 unsafe.Pointer) {
	var _id, _rc, _cnt int32
	var _db unsafe.Pointer
	_id = int32(uintptr(_workerArg))
	_cnt = i32(0)
	crt.Xfprintf(tls, (*crt.XFILE)(Xstderr), str(147), _id)
_0:
	if Xall_stop != 0 || postInc0(&_cnt, int32(1)) >= i32(10000) {
		goto _1
	}
	if (_cnt % i32(100)) == i32(0) {
		crt.Xprintf(tls, str(167), _id, _cnt)
	}
_3:
	if bin.Xsqlite3_open(tls, str(0), (**bin.Xsqlite3)(unsafe.Pointer(&_db))) != i32(0) {
		crt.Xsched_yield(tls)
		goto _3
	}
	bin.Xsqlite3_exec(tls, (*bin.Xsqlite3)(_db), str(175), nil, nil, nil)
	if Xall_stop != 0 {
		bin.Xsqlite3_close(tls, (*bin.Xsqlite3)(_db))
		goto _1
	}
	_rc = bin.Xsqlite3_exec(tls, (*bin.Xsqlite3)(_db), str(198), nil, nil, nil)
	bin.Xsqlite3_close(tls, (*bin.Xsqlite3)(_db))
	goto _0
_1:
	crt.Xfprintf(tls, (*crt.XFILE)(Xstderr), str(234), _id)
	return nil

	_ = _rc
	panic(0)
}

var Xall_stop int32

func bool2int(b bool) int32 {
	if b {
		return 1
	}
	return 0
}
func bug20530(interface{}) {} //TODO remove when https://github.com/golang/go/issues/20530 is fixed.
func i16(n int16) int16    { return n }
func i32(n int32) int32    { return n }
func i64(n int64) int64    { return n }
func i8(n int8) int8       { return n }
func init()                { nzf32 *= -1; nzf64 *= -1 }
func u16(n uint16) uint16  { return n }
func u32(n uint32) uint32  { return n }
func u64(n uint64) uint64  { return n }
func u8(n byte) byte       { return n }

var inf = math.Inf(1)
var nzf32 float32                      // -0.0
var nzf64 float64                      // -0.0
func postInc0(p *int32, d int32) int32 { v := *p; *p += d; return v }
func str(n int) *int8                  { return (*int8)(unsafe.Pointer(&strTab[n])) }
func wstr(n int) *int32                { return (*int32)(unsafe.Pointer(&strTab[n])) }

var strTab = []byte("test.db\x00:memory:\x00%s-journal\x00unable to initialize database\x0a\x00CREATE TABLE t1(x);\x00cannot create table t1: %d\x0a\x00Everything seems ok.\x0a\x00We hit an error.\x0a\x00Starting worker %d\x0a\x00%d: %d\x0a\x00PRAGMA synchronous=OFF\x00INSERT INTO t1 VALUES('bogus data')\x00Worker %d finished\x0a\x00")