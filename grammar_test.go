package gositter

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	g := CreateGrammar("root", map[string]Expression{
		"root": Repeat1(
			Choice(
				Ref("whitespace"),
				Ref("tag"),
				Terminal("\n"))),
		"tag": Choice(
			Ref("h6"),
			Ref("h5"),
			Ref("h4"),
			Ref("h3"),
			Ref("h2"),
			Ref("h1"),
			Ref("p"),
			Ref("a")),
		"h1": Seq(Terminal("#"), Ref("p")),
		"h2": Seq(Terminal("##"), Ref("p")),
		"h3": Seq(Terminal("###"), Ref("p")),
		"h4": Seq(Terminal("####"), Ref("p")),
		"h5": Seq(Terminal("#####"), Ref("p")),
		"h6": Seq(Terminal("######"), Ref("p")),
		"p": Repeat1(
			Seq(
				Repeat1(Choice(
					Ref("a"),
					Ref("img"),
					Ref("char"),
					Ref("whitespace"))),
				Optional(Terminal("\n")))),
		"char":       Regex(`[^\s]`),
		"whitespace": Regex(` |\t`),
		"a": Seq(
			Terminal("["),
			Regex(`[^\]]*`),
			Terminal("]"),
			Optional(Seq(
				Terminal("("),
				Regex(`[^\)]*`),
				Terminal(")")))),
		"img": Seq(
			Terminal("!["),
			Regex(`[^\]]*`),
			Terminal("]"),
			Optional(Seq(
				Terminal("("),
				Regex(`[^\)]*`),
				Terminal(")")))),
	})
	input := `
    # Hello!

    ## Brave

    ### New

    #### World!!

    My name is [Samuel](http://samuellando.com) 
    Lando
    yesy yes ![image](http://image.com)
    Hahaha
    `
	st, err := g.Parse(input)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(st.Tree())
	if len(st.Find("tag")) != 5 {
		t.Fatal("5 tags")
	}
	if len(st.Find("a")) != 1 {
		t.Fail()
	}
	if len(st.Find("img")) != 1 {
		t.Fail()
	}
}
