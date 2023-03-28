//go:build darwin

package tokenizer

/*
#cgo LDFLAGS: ${SRCDIR}/tiktoken-cffi/libtiktoken.a -ldl -framework Security -framework CoreFoundation
#include <stdlib.h>

extern unsigned int count_tokens(const char*, const char*);
extern unsigned int get_context_size(const char*);
*/
import "C"
import (
	"github.com/kamushadenes/chloe/models"
	"unsafe"
)

func CountTokens(model models.Model, prompt string) int {
	m := C.CString(string(model))
	p := C.CString(prompt)
	count := C.count_tokens(m, p)
	C.free(unsafe.Pointer(m))
	C.free(unsafe.Pointer(p))
	return int(count)
}

func GetContextSize(model models.Model) int {
	m := C.CString(string(model))
	count := C.get_context_size(m)
	C.free(unsafe.Pointer(m))
	return int(count)
}
