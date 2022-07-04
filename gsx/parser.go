package gsx

import (
	"fmt"
	"strings"

	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/goneric/stack"
)

type Module struct {
	Pos   lexer.Position
	Nodes []*AstNode `@@*`
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

type ForStatement struct {
	Pos        lexer.Position `"for"`
	Index      string         `@Ident ","`
	Key        string         `@Ident`
	Reference  string         `":""=""range" @Ident`
	Statements []*Statement   `"{" @@* "}"`
}

type Statement struct {
	ReturnStatement *ReturnStatement `@@`
}

type ReturnStatement struct {
	Nodes []*AstNode `"return" "(" @@* ")"`
	Tags  []*Tag
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
	Str *string       `@String`
	Ref *string       `| "{" @Ident ( @"." @Ident )* "}"`
	KV  []*KV         `| "{" @@* "}"`
	For *ForStatement `| @@`
}

func (l *Literal) Clone() *Literal {
	if l == nil {
		return nil
	}
	newLiteral := &Literal{}
	if l.Str != nil {
		v := "" + *l.Str
		newLiteral.Str = &v
	}
	if l.Ref != nil {
		v := "" + *l.Ref
		newLiteral.Ref = &v
	}
	if l.KV != nil {
		newLiteral.KV = []*KV{}
		for _, kv := range l.KV {
			newLiteral.KV = append(newLiteral.KV, &KV{
				Key:   "" + kv.Key,
				Value: "" + kv.Value,
			})
		}
	}
	// TODO copy for
	return newLiteral
}

var htmlParser = participle.MustBuild[Module]()

type Tag struct {
	Name        string
	Text        *Literal
	Attributes  []*Attribute
	Children    []*Tag
	SelfClosing bool
}

func (t *Tag) Clone() *Tag {
	newTag := &Tag{
		Name:        t.Name,
		Text:        t.Text.Clone(),
		Attributes:  []*Attribute{},
		SelfClosing: t.SelfClosing,
		Children:    []*Tag{},
	}
	for _, v := range t.Attributes {
		newTag.Attributes = append(newTag.Attributes, &Attribute{
			Key:   v.Key,
			Value: v.Value.Clone(),
		})
	}
	for _, child := range t.Children {
		newTag.Children = append(newTag.Children, child.Clone())
	}
	return newTag
}

func cloneTags(tags []*Tag) []*Tag {
	newTags := []*Tag{}
	for _, v := range tags {
		newTags = append(newTags, v.Clone())
	}
	return newTags
}

func RenderString(tags []*Tag) string {
	s := ""
	for _, t := range tags {
		s += RenderTagString(t, "") + "\n"
	}
	return s
}

func RenderTagString(x *Tag, space string) string {
	if x.Name == "" {
		if x.Text != nil && x.Text.Str != nil {
			return space + strings.ReplaceAll(*x.Text.Str, `"`, "")
		}
		if x.Text != nil && x.Text.Ref != nil {
			return space + "{" + *x.Text.Ref + "}"
		}
	}
	if x.Name == "fragment" {
		s := ""
		for _, c := range x.Children {
			s += RenderTagString(c, space) + "\n"
		}
		return s
	}
	s := space + "<" + x.Name
	for _, a := range x.Attributes {
		if a.Value.Str != nil && *a.Value.Str != "" {
			s += " " + a.Key + `="` + *a.Value.Str + `"`
		}
	}
	if x.SelfClosing {
		s += " />"
	} else {
		s += ">\n"
	}
	if !x.SelfClosing {
		for _, c := range x.Children {
			s += RenderTagString(c, space+"  ") + "\n"
		}
		s += space + "</" + x.Name + ">"
	}
	return s
}

func processTree(nodes []*AstNode) []*Tag {
	tags := []*Tag{}
	var prevTag *Tag
	stack := stack.New[*Tag]()
	for _, n := range nodes {
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
				if !newTag.SelfClosing {
					prevTag = newTag
				}
			}
		} else if n.Close != nil {
			if n.Close.Name == prevTag.Name {
				prevTag, _ = stack.Pop()
			} else {
				panic(fmt.Errorf("Brackets not matching for tag %s in line %d:%d, prevTag: %s", n.Close.Name, n.Close.Pos.Line, n.Close.Pos.Column, prevTag.Name))
			}
		} else if n.Content != nil {
			newTag := &Tag{
				Name: "",
				Text: n.Content,
			}
			if n.Content.For != nil {
				for _, s := range n.Content.For.Statements {
					if s.ReturnStatement != nil {
						s.ReturnStatement.Tags = processTree(s.ReturnStatement.Nodes)
					}
				}
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
	return processTree(ast.Nodes)
}
