package c

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
	. "github.com/alecthomas/chroma/lexers/p" // nolint
)

// Cheetah lexer.
var Cheetah = internal.Register(MustNewLexer(
	&Config{
		Name:      "Cheetah",
		Aliases:   []string{"cheetah", "spitfire"},
		Filenames: []string{"*.tmpl", "*.spt"},
		MimeTypes: []string{"application/x-cheetah", "application/x-spitfire"},
	},
	Rules{
		"root": {
			{`(##[^\n]*)$`, ByGroups(Comment), nil},
			{`#[*](.|\n)*?[*]#`, Comment, nil},
			{`#end[^#\n]*(?:#|$)`, CommentPreproc, nil},
			{`#slurp$`, CommentPreproc, nil},
			{`(#[a-zA-Z]+)([^#\n]*)(#|$)`, ByGroups(CommentPreproc, Using(Python, nil), CommentPreproc), nil},
			{`(\$)([a-zA-Z_][\w.]*\w)`, ByGroups(CommentPreproc, Using(Python, nil)), nil},
			{`(\$\{!?)(.*?)(\})(?s)`, ByGroups(CommentPreproc, Using(Python, nil), CommentPreproc), nil},
			{`(?sx)
                (.+?)               # anything, followed by:
                (?:
                 (?=\#[#a-zA-Z]*) | # an eval comment
                 (?=\$[a-zA-Z_{]) | # a substitution
                 \Z                 # end of string
                )
            `, Other, nil},
			{`\s+`, Text, nil},
		},
	},
))
