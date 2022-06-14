package template

import (
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
)

type Module struct {
	Pos   lexer.Position
	Nodes []*Xml `parser:"{ @@ }"`
}

type KeyValue struct {
	Pos   lexer.Position
	Key   string   `parser:"@\":\"? @Ident ( @\"-\" @Ident )*"`
	Value *Literal `parser:"\"=\" @@"`
}

type Literal struct {
	Str string `parser:"@String"`
	Ref string `parser:"| @\"{\" @Ident ( @\".\" @Ident )* @\"}\""`
}

type Xml struct {
	Pos        lexer.Position
	Name       string      `parser:"\"<\"@Ident"`
	Parameters []*KeyValue `parser:"[ @@ { @@ } ] \">\""`
	Children   []*Xml      `parser:"{ @@ }"`
	Value      *Literal    `parser:"{ @@ }"` // Todo make this match with @String or Literal
	Close      string      `parser:"\"<\"\"/\"@Ident\">\""`
}

var xmlParser = participle.MustBuild(&Module{})
