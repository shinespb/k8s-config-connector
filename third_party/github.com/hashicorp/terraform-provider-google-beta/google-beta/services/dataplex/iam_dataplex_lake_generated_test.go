// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package dataplex_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
)

func TestAccDataplexLakeIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
		"role":          "roles/viewer",
		"project_name":  envvar.GetTestProjectFromEnv(),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataplexLakeIamBinding_basicGenerated(context),
			},
			{
				ResourceName:      "google_dataplex_lake_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/lakes/%s roles/viewer", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-lake%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccDataplexLakeIamBinding_updateGenerated(context),
			},
			{
				ResourceName:      "google_dataplex_lake_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/lakes/%s roles/viewer", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-lake%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccDataplexLakeIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
		"role":          "roles/viewer",
		"project_name":  envvar.GetTestProjectFromEnv(),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccDataplexLakeIamMember_basicGenerated(context),
			},
			{
				ResourceName:      "google_dataplex_lake_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/lakes/%s roles/viewer user:admin@hashicorptest.com", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-lake%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccDataplexLakeIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
		"role":          "roles/viewer",
		"project_name":  envvar.GetTestProjectFromEnv(),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataplexLakeIamPolicy_basicGenerated(context),
				Check:  resource.TestCheckResourceAttrSet("data.google_dataplex_lake_iam_policy.foo", "policy_data"),
			},
			{
				ResourceName:      "google_dataplex_lake_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/lakes/%s", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-lake%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccDataplexLakeIamPolicy_emptyBinding(context),
			},
			{
				ResourceName:      "google_dataplex_lake_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/lakes/%s", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-lake%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccDataplexLakeIamMember_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_dataplex_lake" "example" {
  location     = "us-central1"
  name         = "tf-test-lake%{random_suffix}"
  description  = "Test Lake"
  display_name = "Test Lake"

  labels = {
    my-lake = "exists"
  }

  project = "%{project_name}"
}

resource "google_dataplex_lake_iam_member" "foo" {
  project = google_dataplex_lake.example.project
  location = google_dataplex_lake.example.location
  lake = google_dataplex_lake.example.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccDataplexLakeIamPolicy_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_dataplex_lake" "example" {
  location     = "us-central1"
  name         = "tf-test-lake%{random_suffix}"
  description  = "Test Lake"
  display_name = "Test Lake"

  labels = {
    my-lake = "exists"
  }

  project = "%{project_name}"
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_dataplex_lake_iam_policy" "foo" {
  project = google_dataplex_lake.example.project
  location = google_dataplex_lake.example.location
  lake = google_dataplex_lake.example.name
  policy_data = data.google_iam_policy.foo.policy_data
}

data "google_dataplex_lake_iam_policy" "foo" {
  project = google_dataplex_lake.example.project
  location = google_dataplex_lake.example.location
  lake = google_dataplex_lake.example.name
  depends_on = [
    google_dataplex_lake_iam_policy.foo
  ]
}
`, context)
}

func testAccDataplexLakeIamPolicy_emptyBinding(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_dataplex_lake" "example" {
  location     = "us-central1"
  name         = "tf-test-lake%{random_suffix}"
  description  = "Test Lake"
  display_name = "Test Lake"

  labels = {
    my-lake = "exists"
  }

  project = "%{project_name}"
}

data "google_iam_policy" "foo" {
}

resource "google_dataplex_lake_iam_policy" "foo" {
  project = google_dataplex_lake.example.project
  location = google_dataplex_lake.example.location
  lake = google_dataplex_lake.example.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccDataplexLakeIamBinding_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_dataplex_lake" "example" {
  location     = "us-central1"
  name         = "tf-test-lake%{random_suffix}"
  description  = "Test Lake"
  display_name = "Test Lake"

  labels = {
    my-lake = "exists"
  }

  project = "%{project_name}"
}

resource "google_dataplex_lake_iam_binding" "foo" {
  project = google_dataplex_lake.example.project
  location = google_dataplex_lake.example.location
  lake = google_dataplex_lake.example.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccDataplexLakeIamBinding_updateGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_dataplex_lake" "example" {
  location     = "us-central1"
  name         = "tf-test-lake%{random_suffix}"
  description  = "Test Lake"
  display_name = "Test Lake"

  labels = {
    my-lake = "exists"
  }

  project = "%{project_name}"
}

resource "google_dataplex_lake_iam_binding" "foo" {
  project = google_dataplex_lake.example.project
  location = google_dataplex_lake.example.location
  lake = google_dataplex_lake.example.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:gterraformtest1@gmail.com"]
}
`, context)
}