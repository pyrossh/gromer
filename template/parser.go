package template

import (
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
)

type Module struct {
	Pos   lexer.Position
	Nodes []*Xml `@@`
}

type Attribute struct {
	Pos   lexer.Position
	Key   string   `parser:"@\":\"? @Ident ( @\"-\" @Ident )*"`
	Value *Literal `parser:"\"=\" @@"`
}

type KV struct {
	Pos   lexer.Position
	Key   string `@Ident`
	Value string `":" @"!"? @Ident ( @"." @Ident )*`
}

type Literal struct {
	Pos lexer.Position
	Str string `@String`
	Ref string `| "{" @Ident ( @"." @Ident )* "}"`
	KV  []*KV  `| "{""{" @@* "}""}"`
}

type Xml struct {
	Pos        lexer.Position
	Name       string       `parser:"\"<\"@Ident"`
	Attributes []*Attribute `parser:"[ @@ { @@ } ] \">\""`
	Children   []*Xml       `parser:"{ @@ }"`
	Value      *Literal     `parser:"{ @@ }"` // Todo make this match with @String or Literal
	Close      string       `parser:"\"<\"\"/\"@Ident\">\""`
}

var xmlParser = participle.MustBuild(&Module{})
