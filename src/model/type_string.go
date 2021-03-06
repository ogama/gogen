package model

import (
	"github.com/ogama/gogen/src/reggen"
)

type StringType struct {
	pattern         string
	stringGenerator *reggen.Generator
	name            string
}

func (s *StringType) Generate(context *GeneratorContext, _ GenerationRequest) (result interface{}, err error) {
	if s.stringGenerator == nil {
		var stringGenerator *reggen.Generator
		if stringGenerator, err = reggen.NewGenerator(s.pattern, context.Rand); err != nil {
			return result, err
		}
		s.stringGenerator = stringGenerator
	}
	return s.stringGenerator.Generate(1000), err
}

func (s *StringType) GetName() string {
	return s.name
}

type StringTypeFactory struct{}

func (s StringTypeFactory) DefaultOptions() TypeOptions {
	defaultOptions := TypeOptions{}
	defaultOptions.Add("pattern", "[A-Z]{1}[A-Za-z]{10,25}")
	defaultOptions.Add("name", "")
	return defaultOptions
}

func (s StringTypeFactory) New(parameters TypeFactoryParameter) (generator TypeGenerator, err error) {
	pattern := parameters.Options.GetOptionAsString("pattern")
	return &StringType{
		pattern: pattern,
		name:    parameters.Options.GetOptionAsString("name"),
	}, err
}
