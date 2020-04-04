package moduledoc

import "fmt"

func (s docHandler) BuildTOC() {
	para := s.doc.AddParagraph()
	run := para.AddRun()
	para.SetStyle("Heading1")
	run.AddText("List of Endpoint")
	para = s.doc.AddParagraph()

	nd := s.doc.Numbering.AddDefinition()
	for path := range s.swagger.Paths {
		pathItems := s.swagger.Paths.Find(path)

		if pathItems.Get != nil {
			para = s.doc.AddParagraph()
			para.SetNumberingDefinition(nd)
			para.AddRun().AddText(fmt.Sprintf("%s %s", MethodGet, path))
		}

		if pathItems.Post != nil {
			para = s.doc.AddParagraph()
			para.SetNumberingDefinition(nd)
			para.AddRun().AddText(fmt.Sprintf("%s %s", MethodPost, path))
		}

		if pathItems.Patch != nil {
			para = s.doc.AddParagraph()
			para.SetNumberingDefinition(nd)
			para.AddRun().AddText(fmt.Sprintf("%s %s", MethodPatch, path))
		}

		if pathItems.Delete != nil {
			para = s.doc.AddParagraph()
			para.SetNumberingDefinition(nd)
			para.AddRun().AddText(fmt.Sprintf("%s %s", MethodDelete, path))
		}
	}

	para.AddRun().AddPageBreak()
}
