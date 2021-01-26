package ch7

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Basic represents properties common to files and directories
type Basic struct {
	name      string
	createdAt time.Time
	updatedAt time.Time
	parent    *Directory
}

// Directory is a directory in the in-memory filesystem. It can have
// subdirectories and files.
type Directory struct {
	Basic
	subdirectories []*Directory
	files          []*File
}

// File is a file in the in-memory filesystem. It stores binary data.
type File struct {
	Basic
	contents *[]byte
}

// TODO: Users, groups, permissions

// newDirectory creates a directory with the given name under parent.
func newDirectory(name string, parent *Directory) *Directory {
	directory := Directory{
		Basic: Basic{
			name:      name,
			createdAt: time.Now(),
			updatedAt: time.Now(),
			parent:    parent,
		},
	}
	if parent != nil {
		parent.subdirectories = append(parent.subdirectories, &directory)
	}
	return &directory
}

// NewRootDirectory creates a new top-level directory.
func NewRootDirectory() *Directory {
	return newDirectory("", nil)
}

// Mkdir creates a subdirectory under the parent.
func (d *Directory) Mkdir(name string) *Directory {
	return newDirectory(name, d)
}

// Touch creates a 0-byte file.
func (d *Directory) Touch(name string) *File {
	file := File{
		Basic: Basic{
			name:      name,
			createdAt: time.Now(),
			updatedAt: time.Now(),
			parent:    d,
		},
	}
	d.files = append(d.files, &file)
	return &file
}

// Set sets the binary data of the file.
func (f *File) Set(contents *[]byte) {
	f.contents = contents
}

// Get returns the binary data of the file.
func (f *File) Get() *[]byte {
	return f.contents
}

// RenameTo renames the file or directory.
func (b *Basic) RenameTo(name string) {
	b.name = name
	b.updatedAt = time.Now()
}

// GetName returns the name of the file or directory.
func (b *Basic) GetName() string {
	return b.name
}

func Test(t *testing.T) {
	root := NewRootDirectory()
	foo := root.Mkdir("bar")
	foo.RenameTo("foo")
	assert.Equal(t, "foo", foo.GetName())
	qux := foo.Touch("qux.txt")
	contents := []byte("Hello, world!")
	qux.Set(&contents)
	assert.Equal(t, "Hello, world!", string(*qux.Get()))
}
