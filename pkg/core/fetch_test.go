package core

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/cloudquery/cloudquery/internal/firebase"
	"github.com/cloudquery/cloudquery/pkg/core/database"
	"github.com/cloudquery/cloudquery/pkg/plugin"
	"github.com/stretchr/testify/assert"

	"github.com/cloudquery/cloudquery/pkg/plugin/registry"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"

	"github.com/cloudquery/cloudquery/pkg/config"
)

func Test_Fetch(t *testing.T) {
	latestVersion := getLatestVersion(t, "test")
	testCases := []struct {
		Name             string
		Options          FetchOptions
		ExpectedDiags    []diag.FlatDiag
		ExpectedResponse *FetchResponse
		Timeout          time.Duration
	}{
		{
			Name: "fetch-errors",
			Options: FetchOptions{
				ProvidersInfo: []ProviderInfo{
					{
						Provider: registry.Provider{
							Name:    "test",
							Version: registry.LatestVersion,
							Source:  registry.DefaultOrganization,
						},
						Config: &config.Provider{
							Name:          "test",
							Alias:         "test_alias",
							Resources:     []string{"slow_resource", "panic_resource", "error_resource", "very_slow_resource"},
							Env:           nil,
							Configuration: nil,
						},
					},
				},
			},
			ExpectedDiags: []diag.FlatDiag{
				{
					Err:         "error from provider",
					Resource:    "error_resource",
					Type:        diag.RESOLVING,
					Severity:    diag.ERROR,
					Summary:     "error from provider",
					Description: diag.Description{Resource: "error_resource", ResourceID: []string(nil), Summary: "error from provider", Detail: ""}},
				{
					Err:         "failed table panic_resource fetch. Error: resource with panic",
					Resource:    "panic_resource",
					Type:        diag.RESOLVING,
					Severity:    diag.ERROR,
					Summary:     "failed table panic_resource fetch. Error: resource with panic",
					Description: diag.Description{Resource: "panic_resource", ResourceID: []string(nil), Summary: "failed table panic_resource fetch. Error: resource with panic", Detail: ""},
				},
			},
			ExpectedResponse: &FetchResponse{ProviderFetchSummary: map[string]*ProviderFetchSummary{"test(test_alias)": {
				Name:                  "test",
				Alias:                 "test_alias",
				Version:               registry.LatestVersion,
				TotalResourcesFetched: 0,
				Status:                FetchFinished,
			}}},
		},
		{
			Name: "fetch-simple",
			Options: FetchOptions{
				ProvidersInfo: []ProviderInfo{
					{
						Provider: registry.Provider{
							Name:    "test",
							Version: registry.LatestVersion,
							Source:  registry.DefaultOrganization,
						},
						Config: &config.Provider{
							Name:          "test",
							Resources:     []string{"slow_resource", "very_slow_resource"},
							Env:           nil,
							Configuration: nil,
						},
					},
				},
			},
			ExpectedDiags: nil,
			ExpectedResponse: &FetchResponse{ProviderFetchSummary: map[string]*ProviderFetchSummary{"test": {
				Name:                  "test",
				Alias:                 "",
				Version:               registry.LatestVersion,
				TotalResourcesFetched: 0,
				Status:                FetchFinished,
			}}},
		},
		{
			Name: "fetch-timeout",
			Options: FetchOptions{
				ProvidersInfo: []ProviderInfo{
					{
						Provider: registry.Provider{
							Name:    "test",
							Version: registry.LatestVersion,
							Source:  registry.DefaultOrganization,
						},
						Config: &config.Provider{
							Name:          "test",
							Resources:     []string{"slow_resource", "very_slow_resource"},
							Env:           nil,
							Configuration: nil,
						},
					},
				},
			},
			Timeout: time.Second * 4,
			ExpectedDiags: []diag.FlatDiag{
				{
					Err:         "rpc error: code = DeadlineExceeded desc = context deadline exceeded",
					Type:        diag.USER,
					Severity:    diag.ERROR,
					Summary:     "provider fetch was canceled by user / fetch deadline exceeded: rpc error: code = DeadlineExceeded desc = context deadline exceeded",
					Description: diag.Description{Resource: "", ResourceID: []string(nil), Summary: "provider fetch was canceled by user / fetch deadline exceeded: rpc error: code = DeadlineExceeded desc = context deadline exceeded", Detail: ""}},
			},
			ExpectedResponse: &FetchResponse{ProviderFetchSummary: map[string]*ProviderFetchSummary{"test": {
				Name:                  "test",
				Alias:                 "",
				Version:               registry.LatestVersion,
				TotalResourcesFetched: 0,
				Status:                FetchCanceled,
			}}},
		},
		{
			Name: "fetch-default-config",
			Options: FetchOptions{
				ProvidersInfo: []ProviderInfo{
					{
						Provider: registry.Provider{
							Name:    "test",
							Version: registry.LatestVersion,
							Source:  registry.DefaultOrganization,
						},
						Config: &config.Provider{
							Name:      "test",
							Resources: []string{"slow_resource"},
						},
					},
				},
			},
			ExpectedResponse: &FetchResponse{ProviderFetchSummary: map[string]*ProviderFetchSummary{"test": {
				Name:                  "test",
				Alias:                 "",
				Version:               registry.LatestVersion,
				TotalResourcesFetched: 0,
				Status:                FetchFinished,
			}}},
		},
		{
			Name: "fetch-duplicates",
			Options: FetchOptions{
				ProvidersInfo: []ProviderInfo{
					{
						Provider: registry.Provider{
							Name:    "test",
							Version: registry.LatestVersion,
							Source:  registry.DefaultOrganization,
						},
						Config: &config.Provider{
							Name:          "test",
							Resources:     []string{"slow_resource", "slow_resource", "slow_resource"},
							Env:           nil,
							Configuration: nil,
						},
					},
				},
			},
			ExpectedDiags: []diag.FlatDiag{{Err: "resource \"slow_resource\" is duplicate", Type: 7, Severity: 2, Summary: "resource \"slow_resource\" is duplicate", Description: diag.Description{Summary: "resource \"slow_resource\" is duplicate", Detail: "configuration has duplicate resources"}}},
			ExpectedResponse: &FetchResponse{ProviderFetchSummary: map[string]*ProviderFetchSummary{"test": {
				Name:                  "test",
				Alias:                 "",
				Version:               registry.LatestVersion,
				TotalResourcesFetched: 0,
				Status:                FetchFailed,
			}}},
		},
	}

	pManager, err := plugin.NewManager(registry.NewRegistryHub(firebase.CloudQueryRegistryURL))
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			dsn := setupDB(t)
			storage := database.NewStorage(dsn, nil)
			rp := make([]registry.Provider, len(tc.Options.ProvidersInfo))
			for i, p := range tc.Options.ProvidersInfo {
				rp[i] = p.Provider
			}
			// download test provider if it doesn't already exist
			_, diags := Download(context.Background(), pManager, &DownloadOptions{
				Providers: rp,
				NoVerify:  false,
			})
			require.False(t, diags.HasDiags())

			for _, r := range rp {
				// Sync provider in table before fetch
				_, diags := Sync(context.Background(), storage, pManager, &SyncOptions{
					Provider:       r,
					DownloadLatest: false,
				})
				require.False(t, diags.HasDiags())
			}

			require.Nil(t, err)
			var (
				ctx    = context.Background()
				cancel context.CancelFunc
			)
			if tc.Timeout > 0 {
				ctx, cancel = context.WithTimeout(context.Background(), tc.Timeout)
				defer cancel()
			}
			resp, diags := Fetch(ctx, storage, pManager, &tc.Options)
			if tc.ExpectedDiags != nil {
				assert.ElementsMatch(t, tc.ExpectedDiags, diag.FlattenDiags(diags, false))
			} else {
				assert.Equal(t, []diag.FlatDiag{}, diag.FlattenDiags(diags, false))
			}
			if tc.ExpectedResponse == nil {
				require.Nil(t, resp)
			} else {
				for k, p := range tc.ExpectedResponse.ProviderFetchSummary {
					fetchSummary, ok := resp.ProviderFetchSummary[k]
					require.True(t, ok)
					assert.Equal(t, p.Name, fetchSummary.Name)
					if p.Version == registry.LatestVersion {
						assert.Equal(t, latestVersion, fetchSummary.Version)
					} else {
						assert.Equal(t, p.Version, fetchSummary.Version)
					}
					assert.Equal(t, p.Status, fetchSummary.Status)
					assert.Equal(t, p.TotalResourcesFetched, fetchSummary.TotalResourcesFetched)
				}
			}
		})
	}
}

func Test_doNormalizeResources(t *testing.T) {
	tests := []struct {
		name      string
		requested []string
		keepOrder bool
		all       map[string]*schema.Table
		want      []string
		wantErr   bool
	}{
		{
			"wilcard",
			[]string{"*"},
			false,
			map[string]*schema.Table{"3": nil, "2": nil, "1": nil},
			[]string{"1", "2", "3"},
			false,
		},
		{
			"wilcard with explicit",
			[]string{"*", "1"},
			false,
			map[string]*schema.Table{"3": nil, "2": nil, "1": nil},
			nil,
			true,
		},
		{
			"unknown resource",
			[]string{"1", "2", "x"},
			false,
			map[string]*schema.Table{"3": nil, "2": nil, "1": nil},
			nil,
			true,
		},
		{
			"duplicate resource",
			[]string{"1", "2", "1"},
			false,
			map[string]*schema.Table{"3": nil, "2": nil, "1": nil},
			nil,
			true,
		},
		{
			"ok, all explicit, keep order",
			[]string{"2", "1"},
			true,
			map[string]*schema.Table{"3": nil, "2": nil, "1": nil},
			[]string{"2", "1"},
			false,
		},
		{
			"ok, all explicit, shuffle but no dot",
			[]string{"2", "1"},
			false,
			map[string]*schema.Table{"3": nil, "2": nil, "1": nil},
			[]string{"2", "1"},
			false,
		},
		{
			"ok, all explicit, shuffle",
			[]string{"c1.a", "c1.b", "c2.c", "c3.d", "c3.f", "c3.g"},
			false,
			map[string]*schema.Table{"c1.a": nil, "c1.b": nil, "c2.c": nil, "c3.d": nil, "c3.e": nil, "c3.f": nil, "c3.g": nil},
			[]string{"c1.a", "c2.c", "c3.d", "c1.b", "c3.f", "c3.g"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, diags := doNormalizeResources(tt.requested, tt.all, tt.keepOrder)
			if tt.wantErr != diags.HasErrors() {
				t.Errorf("doInterpolate() error = %v, wantErr %v", diags, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("doInterpolate() = %v, want %v", got, tt.want)
			}
		})
	}
}
