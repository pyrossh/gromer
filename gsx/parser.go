package gsx

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
	Key   string   `@":"? @Ident ( @"-" @Ident )*`
	Value *Literal `"=" @@`
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
	Name       string       `"<" @Ident`
	Attributes []*Attribute `[ @@ { @@ } ]">"`
	Children   []*Xml       `{ @@ }`
	Value      *Literal     `{ @@ }` // Todo make this match with @String or Literal
	Close      string       `("<""/"@Ident">")?`
}

var xmlParser = participle.MustBuild(&Module{})
