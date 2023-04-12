package ssmcontacts_test

import (
	"context"
	"fmt"
	"testing"

	sdkacctest "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/names"
)

func testPlanDataSource_basic(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping long-running test in short mode")
	}

	ctx := context.Background()

	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	planDataSourceName := "data.aws_ssmcontacts_plan.test"
	planResourceName := "aws_ssmcontacts_plan.test"

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(ctx, t)
			testAccContactPreCheck(t)
		},
		ErrorCheck:               acctest.ErrorCheck(t, names.SSMContactsEndpointID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckPlanDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPlanDataSourceConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPlanExists(planDataSourceName),
					testAccCheckPlanExists(planResourceName),
					resource.TestCheckResourceAttrPair(
						planDataSourceName, "contact_id",
						planResourceName, "contact_id",
					),
					resource.TestCheckResourceAttrPair(
						planDataSourceName, "stage.#",
						planResourceName, "stage.#",
					),
					resource.TestCheckResourceAttrPair(
						planDataSourceName, "stage.0.duration_in_minutes",
						planResourceName, "stage.0.duration_in_minutes",
					),
					resource.TestCheckResourceAttrPair(
						planDataSourceName, "stage.0.target.#",
						planResourceName, "stage.0.target.#",
					),
					resource.TestCheckResourceAttrPair(
						planDataSourceName, "stage.0.target.0.contact_target_info.is_essential",
						planResourceName, "stage.0.target.0.contact_target_info.contact_id",
					),
					resource.TestCheckResourceAttrPair(
						planDataSourceName, "stage.0.target.1.contact_target_info.is_essential",
						planResourceName, "stage.0.target.1.contact_target_info.contact_id",
					),
					resource.TestCheckResourceAttrPair(
						planDataSourceName, "stage.1.duration_in_minutes",
						planResourceName, "stage.1.duration_in_minutes",
					),
					resource.TestCheckResourceAttrPair(
						planDataSourceName, "stage.1.target.#",
						planResourceName, "stage.1.target.#",
					),
					resource.TestCheckResourceAttrPair(
						planDataSourceName, "stage.1.target.0.contact_target_info.is_essential",
						planResourceName, "stage.1.target.0.contact_target_info.contact_id",
					),
					resource.TestCheckResourceAttrPair(
						planDataSourceName, "stage.1.target.1.contact_target_info.is_essential",
						planResourceName, "stage.1.target.1.contact_target_info.contact_id",
					),
				),
			},
		},
	})
}

func testPlanDataSource_channelTargetInfo(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping long-running test in short mode")
	}

	ctx := context.Background()

	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	planDataSourceName := "data.aws_ssmcontacts_plan.test"
	planResourceName := "aws_ssmcontacts_plan.test"

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(ctx, t)
			testAccContactPreCheck(t)
		},
		ErrorCheck:               acctest.ErrorCheck(t, names.SSMContactsEndpointID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckPlanDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPlanDataSourceConfig_channelTargetInfo(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPlanExists(planDataSourceName),
					resource.TestCheckResourceAttrPair(planDataSourceName, "stage.0.target.#", planResourceName, "stage.0.target.#"),
					resource.TestCheckResourceAttrPair(
						planDataSourceName,
						"stage.0.target.0.channel_target_info.0.contact_channel_id",
						planResourceName,
						"stage.0.target.0.channel_target_info.0.contact_channel_id",
					),
					resource.TestCheckResourceAttrPair(
						planDataSourceName,
						"stage.0.target.0.contact_target_info.0.retry_interval_in_minutes",
						planResourceName,
						"stage.0.target.0.contact_target_info.0.retry_interval_in_minutes",
					),
				),
			},
		},
	})
}

func testAccPlanDataSourceConfig_basic(rName string) string {
	return acctest.ConfigCompose(
		testAccPlanDataSourceConfig_base(rName),
		fmt.Sprintf(`
resource "aws_ssmcontacts_plan" "test" {
	contact_id = aws_ssmcontacts_contact.test_escalation_plan_one.arn
	stage {
		duration_in_minutes = 1
		target {
			contact_target_info {
				is_essential = false
				contact_id = aws_ssmcontacts_contact.test_contact_one.arn
			}
		}
		target {
			contact_target_info {
				is_essential = true
				contact_id = aws_ssmcontacts_contact.test_contact_two.arn
			}
		}
	}
	stage {
		duration_in_minutes = 0
		target {
			contact_target_info {
				is_essential = false
				contact_id = aws_ssmcontacts_contact.test_contact_three.arn
			}
		}
		target {
			contact_target_info {
				is_essential = true
				contact_id = aws_ssmcontacts_contact.test_contact_four.arn
			}
		}
	}
}

data "aws_ssmcontacts_plan" "test" {
	contact_id = aws_ssmcontacts_contact.test_escalation_plan_one.arn

	depends_on = [aws_ssmcontacts_plan.test]
}
`))
}

func testAccPlanDataSourceConfig_channelTargetInfo(rName string) string {
	return acctest.ConfigCompose(
		testAccPlanDataSourceConfig_base(rName),
		fmt.Sprintf(`
resource "aws_ssmcontacts_contact_channel" "test" {
	contact_id = aws_ssmcontacts_contact.test_contact_one.arn
	delivery_address {
		simple_address = "email@example.com"
	}
	name = "Test Contact Channel for %[1]s"
	type = "EMAIL"
}

resource "aws_ssmcontacts_plan" "test" {
	contact_id = aws_ssmcontacts_contact.test_contact_one.arn
	stage {
		duration_in_minutes = 1
		target {
			channel_target_info {
				contact_channel_id = aws_ssmcontacts_contact_channel.test.arn
				retry_interval_in_minutes = 1
			}
		}
	}
}

data "aws_ssmcontacts_plan" "test" {
	contact_id = aws_ssmcontacts_contact.test_contact_one.arn

	depends_on = [aws_ssmcontacts_plan.test]
}
`, rName))
}

func testAccPlanDataSourceConfig_base(alias string) string {
	return fmt.Sprintf(`
resource "aws_ssmincidents_replication_set" "test" {
	region {
		name = %[1]q
	}
}

resource "aws_ssmcontacts_contact" "test_contact_one" {
	alias                   = "test-contact-one-for-%[2]s"
	type                    = "PERSONAL"

	depends_on              = [aws_ssmincidents_replication_set.test]
}

resource "aws_ssmcontacts_contact" "test_contact_two" {
	alias                   = "test-contact-two-for-%[2]s"
	type                    = "PERSONAL"

	depends_on              = [aws_ssmincidents_replication_set.test]
}

resource "aws_ssmcontacts_contact" "test_contact_three" {
	alias                   = "test-contact-three-for-%[2]s"
	type                    = "PERSONAL"

	depends_on              = [aws_ssmincidents_replication_set.test]
}

resource "aws_ssmcontacts_contact" "test_contact_four" {
	alias                   = "test-contact-four-for-%[2]s"
	type                    = "PERSONAL"

	depends_on              = [aws_ssmincidents_replication_set.test]
}

resource "aws_ssmcontacts_contact" "test_escalation_plan_one" {
	alias                   = "test-escalation-plan-for-%[2]s"
	type                    = "ESCALATION"

	depends_on              = [aws_ssmincidents_replication_set.test]
}
`, acctest.Region(), alias)
}
