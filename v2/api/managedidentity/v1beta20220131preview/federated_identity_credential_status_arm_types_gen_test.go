// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20220131preview

import (
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/kr/pretty"
	"github.com/kylelemons/godebug/diff"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"os"
	"reflect"
	"testing"
)

func Test_FederatedIdentityCredential_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FederatedIdentityCredential_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFederatedIdentityCredential_STATUS_ARM, FederatedIdentityCredential_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFederatedIdentityCredential_STATUS_ARM runs a test to see if a specific instance of FederatedIdentityCredential_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForFederatedIdentityCredential_STATUS_ARM(subject FederatedIdentityCredential_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FederatedIdentityCredential_STATUS_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of FederatedIdentityCredential_STATUS_ARM instances for property testing - lazily instantiated by
// FederatedIdentityCredential_STATUS_ARMGenerator()
var federatedIdentityCredential_STATUS_ARMGenerator gopter.Gen

// FederatedIdentityCredential_STATUS_ARMGenerator returns a generator of FederatedIdentityCredential_STATUS_ARM instances for property testing.
// We first initialize federatedIdentityCredential_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func FederatedIdentityCredential_STATUS_ARMGenerator() gopter.Gen {
	if federatedIdentityCredential_STATUS_ARMGenerator != nil {
		return federatedIdentityCredential_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFederatedIdentityCredential_STATUS_ARM(generators)
	federatedIdentityCredential_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(FederatedIdentityCredential_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFederatedIdentityCredential_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForFederatedIdentityCredential_STATUS_ARM(generators)
	federatedIdentityCredential_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(FederatedIdentityCredential_STATUS_ARM{}), generators)

	return federatedIdentityCredential_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForFederatedIdentityCredential_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFederatedIdentityCredential_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForFederatedIdentityCredential_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFederatedIdentityCredential_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(FederatedIdentityCredentialProperties_STATUS_ARMGenerator())
}

func Test_FederatedIdentityCredentialProperties_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FederatedIdentityCredentialProperties_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFederatedIdentityCredentialProperties_STATUS_ARM, FederatedIdentityCredentialProperties_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFederatedIdentityCredentialProperties_STATUS_ARM runs a test to see if a specific instance of FederatedIdentityCredentialProperties_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForFederatedIdentityCredentialProperties_STATUS_ARM(subject FederatedIdentityCredentialProperties_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FederatedIdentityCredentialProperties_STATUS_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of FederatedIdentityCredentialProperties_STATUS_ARM instances for property testing - lazily instantiated by
// FederatedIdentityCredentialProperties_STATUS_ARMGenerator()
var federatedIdentityCredentialProperties_STATUS_ARMGenerator gopter.Gen

// FederatedIdentityCredentialProperties_STATUS_ARMGenerator returns a generator of FederatedIdentityCredentialProperties_STATUS_ARM instances for property testing.
func FederatedIdentityCredentialProperties_STATUS_ARMGenerator() gopter.Gen {
	if federatedIdentityCredentialProperties_STATUS_ARMGenerator != nil {
		return federatedIdentityCredentialProperties_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFederatedIdentityCredentialProperties_STATUS_ARM(generators)
	federatedIdentityCredentialProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(FederatedIdentityCredentialProperties_STATUS_ARM{}), generators)

	return federatedIdentityCredentialProperties_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForFederatedIdentityCredentialProperties_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFederatedIdentityCredentialProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Audiences"] = gen.SliceOf(gen.AlphaString())
	gens["Issuer"] = gen.PtrOf(gen.AlphaString())
	gens["Subject"] = gen.PtrOf(gen.AlphaString())
}
