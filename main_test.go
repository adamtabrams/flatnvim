package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrimPath(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		curDir   string
		vimDir   string
		homeDir  string
		expected string
	}{
		{
			name:     "simple",
			path:     "file.txt",
			expected: "file.txt",
		},
		{
			name:     "two",
			path:     "/dir/file.txt",
			expected: "/dir/file.txt",
		},
		{
			name:     "simple vim dir",
			path:     "/dir/file.txt",
			vimDir:   "/dir/",
			expected: "file.txt",
		},
		{
			name:     "skip vim dir",
			path:     "/dir/file.txt",
			vimDir:   "",
			expected: "/dir/file.txt",
		},
		{
			name:     "not vim dir",
			path:     "/other/file.txt",
			vimDir:   "/dir/",
			expected: "/other/file.txt",
		},
		{
			name:     "inner vim dir",
			path:     "/other/dir/file.txt",
			vimDir:   "/dir/",
			expected: "/other/dir/file.txt",
		},
		{
			name:     "simple home dir",
			path:     "/dir/home/file.txt",
			homeDir:  "/dir/home",
			expected: "~/file.txt",
		},
		{
			name:     "skip home dir",
			path:     "/dir/home/file.txt",
			homeDir:  "",
			expected: "/dir/home/file.txt",
		},
		{
			name:     "not home dir",
			path:     "/dir/other/file.txt",
			homeDir:  "/dir/home",
			expected: "/dir/other/file.txt",
		},
		{
			name:     "inner home dir",
			path:     "/other/dir/home/file.txt",
			homeDir:  "/dir/home",
			expected: "/other/dir/home/file.txt",
		},
		{
			name:     "dot slash",
			path:     "./file.txt",
			expected: "file.txt",
		},
		{
			name:     "inner dot slash",
			path:     "/dir/./file.txt",
			expected: "/dir/file.txt",
		},
		{
			name:     "double dot slash",
			path:     "/dir/../file.txt",
			expected: "/dir/../file.txt",
		},
		{
			name:     "prepend current dir",
			path:     "file.txt",
			curDir:   "/dir/",
			expected: "/dir/file.txt",
		},
		{
			name:     "skip prepend current dir",
			path:     "/dir/file.txt",
			curDir:   "/dir/",
			expected: "/dir/file.txt",
		},
		{
			name:     "already absolute path",
			path:     "/other/file.txt",
			curDir:   "/dir/",
			expected: "/other/file.txt",
		},
		{
			name:     "prepend current and trim home dir",
			path:     "file.txt",
			curDir:   "/home/dir/",
			homeDir:  "/home",
			expected: "~/dir/file.txt",
		},
		{
			name:     "prepend current and trim vim dir",
			path:     "file.txt",
			curDir:   "/vim/dir/",
			vimDir:   "/vim/",
			expected: "dir/file.txt",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := trimPath(tt.path, tt.curDir, tt.vimDir, tt.homeDir)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
