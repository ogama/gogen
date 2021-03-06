package model

type ObjectModel struct {
	FieldName string
	Fields    []FieldModel
}

func (objectTemplate ObjectModel) Generate(context *GeneratorContext, request GenerationRequest) (result interface{}, err error) {
	generatedObject := make(map[string]interface{})
	for _, field := range objectTemplate.Fields {
		var generated interface{}
		if generated, err = field.Generate(context, request); err != nil {
			return result, err
		}
		generatedObject[field.FieldName] = generated
	}
	return generatedObject, err
}
