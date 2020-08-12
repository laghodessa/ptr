package ptr

import "time"

func Bool(v bool) *bool           { return &v }
func Byte(v byte) *byte           { return &v }
func Int(v int) *int              { return &v }
func Int8(v int8) *int8           { return &v }
func Int16(v int16) *int16        { return &v }
func Int32(v int32) *int32        { return &v }
func Int64(v int64) *int64        { return &v }
func String(v string) *string     { return &v }
func Time(v time.Time) *time.Time { return &v }
func UInt(v uint) *uint           { return &v }
func UInt8(v uint8) *uint8        { return &v }
func UInt16(v uint16) *uint16     { return &v }
func UInt32(v uint32) *uint32     { return &v }
func UInt64(v uint64) *uint64     { return &v }

func NewBool(v **bool) *bool           { *v = new(bool); return *v }
func NewByte(v **byte) *byte           { *v = new(byte); return *v }
func NewInt(v **int) *int              { *v = new(int); return *v }
func NewInt8(v **int8) *int8           { *v = new(int8); return *v }
func NewInt16(v **int16) *int16        { *v = new(int16); return *v }
func NewInt32(v **int32) *int32        { *v = new(int32); return *v }
func NewInt64(v **int64) *int64        { *v = new(int64); return *v }
func NewString(v **string) *string     { *v = new(string); return *v }
func NewTime(v **time.Time) *time.Time { *v = new(time.Time); return *v }
func NewUInt(v **uint) *uint           { *v = new(uint); return *v }
func NewUInt8(v **uint8) *uint8        { *v = new(uint8); return *v }
func NewUInt16(v **uint16) *uint16     { *v = new(uint16); return *v }
func NewUInt32(v **uint32) *uint32     { *v = new(uint32); return *v }
func NewUInt64(v **uint64) *uint64     { *v = new(uint64); return *v }
