package moduledoc

import (
	"fmt"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/unidoc/unioffice/color"
	"github.com/unidoc/unioffice/document"
	"github.com/unidoc/unioffice/measurement"
	"github.com/unidoc/unioffice/schema/soo/wml"
	"github.com/weisurya/go-openapi-converter/utils"
)

func (s docHandler) BuildPaths() {
	for path := range s.swagger.Paths {
		pathItems := s.swagger.Paths.Find(path)

		if pathItems.Get != nil {
			buildSpec(s.doc, s.swagger, pathItems.Get, path, MethodGet)
		}

		if pathItems.Post != nil {
			buildSpec(s.doc, s.swagger, pathItems.Post, path, MethodPost)
		}

		if pathItems.Put != nil {
			buildSpec(s.doc, s.swagger, pathItems.Put, path, MethodPut)
		}

		if pathItems.Patch != nil {
			buildSpec(s.doc, s.swagger, pathItems.Patch, path, MethodPatch)
		}

		if pathItems.Delete != nil {
			buildSpec(s.doc, s.swagger, pathItems.Delete, path, MethodDelete)
		}

		if pathItems.Head != nil {
			buildSpec(s.doc, s.swagger, pathItems.Head, path, MethodHead)
		}

		if pathItems.Options != nil {
			buildSpec(s.doc, s.swagger, pathItems.Options, path, MethodOptions)
		}

		if pathItems.Trace != nil {
			buildSpec(s.doc, s.swagger, pathItems.Trace, path, MethodTrace)
		}
	}
}

func buildSpec(doc *document.Document, swagger *openapi3.Swagger, endpoint *openapi3.Operation, path string, method string) {
	// Title section
	{
		para := doc.AddParagraph()
		run := para.AddRun()
		para.SetStyle("Heading2")
		run.AddText(endpoint.Summary)

		para = doc.AddParagraph()
		run = para.AddRun()
		run.AddText(fmt.Sprintf("%s - %s%s", method, swagger.Servers[0].URL, path))
	}

	// Description section
	{
		para := doc.AddParagraph()
		run := para.AddRun()
		para.SetStyle("Heading3")
		run.AddText("Description")

		if endpoint.Description == "" {
			endpoint.Description = "No description"
		}

		descriptions := strings.Split(endpoint.Description, "\n")

		for _, description := range descriptions {

			para = doc.AddParagraph()
			run = para.AddRun()
			run.AddText(description)
		}

	}

	// Tags section
	{
		para := doc.AddParagraph()
		run := para.AddRun()
		para.SetStyle("Heading3")
		run.AddText("Tags")

		for _, tag := range endpoint.Tags {
			para = doc.AddParagraph()
			run = para.AddRun()
			run.AddText(fmt.Sprintf("- %s", tag))
		}

	}

	// Authorization section
	{
		para := doc.AddParagraph()
		run := para.AddRun()
		para.SetStyle("Heading3")
		run.AddText("Authorization")

		table := doc.AddTable()
		borders := table.Properties().Borders()
		borders.SetAll(wml.ST_BorderSingle, color.Auto, 1*measurement.Point)

		if endpoint.Security != nil {

			for _, security := range *endpoint.Security {

				securityKey := utils.ExtractSecurityRequirementKey(security)
				securityDetail := swagger.Components.SecuritySchemes[securityKey]

				row := table.AddRow()
				cell := row.AddCell()
				cell.Properties().SetVerticalMerge(wml.ST_MergeRestart)

				cellPara := cell.AddParagraph()
				cellPara.Properties().SetAlignment(wml.ST_JcCenter)

				paragraph := cellPara.AddRun()
				paragraph.Properties().SetBold(true)
				paragraph.AddText(securityKey)

				row.AddCell().AddParagraph().AddRun().AddText(securityDetail.Value.Name)
			}
		}
	}

	// Header parameters section(
	{
		isExist := utils.IsParameterExist(endpoint.Parameters, utils.ParamHeader)

		if isExist != false {

			para := doc.AddParagraph()
			run := para.AddRun()
			para.SetStyle("Heading3")
			run.AddText("Header Parameters")

			table := doc.AddTable()
			// 4 inches wide
			borders := table.Properties().Borders()
			// thin borders
			borders.SetAll(wml.ST_BorderSingle, color.Auto, 1*measurement.Point)

			for _, param := range endpoint.Parameters {

				if param.Value.In == "header" {

					row := table.AddRow()
					cell := row.AddCell()
					cell.Properties().SetVerticalMerge(wml.ST_MergeRestart)

					cellPara := cell.AddParagraph()
					cellPara.Properties().SetAlignment(wml.ST_JcCenter)

					cellPara.AddRun().AddText(param.Value.Name)

					description := param.Value.Description
					if description == "" {

						description = "No description"
					}
					row.AddCell().AddParagraph().AddRun().AddText(description)

					row = table.AddRow()
					cell = row.AddCell()
					cell.Properties().SetVerticalMerge(wml.ST_MergeContinue)
					cell.AddParagraph().AddRun().AddText("")
					if param.Value.Required != true {

						row.AddCell().AddParagraph().AddRun().AddText("Optional")
					} else {

						row.AddCell().AddParagraph().AddRun().AddText("Required")
					}
				}
			}
		}

	}

	// Path parameters section
	{
		pathItems := swagger.Paths.Find(path)

		if len(pathItems.Parameters) > 0 {

			para := doc.AddParagraph()
			run := para.AddRun()
			para.SetStyle("Heading3")
			run.AddText("Path Parameters")

			table := doc.AddTable()
			borders := table.Properties().Borders()
			borders.SetAll(wml.ST_BorderSingle, color.Auto, 1*measurement.Point)

			for _, param := range pathItems.Parameters {

				row := table.AddRow()
				cell := row.AddCell()
				cell.Properties().SetVerticalMerge(wml.ST_MergeRestart)

				cellPara := cell.AddParagraph()
				cellPara.Properties().SetAlignment(wml.ST_JcCenter)
				cellPara.AddRun().AddText(param.Value.Name)

				description := param.Value.Description
				if description == "" {

					description = "No description"
				}
				row.AddCell().AddParagraph().AddRun().AddText(description)

				row = table.AddRow()
				cell = row.AddCell()
				cell.Properties().SetVerticalMerge(wml.ST_MergeContinue)
				cell.AddParagraph().AddRun().AddText("")
				if param.Value.Required != true {

					row.AddCell().AddParagraph().AddRun().AddText("Optional")
				} else {

					row.AddCell().AddParagraph().AddRun().AddText("Required")
				}
			}
		}
	}

	// Query parameters section
	{
		isExist := utils.IsParameterExist(endpoint.Parameters, utils.ParamQuery)

		if isExist != false {

			para := doc.AddParagraph()
			run := para.AddRun()
			para.SetStyle("Heading3")
			run.AddText("Query Parameters")

			table := doc.AddTable()
			// 4 inches wide
			borders := table.Properties().Borders()
			// thin borders
			borders.SetAll(wml.ST_BorderSingle, color.Auto, 1*measurement.Point)

			for _, param := range endpoint.Parameters {

				if param.Value.In == "query" {

					row := table.AddRow()
					cell := row.AddCell()
					cell.Properties().SetVerticalMerge(wml.ST_MergeRestart)

					cellPara := cell.AddParagraph()
					cellPara.Properties().SetAlignment(wml.ST_JcCenter)

					paragraph := cell.AddParagraph().AddRun()
					paragraph.Properties().SetBold(true)
					paragraph.AddText(param.Value.Name)

					description := param.Value.Description
					if description == "" {

						description = "No description"
					}
					row.AddCell().AddParagraph().AddRun().AddText(description)

					row = table.AddRow()
					cell = row.AddCell()
					cell.Properties().SetVerticalMerge(wml.ST_MergeContinue)
					cell.AddParagraph().AddRun().AddText("")
					if param.Value.Required != true {

						row.AddCell().AddParagraph().AddRun().AddText("Optional")
					} else {

						row.AddCell().AddParagraph().AddRun().AddText("Required")
					}
				}
			}
		}

	}

	// Request body schema & example section
	{
		if endpoint.RequestBody != nil {
			para := doc.AddParagraph()
			run := para.AddRun()
			para.SetStyle("Heading3")
			run.AddText("Request Body Schema")

			if endpoint.RequestBody.Ref != "" {
				component := swagger.Components.Schemas[endpoint.RequestBody.Ref].Value

				createBodyComponent(component, doc)

				para = doc.AddParagraph()

				createExampleComponent(swagger.Components.Examples, doc)

			} else if endpoint.RequestBody.Value.Content.Get(MimeTypeApplicationJSON) != nil {
				component := endpoint.RequestBody.Value.Content.Get(MimeTypeApplicationJSON)

				createBodyComponent(component.Schema.Value, doc)

				para = doc.AddParagraph()

				createExampleComponent(component.Examples, doc)
			}
		}
	}

	// Response body schema & example section
	{
		if len(endpoint.Responses) > 0 {
			para := doc.AddParagraph()
			run := para.AddRun()
			para.SetStyle("Heading3")

			for httpStatus := range endpoint.Responses {
				run.AddText("Response Body Schema - " + httpStatus)

				if endpoint.Responses[httpStatus].Ref != "" {
					component := swagger.Components.Schemas[endpoint.Responses[httpStatus].Ref].Value

					createBodyComponent(component, doc)

					para = doc.AddParagraph()

					createExampleComponent(swagger.Components.Examples, doc)

				} else if endpoint.Responses[httpStatus].Value.Content.Get(MimeTypeApplicationJSON) != nil {
					component := endpoint.Responses[httpStatus].Value.Content.Get(MimeTypeApplicationJSON)

					createBodyComponent(component.Schema.Value, doc)

					para = doc.AddParagraph()

					createExampleComponent(component.Examples, doc)
				}
			}
		}
	}
}
