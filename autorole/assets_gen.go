// Code generated by "esc -o assets_gen.go -pkg autorole -ignore .go assets/"; DO NOT EDIT.

package autorole

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"sync"
	"time"
)

type _escLocalFS struct{}

var _escLocal _escLocalFS

type _escStaticFS struct{}

var _escStatic _escStaticFS

type _escDirectory struct {
	fs   http.FileSystem
	name string
}

type _escFile struct {
	compressed string
	size       int64
	modtime    int64
	local      string
	isDir      bool

	once sync.Once
	data []byte
	name string
}

func (_escLocalFS) Open(name string) (http.File, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	return os.Open(f.local)
}

func (_escStaticFS) prepare(name string) (*_escFile, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	var err error
	f.once.Do(func() {
		f.name = path.Base(name)
		if f.size == 0 {
			return
		}
		var gr *gzip.Reader
		b64 := base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(f.compressed))
		gr, err = gzip.NewReader(b64)
		if err != nil {
			return
		}
		f.data, err = ioutil.ReadAll(gr)
	})
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (fs _escStaticFS) Open(name string) (http.File, error) {
	f, err := fs.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.File()
}

func (dir _escDirectory) Open(name string) (http.File, error) {
	return dir.fs.Open(dir.name + name)
}

func (f *_escFile) File() (http.File, error) {
	type httpFile struct {
		*bytes.Reader
		*_escFile
	}
	return &httpFile{
		Reader:   bytes.NewReader(f.data),
		_escFile: f,
	}, nil
}

func (f *_escFile) Close() error {
	return nil
}

func (f *_escFile) Readdir(count int) ([]os.FileInfo, error) {
	return nil, nil
}

func (f *_escFile) Stat() (os.FileInfo, error) {
	return f, nil
}

func (f *_escFile) Name() string {
	return f.name
}

func (f *_escFile) Size() int64 {
	return f.size
}

func (f *_escFile) Mode() os.FileMode {
	return 0
}

func (f *_escFile) ModTime() time.Time {
	return time.Unix(f.modtime, 0)
}

func (f *_escFile) IsDir() bool {
	return f.isDir
}

func (f *_escFile) Sys() interface{} {
	return f
}

// FS returns a http.Filesystem for the embedded assets. If useLocal is true,
// the filesystem's contents are instead used.
func FS(useLocal bool) http.FileSystem {
	if useLocal {
		return _escLocal
	}
	return _escStatic
}

// Dir returns a http.Filesystem for the embedded assets on a given prefix dir.
// If useLocal is true, the filesystem's contents are instead used.
func Dir(useLocal bool, name string) http.FileSystem {
	if useLocal {
		return _escDirectory{fs: _escLocal, name: name}
	}
	return _escDirectory{fs: _escStatic, name: name}
}

// FSByte returns the named file from the embedded assets. If useLocal is
// true, the filesystem's contents are instead used.
func FSByte(useLocal bool, name string) ([]byte, error) {
	if useLocal {
		f, err := _escLocal.Open(name)
		if err != nil {
			return nil, err
		}
		b, err := ioutil.ReadAll(f)
		_ = f.Close()
		return b, err
	}
	f, err := _escStatic.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.data, nil
}

// FSMustByte is the same as FSByte, but panics if name is not present.
func FSMustByte(useLocal bool, name string) []byte {
	b, err := FSByte(useLocal, name)
	if err != nil {
		panic(err)
	}
	return b
}

// FSString is the string version of FSByte.
func FSString(useLocal bool, name string) (string, error) {
	b, err := FSByte(useLocal, name)
	return string(b), err
}

// FSMustString is the string version of FSMustByte.
func FSMustString(useLocal bool, name string) string {
	return string(FSMustByte(useLocal, name))
}

var _escData = map[string]*_escFile{

	"/assets/settings.html": {
		local:   "assets/settings.html",
		size:    3217,
		modtime: 1518284137,
		compressed: `
H4sIAAAAAAAC/8xXTW/jNhO++1fMy8u7ASoL6XFhCwiaxXYLpCm2AXosKGkksaE4LDmyawj67wUpWbEd
O+4HttgcYofDGc48z3yl70uslEEQhf1VdkyONIphWCz6nrG1WvIoa1CWApbDsFiVagOFlt6vhaOtyBYA
AIenBelE18ntt5MsypvbvdjKGpNgD53I7qY3V2lze3DbZh/VBoE6hiD1EHxrJatCar2Dihy02Obo/Cq1
kwdpqTbT1/8lCaTL2Q9IkmwxyU/ikhod+ymyUc3RdlSoyLXQIjdUroUlzwJkwYrMWqStNLLGtO+XdwWr
DX7slC6Xn+6HIZ1hfI3MC15/BbPTO1Ya1BB/JyVWstN8cvusRpJTuTtzcQL6u0aaGj20cgcsnxE6C0wg
oVWmY4QcK3II3OAOGrlBkGYHWFVY8Iz9Wz4EFJPaUWcv+BAVtMxRB17XYo9fMoJ4d8S89F7VZk8+cKM8
jPkTLbzxgkeNBYMqT184crQgw460ACNbXIvPLzxe+un7YOXRhsTwcJQOn2PqLr9XdYOew1+w3Cd8FIL4
kQzCu1J5mWssb0LpXYwgHUO4APlL+n8ZNsrOyRCjyB5iYnigas9Doyw4/L1TDstYnEED3q3y7F75QkvV
onu/SvMMflFagyGGLblnUBXsqBvTSjF45JB6ZPQO6lD/Y+2Tgd9ImZvrHCtjOwbeWVwL0wXXLrB7lAVz
ZHvSp0ju53OrZYEN6RLdWgjYSN3hWoTin9k80RkG8V/xlE1vQ0glqkKl+j10TJAjWIceDQcgucGJtD2a
q9yl18vmGJmY2DO0badZTfeOYY4Sq3G6Y0O1BeinXIn15/9OfT0EM+eKzCgNr9mIsq+opLJPtQnN1CJZ
jbBV3ERGKtKatsrUI23/hJrR8r8lRkUrX4yXAye/DlYOJ3CDxXNOf1wJe3WlB53tRbPtiaxHo3eP5gdS
RkDfq+oAohfRMEQ1LPseTTkMGQTZfgLGDrtt0IyTOTTIb6Ck2Ftj71QMuSyew6fi/3tw2NImtGdHbdBp
QVaMbitd6d+O+GrbfYORUXRx9+icQ8NzVPsKCH1rVVCJWd8vf3JUoPfK1MOwSuPpfu4s4cPT3ftzVz88
3c2322levXu4uuJATgyepWN/4pK/ubzt2OyJwDPZcYI1OAcR57s4q3kGtXNHeccc+nbMI9/lreK5vnM2
kLNJrFOtdLv4XdfxI9dUPAvIfpYbXKWjkQwAjo2PC++4V4aV94IjZ/fp8zv3tDwfvrJYpaEYX63eFRGH
Ab0c/92IOb74MwAA//8TC/AakQwAAA==
`,
	},

	"/": {
		isDir: true,
		local: "",
	},

	"/assets": {
		isDir: true,
		local: "assets",
	},
}
