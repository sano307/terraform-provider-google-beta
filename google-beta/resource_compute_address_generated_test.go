// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
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

package google

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccComputeAddress_addressBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeAddressDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeAddress_addressBasicExample(context),
			},
			{
				ResourceName:            "google_compute_address.ip_address",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"subnetwork", "region"},
			},
		},
	})
}

func testAccComputeAddress_addressBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_address" "ip_address" {
  name = "tf-test-my-address%{random_suffix}"
}
`, context)
}

func TestAccComputeAddress_addressWithSubnetworkExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeAddressDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeAddress_addressWithSubnetworkExample(context),
			},
			{
				ResourceName:            "google_compute_address.internal_with_subnet_and_address",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"subnetwork", "region"},
			},
		},
	})
}

func testAccComputeAddress_addressWithSubnetworkExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_network" "default" {
  name = "tf-test-my-network%{random_suffix}"
}

resource "google_compute_subnetwork" "default" {
  name          = "tf-test-my-subnet%{random_suffix}"
  ip_cidr_range = "10.0.0.0/16"
  region        = "us-central1"
  network       = google_compute_network.default.id
}

resource "google_compute_address" "internal_with_subnet_and_address" {
  name         = "tf-test-my-internal-address%{random_suffix}"
  subnetwork   = google_compute_subnetwork.default.id
  address_type = "INTERNAL"
  address      = "10.0.42.42"
  region       = "us-central1"
}
`, context)
}

func TestAccComputeAddress_addressWithGceEndpointExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeAddressDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeAddress_addressWithGceEndpointExample(context),
			},
			{
				ResourceName:            "google_compute_address.internal_with_gce_endpoint",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"subnetwork", "region"},
			},
		},
	})
}

func testAccComputeAddress_addressWithGceEndpointExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_address" "internal_with_gce_endpoint" {
  name         = "tf-test-my-internal-address-%{random_suffix}"
  address_type = "INTERNAL"
  purpose      = "GCE_ENDPOINT"
}
`, context)
}

func TestAccComputeAddress_addressWithSharedLoadbalancerVipExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckComputeAddressDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeAddress_addressWithSharedLoadbalancerVipExample(context),
			},
		},
	})
}

func testAccComputeAddress_addressWithSharedLoadbalancerVipExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_address" "internal_with_shared_loadbalancer_vip" {
  provider     = google-beta 
  name         = "tf-test-my-internal-address%{random_suffix}"
  address_type = "INTERNAL"
  purpose      = "SHARED_LOADBALANCER_VIP"
}
`, context)
}

func TestAccComputeAddress_instanceWithIpExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeAddressDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeAddress_instanceWithIpExample(context),
			},
			{
				ResourceName:            "google_compute_address.static",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"subnetwork", "region"},
			},
		},
	})
}

func testAccComputeAddress_instanceWithIpExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_address" "static" {
  name = "tf-test-ipv4-address%{random_suffix}"
}

data "google_compute_image" "debian_image" {
  family  = "debian-9"
  project = "debian-cloud"
}

resource "google_compute_instance" "instance_with_ip" {
  name         = "tf-test-vm-instance%{random_suffix}"
  machine_type = "f1-micro"
  zone         = "us-central1-a"

  boot_disk {
    initialize_params {
      image = data.google_compute_image.debian_image.self_link
    }
  }

  network_interface {
    network = "default"
    access_config {
      nat_ip = google_compute_address.static.address
    }
  }
}
`, context)
}

func testAccCheckComputeAddressDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_compute_address" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/addresses/{{name}}")
			if err != nil {
				return err
			}

			_, err = sendRequest(config, "GET", "", url, nil)
			if err == nil {
				return fmt.Errorf("ComputeAddress still exists at %s", url)
			}
		}

		return nil
	}
}
