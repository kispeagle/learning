package template

import "fmt"

type ITemplateHtml interface {
	Header() string
	Body() string
	Footer() string
}

type TemplateHtml struct {
	ITH ITemplateHtml
}

func (t *TemplateHtml) ShowHtmlPage() {
	fmt.Println(t.ITH.Header())
	fmt.Println(t.ITH.Body())
	fmt.Println(t.ITH.Footer())
}

type HomePage struct {
}

func (h *HomePage) Header() string {
	return "<Header> Home Page </Header>"
}

func (h *HomePage) Body() string {
	return "<Body> Content </Body>"
}

func (h *HomePage) Footer() string {
	return "<Footer> footer </Footer>"
}

type ContactPage struct {
}

func (c *ContactPage) Header() string {
	return "<Header> Contact Pages </Header>"
}

func (c *ContactPage) Body() string {
	return "<Body> Content </Body>"
}

func (c *ContactPage) Footer() string {
	return "<Footer> footer </Footer>"
}

// func main() {
// 	c := &template.ContactPage{}
// 	t := template.TemplateHtml{c}
// 	t.ShowHtmlPage()

// 	h := &template.HomePage{}
// 	t = template.TemplateHtml{h}

// 	t.ShowHtmlPage()

// }
