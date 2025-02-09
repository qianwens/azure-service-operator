// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20211101storage

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

func Test_NamespacesAuthorizationRule_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NamespacesAuthorizationRule via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespacesAuthorizationRule, NamespacesAuthorizationRuleGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespacesAuthorizationRule runs a test to see if a specific instance of NamespacesAuthorizationRule round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespacesAuthorizationRule(subject NamespacesAuthorizationRule) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NamespacesAuthorizationRule
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

// Generator of NamespacesAuthorizationRule instances for property testing - lazily instantiated by
// NamespacesAuthorizationRuleGenerator()
var namespacesAuthorizationRuleGenerator gopter.Gen

// NamespacesAuthorizationRuleGenerator returns a generator of NamespacesAuthorizationRule instances for property testing.
func NamespacesAuthorizationRuleGenerator() gopter.Gen {
	if namespacesAuthorizationRuleGenerator != nil {
		return namespacesAuthorizationRuleGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForNamespacesAuthorizationRule(generators)
	namespacesAuthorizationRuleGenerator = gen.Struct(reflect.TypeOf(NamespacesAuthorizationRule{}), generators)

	return namespacesAuthorizationRuleGenerator
}

// AddRelatedPropertyGeneratorsForNamespacesAuthorizationRule is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespacesAuthorizationRule(gens map[string]gopter.Gen) {
	gens["Spec"] = Namespaces_AuthorizationRule_SpecGenerator()
	gens["Status"] = AuthorizationRule_STATUSGenerator()
}

func Test_AuthorizationRule_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AuthorizationRule_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAuthorizationRule_STATUS, AuthorizationRule_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAuthorizationRule_STATUS runs a test to see if a specific instance of AuthorizationRule_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForAuthorizationRule_STATUS(subject AuthorizationRule_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AuthorizationRule_STATUS
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

// Generator of AuthorizationRule_STATUS instances for property testing - lazily instantiated by
// AuthorizationRule_STATUSGenerator()
var authorizationRule_STATUSGenerator gopter.Gen

// AuthorizationRule_STATUSGenerator returns a generator of AuthorizationRule_STATUS instances for property testing.
// We first initialize authorizationRule_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func AuthorizationRule_STATUSGenerator() gopter.Gen {
	if authorizationRule_STATUSGenerator != nil {
		return authorizationRule_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAuthorizationRule_STATUS(generators)
	authorizationRule_STATUSGenerator = gen.Struct(reflect.TypeOf(AuthorizationRule_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAuthorizationRule_STATUS(generators)
	AddRelatedPropertyGeneratorsForAuthorizationRule_STATUS(generators)
	authorizationRule_STATUSGenerator = gen.Struct(reflect.TypeOf(AuthorizationRule_STATUS{}), generators)

	return authorizationRule_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForAuthorizationRule_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAuthorizationRule_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Rights"] = gen.SliceOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForAuthorizationRule_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForAuthorizationRule_STATUS(gens map[string]gopter.Gen) {
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSGenerator())
}

func Test_Namespaces_AuthorizationRule_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Namespaces_AuthorizationRule_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespaces_AuthorizationRule_Spec, Namespaces_AuthorizationRule_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespaces_AuthorizationRule_Spec runs a test to see if a specific instance of Namespaces_AuthorizationRule_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespaces_AuthorizationRule_Spec(subject Namespaces_AuthorizationRule_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Namespaces_AuthorizationRule_Spec
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

// Generator of Namespaces_AuthorizationRule_Spec instances for property testing - lazily instantiated by
// Namespaces_AuthorizationRule_SpecGenerator()
var namespaces_AuthorizationRule_SpecGenerator gopter.Gen

// Namespaces_AuthorizationRule_SpecGenerator returns a generator of Namespaces_AuthorizationRule_Spec instances for property testing.
func Namespaces_AuthorizationRule_SpecGenerator() gopter.Gen {
	if namespaces_AuthorizationRule_SpecGenerator != nil {
		return namespaces_AuthorizationRule_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespaces_AuthorizationRule_Spec(generators)
	namespaces_AuthorizationRule_SpecGenerator = gen.Struct(reflect.TypeOf(Namespaces_AuthorizationRule_Spec{}), generators)

	return namespaces_AuthorizationRule_SpecGenerator
}

// AddIndependentPropertyGeneratorsForNamespaces_AuthorizationRule_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespaces_AuthorizationRule_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["Rights"] = gen.SliceOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}
