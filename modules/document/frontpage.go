package moduledoc

import (
	"github.com/unidoc/unioffice/color"
	"github.com/unidoc/unioffice/measurement"
	"github.com/unidoc/unioffice/schema/soo/wml"
)

func (s docHandler) BuildFrontpage() {
	para := s.doc.Paragraphs()[0]
	// Project title section
	{
		run := para.AddRun()
		para.SetStyle("Heading1")
		run.AddText(s.swagger.Info.Title)
		para = s.doc.AddParagraph()
	}

	// Project description section
	{
		para = s.doc.AddParagraph()
		run := para.AddRun()
		run.AddText(s.swagger.Info.Description)
	}

	// Project metadata section
	{
		para = s.doc.AddParagraph()

		table := s.doc.AddTable()
		borders := table.Properties().Borders()
		borders.SetAll(wml.ST_BorderSingle, color.Auto, 1*measurement.Point)

		row := table.AddRow()
		run := row.AddCell().AddParagraph().AddRun()
		run.Properties().SetBold(true)
		run.AddText("Version")
		row.AddCell().AddParagraph().AddRun().AddText(s.swagger.Info.Version)

		if s.swagger.Info.Contact != nil {
			row = table.AddRow()
			run = row.AddCell().AddParagraph().AddRun()
			run.Properties().SetBold(true)
			run.AddText("Contact Person")

			row.AddCell().AddParagraph().AddRun().AddText(s.swagger.Info.Contact.Name)

			row = table.AddRow()
			run = row.AddCell().AddParagraph().AddRun()
			run.Properties().SetBold(true)
			run.AddText("Contact URL")
			row.AddCell().AddParagraph().AddRun().AddText(s.swagger.Info.Contact.URL)

			row = table.AddRow()
			run = row.AddCell().AddParagraph().AddRun()
			run.Properties().SetBold(true)
			run.AddText("Contact Email")
			row.AddCell().AddParagraph().AddRun().AddText(s.swagger.Info.Contact.Email)

		}
	}

	// Servers section
	{
		para = s.doc.AddParagraph()
		run := para.AddRun()
		para.SetStyle("Heading2")
		run.AddText("Servers")

		para = s.doc.AddParagraph()
		table := s.doc.AddTable()
		borders := table.Properties().Borders()
		borders.SetAll(wml.ST_BorderSingle, color.Auto, 1*measurement.Point)

		for _, server := range s.swagger.Servers {
			row := table.AddRow()
			cell := row.AddCell()
			cell.Properties().SetVerticalMerge(wml.ST_MergeRestart)

			cellPara := cell.AddParagraph()
			cellPara.Properties().SetAlignment(wml.ST_JcLeft)

			serverName := server.Description
			if serverName == "" {
				serverName = "Default server"
			}

			paragraph := cellPara.AddRun()
			paragraph.Properties().SetBold(true)
			paragraph.AddText(serverName)

			row.AddCell().AddParagraph().AddRun().AddText(server.URL)
		}
	}

	// Security schemes section
	{
		para = s.doc.AddParagraph()
		run := para.AddRun()
		para.SetStyle("Heading2")
		run.AddText("Security Schemes")

		para = s.doc.AddParagraph()
		table := s.doc.AddTable()
		borders := table.Properties().Borders()
		borders.SetAll(wml.ST_BorderSingle, color.Auto, 1*measurement.Point)

		headerRow := table.AddRow()
		cell := headerRow.AddCell()
		cellPara := cell.AddParagraph()
		cellPara.Properties().SetAlignment(wml.ST_JcCenter)
		run = cellPara.AddRun()
		run.Properties().SetBold(true)
		run.AddText("Parameter Name")

		cell = headerRow.AddCell()
		cellPara = cell.AddParagraph()
		cellPara.Properties().SetAlignment(wml.ST_JcCenter)
		run = cellPara.AddRun()
		run.Properties().SetBold(true)
		run.AddText("Location")

		cell = headerRow.AddCell()
		cellPara = cell.AddParagraph()
		cellPara.Properties().SetAlignment(wml.ST_JcCenter)
		run = cellPara.AddRun()
		run.Properties().SetBold(true)
		run.AddText("Description")

		for key := range s.swagger.Components.SecuritySchemes {
			securityScheme := s.swagger.Components.SecuritySchemes[key]

			row := table.AddRow()
			cell := row.AddCell()
			paragraph := cell.AddParagraph().AddRun()
			paragraph.Properties().SetBold(true)
			paragraph.AddText(key)

			cell = row.AddCell()
			cell.AddParagraph().AddRun().AddText(securityScheme.Value.In)

			cell = row.AddCell()
			cell.AddParagraph().AddRun().AddText(securityScheme.Value.Name)
		}
	}

	para = s.doc.AddParagraph()
	para.AddRun().AddPageBreak()
}
