// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20180901

import (
	"encoding/json"
	v20180901s "github.com/Azure/azure-service-operator/v2/api/network/v1beta20180901storage"
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

func Test_PrivateDnsZone_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from PrivateDnsZone to hub returns original",
		prop.ForAll(RunResourceConversionTestForPrivateDnsZone, PrivateDnsZoneGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForPrivateDnsZone tests if a specific instance of PrivateDnsZone round trips to the hub storage version and back losslessly
func RunResourceConversionTestForPrivateDnsZone(subject PrivateDnsZone) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v20180901s.PrivateDnsZone
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual PrivateDnsZone
	err = actual.ConvertFrom(&hub)
	if err != nil {
		return err.Error()
	}

	// Compare actual with what we started with
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_PrivateDnsZone_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from PrivateDnsZone to PrivateDnsZone via AssignProperties_To_PrivateDnsZone & AssignProperties_From_PrivateDnsZone returns original",
		prop.ForAll(RunPropertyAssignmentTestForPrivateDnsZone, PrivateDnsZoneGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForPrivateDnsZone tests if a specific instance of PrivateDnsZone can be assigned to v1beta20180901storage and back losslessly
func RunPropertyAssignmentTestForPrivateDnsZone(subject PrivateDnsZone) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20180901s.PrivateDnsZone
	err := copied.AssignProperties_To_PrivateDnsZone(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual PrivateDnsZone
	err = actual.AssignProperties_From_PrivateDnsZone(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_PrivateDnsZone_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateDnsZone via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateDnsZone, PrivateDnsZoneGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateDnsZone runs a test to see if a specific instance of PrivateDnsZone round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateDnsZone(subject PrivateDnsZone) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateDnsZone
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

// Generator of PrivateDnsZone instances for property testing - lazily instantiated by PrivateDnsZoneGenerator()
var privateDnsZoneGenerator gopter.Gen

// PrivateDnsZoneGenerator returns a generator of PrivateDnsZone instances for property testing.
func PrivateDnsZoneGenerator() gopter.Gen {
	if privateDnsZoneGenerator != nil {
		return privateDnsZoneGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForPrivateDnsZone(generators)
	privateDnsZoneGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZone{}), generators)

	return privateDnsZoneGenerator
}

// AddRelatedPropertyGeneratorsForPrivateDnsZone is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrivateDnsZone(gens map[string]gopter.Gen) {
	gens["Spec"] = PrivateDnsZone_SpecGenerator()
	gens["Status"] = PrivateZone_STATUSGenerator()
}

func Test_PrivateDnsZone_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from PrivateDnsZone_Spec to PrivateDnsZone_Spec via AssignProperties_To_PrivateDnsZone_Spec & AssignProperties_From_PrivateDnsZone_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForPrivateDnsZone_Spec, PrivateDnsZone_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForPrivateDnsZone_Spec tests if a specific instance of PrivateDnsZone_Spec can be assigned to v1beta20180901storage and back losslessly
func RunPropertyAssignmentTestForPrivateDnsZone_Spec(subject PrivateDnsZone_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20180901s.PrivateDnsZone_Spec
	err := copied.AssignProperties_To_PrivateDnsZone_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual PrivateDnsZone_Spec
	err = actual.AssignProperties_From_PrivateDnsZone_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_PrivateDnsZone_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateDnsZone_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateDnsZone_Spec, PrivateDnsZone_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateDnsZone_Spec runs a test to see if a specific instance of PrivateDnsZone_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateDnsZone_Spec(subject PrivateDnsZone_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateDnsZone_Spec
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

// Generator of PrivateDnsZone_Spec instances for property testing - lazily instantiated by
// PrivateDnsZone_SpecGenerator()
var privateDnsZone_SpecGenerator gopter.Gen

// PrivateDnsZone_SpecGenerator returns a generator of PrivateDnsZone_Spec instances for property testing.
func PrivateDnsZone_SpecGenerator() gopter.Gen {
	if privateDnsZone_SpecGenerator != nil {
		return privateDnsZone_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZone_Spec(generators)
	privateDnsZone_SpecGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZone_Spec{}), generators)

	return privateDnsZone_SpecGenerator
}

// AddIndependentPropertyGeneratorsForPrivateDnsZone_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateDnsZone_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

func Test_PrivateZone_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from PrivateZone_STATUS to PrivateZone_STATUS via AssignProperties_To_PrivateZone_STATUS & AssignProperties_From_PrivateZone_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForPrivateZone_STATUS, PrivateZone_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForPrivateZone_STATUS tests if a specific instance of PrivateZone_STATUS can be assigned to v1beta20180901storage and back losslessly
func RunPropertyAssignmentTestForPrivateZone_STATUS(subject PrivateZone_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20180901s.PrivateZone_STATUS
	err := copied.AssignProperties_To_PrivateZone_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual PrivateZone_STATUS
	err = actual.AssignProperties_From_PrivateZone_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_PrivateZone_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateZone_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateZone_STATUS, PrivateZone_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateZone_STATUS runs a test to see if a specific instance of PrivateZone_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateZone_STATUS(subject PrivateZone_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateZone_STATUS
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

// Generator of PrivateZone_STATUS instances for property testing - lazily instantiated by PrivateZone_STATUSGenerator()
var privateZone_STATUSGenerator gopter.Gen

// PrivateZone_STATUSGenerator returns a generator of PrivateZone_STATUS instances for property testing.
func PrivateZone_STATUSGenerator() gopter.Gen {
	if privateZone_STATUSGenerator != nil {
		return privateZone_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateZone_STATUS(generators)
	privateZone_STATUSGenerator = gen.Struct(reflect.TypeOf(PrivateZone_STATUS{}), generators)

	return privateZone_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForPrivateZone_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateZone_STATUS(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["MaxNumberOfRecordSets"] = gen.PtrOf(gen.Int())
	gens["MaxNumberOfVirtualNetworkLinks"] = gen.PtrOf(gen.Int())
	gens["MaxNumberOfVirtualNetworkLinksWithRegistration"] = gen.PtrOf(gen.Int())
	gens["NumberOfRecordSets"] = gen.PtrOf(gen.Int())
	gens["NumberOfVirtualNetworkLinks"] = gen.PtrOf(gen.Int())
	gens["NumberOfVirtualNetworkLinksWithRegistration"] = gen.PtrOf(gen.Int())
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		PrivateZoneProperties_ProvisioningState_STATUS_Canceled,
		PrivateZoneProperties_ProvisioningState_STATUS_Creating,
		PrivateZoneProperties_ProvisioningState_STATUS_Deleting,
		PrivateZoneProperties_ProvisioningState_STATUS_Failed,
		PrivateZoneProperties_ProvisioningState_STATUS_Succeeded,
		PrivateZoneProperties_ProvisioningState_STATUS_Updating))
}
