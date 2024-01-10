package moduledoc

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/unidoc/unioffice/color"
	"github.com/unidoc/unioffice/document"
	"github.com/unidoc/unioffice/measurement"
	"github.com/unidoc/unioffice/schema/soo/wml"
	"github.com/weisurya/go-openapi-converter/utils"
)

func createBodyComponent(component *openapi3.Schema, doc *document.Document) {

	doc.AddParagraph()

	table := doc.AddTable()
	borders := table.Properties().Borders()
	borders.SetAll(wml.ST_BorderSingle, color.Auto, 1*measurement.Point)

	headerRow := table.AddRow()
	cell := headerRow.AddCell()
	cellPara := cell.AddParagraph()
	cellPara.Properties().SetAlignment(wml.ST_JcCenter)
	run := cellPara.AddRun()
	run.Properties().SetBold(true)
	run.AddText("Name")

	cell = headerRow.AddCell()
	cellPara = cell.AddParagraph()
	cellPara.Properties().SetAlignment(wml.ST_JcCenter)
	run = cellPara.AddRun()
	run.Properties().SetBold(true)
	run.AddText("Type")

	cell = headerRow.AddCell()
	cellPara = cell.AddParagraph()
	cellPara.Properties().SetAlignment(wml.ST_JcCenter)
	run = cellPara.AddRun()
	run.Properties().SetBold(true)
	run.AddText("Description")

	cell = headerRow.AddCell()
	cellPara = cell.AddParagraph()
	cellPara.Properties().SetAlignment(wml.ST_JcCenter)
	run = cellPara.AddRun()
	run.Properties().SetBold(true)
	run.AddText("Additional Spec.")

	// Check if the response schema is of the array type
	if component.Type == "array" {
		// Get the schema of the array items
		arrayItemsSchema := component.Items.Value

		// Process the array items schema
		processSchemaProperties(arrayItemsSchema.Properties, arrayItemsSchema.Required, table)
	} else {
		// Process the schema directly
		processSchemaProperties(component.Properties, component.Required, table)
	}
}

func processSchemaProperties(properties map[string]*openapi3.SchemaRef, required []string, table document.Table) {
	for fieldName := range properties {

		fieldData := properties[fieldName].Value

		isRequired := false
		if utils.IsDataExistInArray(required, fieldName) == true {
			isRequired = true
		}

		row := table.AddRow()

		cell := row.AddCell()
		paragraph := cell.AddParagraph().AddRun()
		paragraph.Properties().SetBold(true)
		paragraph.AddText(fieldName)

		cell = row.AddCell()
		cell.AddParagraph().AddRun().AddText(fieldData.Type)

		cell = row.AddCell()
		cell.AddParagraph().AddRun().AddText(fieldData.Description)

		cell = row.AddCell()
		cell.AddParagraph().AddRun().AddText(fmt.Sprintf("Required: %s", strconv.FormatBool(isRequired)))

		if fieldData.Default != nil {
			cell.AddParagraph().AddRun().AddText(fmt.Sprintf("Default: %v", fieldData.Default))
		}

		if len(fieldData.Enum) >= 1 {
			listEnum := ""
			for i, datum := range fieldData.Enum {
				if i != 0 {
					listEnum += ", "
				}
				listEnum += datum.(string)
			}

			cell.AddParagraph().AddRun().AddText(fmt.Sprintf("Enum: %s", listEnum))
		}

		if fieldData.Example != nil {
			cell.AddParagraph().AddRun().AddText(fmt.Sprintf("Example: %s", fieldData.Example.(string)))
		}

		if fieldData.Pattern != "" {
			cell.AddParagraph().AddRun().AddText(fmt.Sprintf("Pattern: %s", fieldData.Pattern))
		}

		if fieldData.MinLength != 0 {
			cell.AddParagraph().AddRun().AddText(fmt.Sprintf("Min. Length: %s", strconv.FormatUint(fieldData.MinLength, 10)))
		}

		if fieldData.MaxLength != nil {
			cell.AddParagraph().AddRun().AddText(fmt.Sprintf("Max. Length: %s", strconv.FormatUint(*fieldData.MaxLength, 10)))
		}
	}
}

func createExampleComponent(examples map[string]*openapi3.ExampleRef, doc *document.Document) {
	for key := range examples {
		_, exampleHashmap := utils.ExtractInterfaceOfMap(examples[key].Value.Value)

		example, _ := json.Marshal(exampleHashmap)

		para := doc.AddParagraph()
		run := para.AddRun()
		run.Properties().SetBold(true)
		run.AddText(key)

		table := doc.AddTable()
		borders := table.Properties().Borders()
		borders.SetAll(wml.ST_BorderSingle, color.Auto, 1*measurement.Point)

		row := table.AddRow()
		cell := row.AddCell()
		cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)

		para = cell.AddParagraph()
		run = para.AddRun()
		run.AddText(fmt.Sprintf("%s", string(example)))
	}
}
