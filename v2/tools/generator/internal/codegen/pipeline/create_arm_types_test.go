/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/test"
)

// We don't need to test everything here because a lot of the common cases are covered in the existing golden
// files tests. We focus on complicated edge cases here, such as flattening. Because of the difficulty in setting these
// things up, we combine testing of the CreateARMTypes and the ApplyARMConversionInterface stages

func TestCreateFlattenedARMType_CreatesExpectedConversions(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Define a test resource
	specProperties := test.CreateObjectDefinition(
		test.Pkg2020,
		"PersonProperties",
		test.FullNameProperty,
		test.FamilyNameProperty,
		test.KnownAsProperty)
	specPropertiesProp := astmodel.NewPropertyDefinition(
		"Properties",
		"properties",
		specProperties.Name()).SetFlatten(true).MakeTypeOptional()
	spec := test.CreateSpec(test.Pkg2020, "Person", specPropertiesProp, test.NameProperty)
	status := test.CreateStatus(test.Pkg2020, "Person")
	resource := test.CreateARMResource(test.Pkg2020, "Person", spec, status, test.Pkg2020APIVersion)

	defs := make(astmodel.TypeDefinitionSet)
	defs.AddAll(resource, status, spec, specProperties, test.Pkg2020APIVersion)

	idFactory := astmodel.NewIdentifierFactory()

	createARMTypes := CreateARMTypes(idFactory)
	applyARMConversionInterface := ApplyARMConversionInterface(idFactory)
	flatten := FlattenProperties()
	simplify := SimplifyDefinitions()
	strip := StripUnreferencedTypeDefinitions()

	state, err := RunTestPipeline(
		NewState().WithDefinitions(defs),
		createARMTypes,
		applyARMConversionInterface,
		flatten,
		simplify,
		strip)
	g.Expect(err).ToNot(HaveOccurred())

	test.AssertPackagesGenerateExpectedCode(t, state.Definitions())
}

func TestCreateFlattenedARMTypeWithResourceRef_CreatesExpectedConversions(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Define a test resource
	specProperties := test.CreateObjectDefinition(
		test.Pkg2020,
		"PersonProperties",
		test.FullNameProperty,
		test.FamilyNameProperty,
		test.KnownAsProperty)
	specPropertiesProp := astmodel.NewPropertyDefinition(
		"Properties",
		"properties",
		specProperties.Name()).SetFlatten(true).MakeTypeOptional()
	spec := test.CreateSpec(test.Pkg2020, "Person", specPropertiesProp, test.NameProperty)
	status := test.CreateStatus(test.Pkg2020, "Person")
	resource := test.CreateARMResource(test.Pkg2020, "Person", spec, status, test.Pkg2020APIVersion)

	defs := make(astmodel.TypeDefinitionSet)
	defs.AddAll(resource, status, spec, specProperties, test.Pkg2020APIVersion)

	idFactory := astmodel.NewIdentifierFactory()
	omc := config.NewObjectModelConfiguration()
	g.Expect(
		omc.ModifyProperty(
			specProperties.Name(),
			test.FamilyNameProperty.PropertyName(),
			func(propertyConfiguration *config.PropertyConfiguration) error {
				propertyConfiguration.SetARMReference(true)
				return nil
			})).To(Succeed())

	configuration := config.NewConfiguration()
	configuration.ObjectModelConfiguration = omc

	crossResourceRefs := AddCrossResourceReferences(configuration, idFactory)
	createARMTypes := CreateARMTypes(idFactory)
	applyARMConversionInterface := ApplyARMConversionInterface(idFactory)
	flatten := FlattenProperties()
	simplify := SimplifyDefinitions()
	strip := StripUnreferencedTypeDefinitions()

	state, err := RunTestPipeline(
		NewState().WithDefinitions(defs),
		crossResourceRefs,
		createARMTypes,
		applyARMConversionInterface,
		flatten,
		simplify,
		strip)
	g.Expect(err).ToNot(HaveOccurred())

	test.AssertPackagesGenerateExpectedCode(t, state.Definitions())
}

func TestCreateFlattenedARMTypeWithConfigMap_CreatesExpectedConversions(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Define a test resource
	specProperties := test.CreateObjectDefinition(
		test.Pkg2020,
		"PersonProperties",
		test.FullNameProperty,
		test.FamilyNameProperty,
		test.KnownAsProperty)
	specPropertiesProp := astmodel.NewPropertyDefinition(
		"Properties",
		"properties",
		specProperties.Name()).SetFlatten(true).MakeTypeOptional()
	spec := test.CreateSpec(test.Pkg2020, "Person", specPropertiesProp, test.NameProperty)
	status := test.CreateStatus(test.Pkg2020, "Person")
	resource := test.CreateARMResource(test.Pkg2020, "Person", spec, status, test.Pkg2020APIVersion)

	defs := make(astmodel.TypeDefinitionSet)
	defs.AddAll(resource, status, spec, specProperties, test.Pkg2020APIVersion)

	idFactory := astmodel.NewIdentifierFactory()
	omc := config.NewObjectModelConfiguration()
	g.Expect(
		omc.ModifyProperty(
			specProperties.Name(),
			test.FullNameProperty.PropertyName(),
			func(pc *config.PropertyConfiguration) error {
				pc.SetImportConfigMapMode(config.ImportConfigMapModeRequired)
				return nil
			})).
		To(Succeed())
	g.Expect(
		omc.ModifyProperty(
			specProperties.Name(),
			test.FamilyNameProperty.PropertyName(),
			func(pc *config.PropertyConfiguration) error {
				pc.SetImportConfigMapMode(config.ImportConfigMapModeOptional)
				return nil
			})).
		To(Succeed())

	configuration := config.NewConfiguration()
	configuration.ObjectModelConfiguration = omc

	addConfigMaps := AddConfigMaps(configuration)
	createARMTypes := CreateARMTypes(idFactory)
	applyARMConversionInterface := ApplyARMConversionInterface(idFactory)
	flatten := FlattenProperties()
	simplify := SimplifyDefinitions()
	strip := StripUnreferencedTypeDefinitions()

	state, err := RunTestPipeline(
		NewState().WithDefinitions(defs),
		addConfigMaps,
		createARMTypes,
		applyARMConversionInterface,
		flatten,
		simplify,
		strip)
	g.Expect(err).ToNot(HaveOccurred())

	test.AssertPackagesGenerateExpectedCode(t, state.Definitions())
}

func TestCreateARMTypeWithConfigMap_CreatesExpectedConversions(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Define a test resource
	specProperties := test.CreateObjectDefinition(
		test.Pkg2020,
		"PersonProperties",
		test.FullNameProperty,
		test.FamilyNameProperty,
		test.KnownAsProperty)
	specPropertiesProp := astmodel.NewPropertyDefinition(
		"Properties",
		"properties",
		specProperties.Name()).MakeTypeOptional()
	spec := test.CreateSpec(test.Pkg2020, "Person", specPropertiesProp, test.NameProperty)
	status := test.CreateStatus(test.Pkg2020, "Person")
	resource := test.CreateARMResource(test.Pkg2020, "Person", spec, status, test.Pkg2020APIVersion)

	defs := make(astmodel.TypeDefinitionSet)
	defs.AddAll(resource, status, spec, specProperties, test.Pkg2020APIVersion)

	idFactory := astmodel.NewIdentifierFactory()
	omc := config.NewObjectModelConfiguration()
	g.Expect(
		omc.ModifyProperty(
			specProperties.Name(),
			test.FullNameProperty.PropertyName(),
			func(pc *config.PropertyConfiguration) error {
				pc.SetImportConfigMapMode(config.ImportConfigMapModeRequired)
				return nil
			})).
		To(Succeed())
	g.Expect(
		omc.ModifyProperty(
			specProperties.Name(),
			test.FamilyNameProperty.PropertyName(),
			func(pc *config.PropertyConfiguration) error {
				pc.SetImportConfigMapMode(config.ImportConfigMapModeOptional)
				return nil
			})).
		To(Succeed())

	configuration := config.NewConfiguration()
	configuration.ObjectModelConfiguration = omc

	addConfigMaps := AddConfigMaps(configuration)
	createARMTypes := CreateARMTypes(idFactory)
	applyARMConversionInterface := ApplyARMConversionInterface(idFactory)
	simplify := SimplifyDefinitions()
	strip := StripUnreferencedTypeDefinitions()

	state, err := RunTestPipeline(
		NewState().WithDefinitions(defs),
		addConfigMaps,
		createARMTypes,
		applyARMConversionInterface,
		simplify,
		strip)
	g.Expect(err).ToNot(HaveOccurred())

	test.AssertPackagesGenerateExpectedCode(t, state.Definitions())
}
