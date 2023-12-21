package main

import (
	"fmt"
	"strings"
)

const TagSelfClose string = "/>"

const TagDoctype string = "<!doctype html>"

const (
	TagHTMLEnglish string = "<html lang=\"en\">"
	TagHTMLClose   string = "</html>"
)

const (
	TagHead      string = "<head>"
	TagHeadClose string = "</head>"
)

const TagMetaOpen string = "<meta"

const (
	TagScriptOpen  string = "<script"
	TagScriptClose string = "></script>"
)

const TagLinkOpen string = "<link"

type DocumentBuilder struct {
	b *strings.Builder
}

func Builder() (b *DocumentBuilder) {
	var sb strings.Builder
	sb.WriteString(TagDoctype)
	sb.WriteString(TagHTMLEnglish)
	b = &DocumentBuilder{
		b: &sb,
	}
	return
}

func (b *DocumentBuilder) Build() string {
	b.b.WriteString(TagHTMLClose)
	return b.b.String()
}

type HeadBuilder struct {
	b *strings.Builder
}

func Head() (b *HeadBuilder) {
	var sb strings.Builder
	sb.WriteString(TagHead)
	b = &HeadBuilder{
		b: &sb,
	}
	return
}

func (b *HeadBuilder) Prepare() string {
	return b.b.String()
}

func (b *HeadBuilder) Build() string {
	b.b.WriteString(TagHeadClose)
	return b.b.String()
}

func (b *HeadBuilder) Meta(attrs ...string) {
	b.b.WriteString(TagMetaOpen)
	for _, attr := range attrs {
		fmt.Fprintf(b.b, " %s", attr)
	}
	b.b.WriteString(TagSelfClose)
}

func (b *HeadBuilder) Script(attrs ...string) {
	b.b.WriteString(TagScriptOpen)
	for _, attr := range attrs {
		fmt.Fprintf(b.b, " %s", attr)
	}
	b.b.WriteString(TagScriptClose)
}

func (b *HeadBuilder) Link(attrs ...string) {
	b.b.WriteString(TagScriptOpen)
	for _, attr := range attrs {
		fmt.Fprintf(b.b, " %s", attr)
	}
	b.b.WriteString(TagSelfClose)
}

type BodyBuilder struct{}
