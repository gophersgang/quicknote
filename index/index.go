// Quicknote stores and searches tens of thousands of short notes.
//
// Copyright (C) 2017  Andrew Miller <amiller@amilx.com>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package index

import (
	"errors"

	"github.com/anmil/quicknote/index/bleve"
	"github.com/anmil/quicknote/index/elastic"
	"github.com/anmil/quicknote/note"
)

// ErrProviderNotSupported index provider given is not supported
var ErrProviderNotSupported = errors.New("Unsupported provider given")

// Index interface for the index providers
type Index interface {
	IndexNote(n *note.Note) error
	IndexNotes(notes []*note.Note) error
	SearchNote(query string, limit, offset int) ([]int64, uint64, error)
	SearchNotePhrase(query string, bk *note.Book, sort string, limit, offset int) ([]int64, uint64, error)
	DeleteNote(n *note.Note) error
	DeleteBook(bk *note.Book) error
}

// NewIndex returns a new indexer for the given provider
func NewIndex(provider string, options ...string) (Index, error) {
	switch provider {
	case "bleve":
		return bleve.NewIndex(options[0])
	case "elastic":
		return elastic.NewIndex(options[0], options[1])
	default:
		return nil, ErrProviderNotSupported
	}
}
