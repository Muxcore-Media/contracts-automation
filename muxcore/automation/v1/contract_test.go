package automationv1

import (
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func TestAutomationServiceRPCNamesStable(t *testing.T) {
	want := map[string]bool{
		"SearchItem":           true,
		"Dispatch":             true,
		"AddToQueue":           true,
		"GetQueue":             true,
		"UpdateQueueItem":      true,
		"RemoveFromQueue":      true,
		"GetHistory":           true,
		"SearchNow":            true,
		"ListBlocklist":        true,
		"ClearBlocklist":       true,
		"BlocklistRelease":     true,
		"RetryImport":          true,
		"ListDelayProfiles":    true,
		"UpsertDelayProfile":   true,
		"ListCutoffUnmet":      true,
		"ListSeriesOverrides":  true,
		"UpsertSeriesOverride": true,
		"DeleteSeriesOverride": true,
		"GetCapabilities":      true,
	}
	got := make(map[string]bool, len(AutomationService_ServiceDesc.Methods))
	for _, m := range AutomationService_ServiceDesc.Methods {
		got[m.MethodName] = true
	}
	for name := range want {
		if !got[name] {
			t.Errorf("missing RPC %s", name)
		}
	}
}

func TestQueueItemFieldNumbers(t *testing.T) {
	want := map[string]protoreflect.FieldNumber{
		"id":                 1,
		"item_type":          2,
		"item_id":            3,
		"tmdb_id":            4,
		"title":              5,
		"year":               6,
		"season_number":      7,
		"episode_number":     8,
		"monitored":          9,
		"missing":            10,
		"last_searched":      11,
		"created_at":         12,
		"updated_at":         13,
		"quality_profile_id": 14,
		"series_id":          15,
		"absolute_number":    16,
		"series_type":        17,
	}
	assertFieldNumbers(t, (&QueueItem{}).ProtoReflect().Descriptor(), want)
}

func TestReleaseMatchRejectionFields(t *testing.T) {
	want := map[string]protoreflect.FieldNumber{
		"guid":              1,
		"title":             2,
		"size":              3,
		"seeders":           4,
		"peers":             5,
		"indexer_name":      6,
		"download_url":      7,
		"info_url":          8,
		"download_protocol": 9,
		"score":             10,
		"category":          11,
		"sub_category":      12,
		"rejected":          13,
		"rejection_reason":  14,
		"format_score":      15,
	}
	assertFieldNumbers(t, (&ReleaseMatch{}).ProtoReflect().Descriptor(), want)
}

func TestGetQueueRequestFilters(t *testing.T) {
	want := map[string]protoreflect.FieldNumber{
		"page":      1,
		"page_size": 2,
		"filter":    3,
		"missing":   4,
		"monitored": 5,
	}
	assertFieldNumbers(t, (&GetQueueRequest{}).ProtoReflect().Descriptor(), want)
}

func TestGetHistoryRequestFilters(t *testing.T) {
	want := map[string]protoreflect.FieldNumber{
		"page":           1,
		"page_size":      2,
		"status":         3,
		"wanted_item_id": 4,
	}
	assertFieldNumbers(t, (&GetHistoryRequest{}).ProtoReflect().Descriptor(), want)
}

func TestSearchNowRequestTargeting(t *testing.T) {
	want := map[string]protoreflect.FieldNumber{
		"queue_id":  1,
		"item_type": 2,
		"item_id":   3,
	}
	assertFieldNumbers(t, (&SearchNowRequest{}).ProtoReflect().Descriptor(), want)
}

func TestSeriesOverrideFieldNumbers(t *testing.T) {
	want := map[string]protoreflect.FieldNumber{
		"series_id":        1,
		"delay_minutes":    2,
		"preferred_groups": 3,
		"ignored_groups":   4,
	}
	assertFieldNumbers(t, (&SeriesOverride{}).ProtoReflect().Descriptor(), want)
}

func TestGetCapabilitiesResponseFields(t *testing.T) {
	want := map[string]protoreflect.FieldNumber{
		"supported_item_types":    1,
		"supports_delay_profiles": 2,
		"supports_cutoff":         3,
		"supports_blocklist":      4,
		"supports_anime_absolute": 5,
		"supported_protocols":     6,
	}
	assertFieldNumbers(t, (&GetCapabilitiesResponse{}).ProtoReflect().Descriptor(), want)
}

func TestImportPath(t *testing.T) {
	var _ grpc.ServiceDesc = AutomationService_ServiceDesc
}

func TestCapabilityIDs(t *testing.T) {
	const consumerCapability = "automation"
	const contractCapability = "contracts.automation"
	if consumerCapability == contractCapability {
		t.Fatal("consumer and contract capability ids must differ")
	}
}

func assertFieldNumbers(t *testing.T, desc protoreflect.MessageDescriptor, want map[string]protoreflect.FieldNumber) {
	t.Helper()
	for name, num := range want {
		fd := desc.Fields().ByName(protoreflect.Name(name))
		if fd == nil {
			t.Errorf("missing field %q", name)
			continue
		}
		if fd.Number() != num {
			t.Errorf("field %q number = %d, want %d", name, fd.Number(), num)
		}
	}
}
