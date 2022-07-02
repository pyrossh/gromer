package gsx

import (
	"fmt"
	"strings"

	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/alecthomas/repr"
	"github.com/goneric/stack"
)

type Module struct {
	Pos     lexer.Position
	AstNode []*AstNode `@@*`
}

type AstNode struct {
	Pos     lexer.Position
	Open    *Open    `@@`
	Close   *Close   `| @@`
	Content *Literal `| @@`
}

type Open struct {
	Pos        lexer.Position
	Name       string       `"<" @Ident`
	Attributes []*Attribute `[ @@ { @@ } ]`
	SelfClose  string       `@("/")? ">"`
}

type Close struct {
	Pos  lexer.Position
	Name string `"<""/"@Ident">"`
}

type Attribute struct {
	Pos   lexer.Position
	Key   string   `@":"? @Ident ( @"-" @Ident )*`
	Value *Literal `"=" @@`
}

type KV struct {
	Pos   lexer.Position
	Key   string `@String`
	Value string `":" @"!"? @Ident ( @"." @Ident )*`
}

type Literal struct {
	Pos lexer.Position
	Str *string `@String`
	Ref *string `| "{" @Ident ( @"." @Ident )* "}"`
	KV  []*KV   `| "{" @@* "}"`
}

var htmlParser = participle.MustBuild[Module]()

type Tag struct {
	Name        string
	Text        *Literal
	Attributes  []*Attribute
	Children    []*Tag
	SelfClosing bool
}

func renderString(tags []*Tag) string {
	s := ""
	for _, t := range tags {
		s += renderTagString(t, "") + "\n"
	}
	return s
}

func renderTagString(x *Tag, space string) string {
	if x.Name == "" {
		if x.Text != nil && x.Text.Str != nil {
			return space + strings.ReplaceAll(*x.Text.Str, `"`, "")
		}
		if x.Text != nil && x.Text.Ref != nil {
			return space + "{" + *x.Text.Ref + "}"
		}
	}
	s := space + "<" + x.Name
	if len(x.Attributes) > 0 {
		s += " "
	}
	for _, a := range x.Attributes {
		if a.Value.Str != nil {
			s += a.Key + "=" + *a.Value.Str
		}
	}
	if x.SelfClosing {
		s += " />"
	} else {
		s += ">\n"
	}
	if !x.SelfClosing {
		for _, c := range x.Children {
			s += renderTagString(c, space+"  ") + "\n"
		}
		s += space + "</" + x.Name + ">"
	}
	return s
}

func processTree(module *Module) []*Tag {
	tags := []*Tag{}
	var prevTag *Tag
	stack := stack.New[*Tag]()
	for _, n := range module.AstNode {
		if n.Open != nil {
			newTag := &Tag{
				Name:        n.Open.Name,
				Attributes:  n.Open.Attributes,
				SelfClosing: n.Open.SelfClose == "/",
			}
			if prevTag != nil {
				prevTag.Children = append(prevTag.Children, newTag)
				if !newTag.SelfClosing {
					stack.Push(prevTag)
					prevTag = newTag
				}
			} else {
				tags = append(tags, newTag)
				prevTag = newTag
			}
		} else if n.Close != nil {
			repr.Println("close", n.Close.Name, prevTag.Name)
			if n.Close.Name == prevTag.Name {
				prevTag, _ = stack.Pop()
			} else {
				panic(fmt.Errorf("Brackets not matching in line %d:%d", n.Close.Pos.Line, n.Close.Pos.Column))
			}
		} else if n.Content != nil {
			newTag := &Tag{
				Name: "",
				Text: n.Content,
			}
			if prevTag != nil {
				prevTag.Children = append(prevTag.Children, newTag)
			} else {
				tags = append(tags, newTag)
			}
		}
	}
	return tags
}

func parse(name, s string) []*Tag {
	ast, err := htmlParser.ParseString(name, s)
	if err != nil {
		panic(err)
	}
	return processTree(ast)
}
