package aws

import (
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccAWSOpsworksCustomLayer_importBasic(t *testing.T) {
	name := acctest.RandString(10)

	resourceName := "aws_opsworks_custom_layer.tf-acc"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAwsOpsworksCustomLayerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAwsOpsworksCustomLayerConfigVpcCreate(name),
			},

			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
