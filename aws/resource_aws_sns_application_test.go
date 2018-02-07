package aws

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/sns"
)

/**
 Before running this test, a few ENV variables must be set:
 GCM_API_KEY - Google Cloud Messaging API Key
 APNS_SANDBOX_CREDENTIAL - Apple Push Notification Sandbox Private Key
 APNS_SANDBOX_PRINCIPAL - Apple Push Notification Sandbox Certificate
**/

func TestAccAWSSNSApplication_gcm_create_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			if os.Getenv("GCM_API_KEY") == "" {
				t.Fatal("GCM_API_KEY must be set.")
			}
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAWSSNSApplicationDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccAWSSNSApplicationGCMConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"aws_sns_application.gcm_test", "name", "aws_sns_application_test"),
					resource.TestCheckResourceAttr(
						"aws_sns_application.gcm_test", "platform", "GCM"),
					resource.TestCheckResourceAttr(
						"aws_sns_application.gcm_test", "success_feedback_sample_rate", "100"),
					resource.TestCheckResourceAttr(
						"aws_sns_application.gcm_test", "event_endpoint_created_topic_arn", "arn:aws:sns:us-east-1:638386993804:endpoint-created-topic"),
					resource.TestCheckResourceAttr(
						"aws_sns_application.gcm_test", "event_endpoint_updated_topic_arn", "arn:aws:sns:us-east-1:638386993804:endpoint-updated-topic"),
					resource.TestCheckResourceAttr(
						"aws_sns_application.gcm_test", "event_delivery_failure_topic_arn", "arn:aws:sns:us-east-1:638386993804:endpoint-failure-topic"),
					resource.TestCheckResourceAttr(
						"aws_sns_application.gcm_test", "event_endpoint_deleted_topic_arn", "arn:aws:sns:us-east-1:638386993804:endpoint-deleted-topic"),
				),
			},
			resource.TestStep{
				Config: testAccAWSSNSApplicationGCMConfigUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"aws_sns_application.gcm_test", "name", "aws_sns_application_test"),
					resource.TestCheckResourceAttr(
						"aws_sns_application.gcm_test", "platform", "GCM"),
					resource.TestCheckResourceAttr(
						"aws_sns_application.gcm_test", "success_feedback_sample_rate", "99"),
					resource.TestCheckResourceAttr(
						"aws_sns_application.gcm_test", "event_endpoint_created_topic_arn", "arn:aws:sns:us-east-1:638386993804:endpoint-created-topic-update"),
					resource.TestCheckResourceAttr(
						"aws_sns_application.gcm_test", "event_endpoint_updated_topic_arn", "arn:aws:sns:us-east-1:638386993804:endpoint-updated-topic-update"),
					resource.TestCheckResourceAttr(
						"aws_sns_application.gcm_test", "event_delivery_failure_topic_arn", "arn:aws:sns:us-east-1:638386993804:endpoint-failure-topic-update"),
					resource.TestCheckResourceAttr(
						"aws_sns_application.gcm_test", "event_endpoint_deleted_topic_arn", "arn:aws:sns:us-east-1:638386993804:endpoint-deleted-topic-update"),
				),
			},
		},
	})
}

func TestAccAWSSNSApplication_apns_sandbox_create_update(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			if os.Getenv("APNS_SANDBOX_CREDENTIAL") == "" {
				t.Fatal("APNS_SANDBOX_CREDENTIAL must be set.")
			}
			if os.Getenv("APNS_SANDBOX_PRINCIPAL") == "" {
				t.Fatal("APNS_SANDBOX_PRINCIPAL must be set.")
			}
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAWSSNSApplicationDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccAWSSNSApplicationAPNSSandBoxConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"aws_sns_application.apns_test", "name", "aws_sns_application_test"),
					resource.TestCheckResourceAttr(
						"aws_sns_application.apns_test", "platform", "APNS_SANDBOX"),
					resource.TestCheckResourceAttr(
						"aws_sns_application.apns_test", "success_feedback_sample_rate", "100"),
					resource.TestCheckResourceAttr(
						"aws_sns_application.apns_test", "event_endpoint_created_topic_arn", "arn:aws:sns:us-east-1:638386993804:endpoint-created-topic"),
					resource.TestCheckResourceAttr(
						"aws_sns_application.apns_test", "event_endpoint_updated_topic_arn", "arn:aws:sns:us-east-1:638386993804:endpoint-updated-topic"),
					resource.TestCheckResourceAttr(
						"aws_sns_application.apns_test", "event_delivery_failure_topic_arn", "arn:aws:sns:us-east-1:638386993804:endpoint-failure-topic"),
					resource.TestCheckResourceAttr(
						"aws_sns_application.apns_test", "event_endpoint_deleted_topic_arn", "arn:aws:sns:us-east-1:638386993804:endpoint-deleted-topic"),
				),
			},
			resource.TestStep{
				Config: testAccAWSSNSApplicationAPNSSandBoxConfigUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"aws_sns_application.apns_test", "name", "aws_sns_application_test"),
					resource.TestCheckResourceAttr(
						"aws_sns_application.apns_test", "platform", "APNS_SANDBOX"),
					resource.TestCheckResourceAttr(
						"aws_sns_application.apns_test", "success_feedback_sample_rate", "99"),
					resource.TestCheckResourceAttr(
						"aws_sns_application.apns_test", "event_endpoint_created_topic_arn", "arn:aws:sns:us-east-1:638386993804:endpoint-created-topic-update"),
					resource.TestCheckResourceAttr(
						"aws_sns_application.apns_test", "event_endpoint_updated_topic_arn", "arn:aws:sns:us-east-1:638386993804:endpoint-updated-topic-update"),
					resource.TestCheckResourceAttr(
						"aws_sns_application.apns_test", "event_delivery_failure_topic_arn", "arn:aws:sns:us-east-1:638386993804:endpoint-failure-topic-update"),
					resource.TestCheckResourceAttr(
						"aws_sns_application.apns_test", "event_endpoint_deleted_topic_arn", "arn:aws:sns:us-east-1:638386993804:endpoint-deleted-topic-update"),
				),
			},
		},
	})
}

func testAccCheckAWSSNSApplicationDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*AWSClient).snsconn

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "aws_sns_application" {
			continue
		}
		_, err := conn.DeletePlatformApplication(&sns.DeletePlatformApplicationInput{
			PlatformApplicationArn: aws.String(rs.Primary.ID),
		})
		if err != nil {
			if awsErr, ok := err.(awserr.Error); ok && awsErr.Code() == "NoSNSApplication" {
				return nil
			}
			return err
		}
	}
	return nil
}

var testAccAWSSNSApplicationGCMConfig = `
resource "aws_sns_application" "gcm_test" {
	event_delivery_failure_topic_arn = "arn:aws:sns:us-east-1:638386993804:endpoint-failure-topic"
	event_endpoint_created_topic_arn = "arn:aws:sns:us-east-1:638386993804:endpoint-created-topic"
	event_endpoint_deleted_topic_arn = "arn:aws:sns:us-east-1:638386993804:endpoint-deleted-topic"
	event_endpoint_updated_topic_arn = "arn:aws:sns:us-east-1:638386993804:endpoint-updated-topic"
	name = "aws_sns_application_test"
	platform = "GCM"
	platform_credential = "` + os.Getenv("GCM_API_KEY") + `"
	success_feedback_sample_rate = 100
}
`

var testAccAWSSNSApplicationGCMConfigUpdate = `
	event_delivery_failure_topic_arn = "arn:aws:sns:us-east-1:638386993804:endpoint-failure-topic-update"
	event_endpoint_created_topic_arn = "arn:aws:sns:us-east-1:638386993804:endpoint-created-topic-update"
	event_endpoint_deleted_topic_arn = "arn:aws:sns:us-east-1:638386993804:endpoint-deleted-topic-update"
	event_endpoint_updated_topic_arn = "arn:aws:sns:us-east-1:638386993804:endpoint-updated-topic-update"
	name = "aws_sns_application_test"
	platform = "GCM"
	platform_credential = "` + os.Getenv("GCM_API_KEY") + `"
	resource "aws_sns_application" "gcm_test" {
	success_feedback_sample_rate = 99
}
`

var testAccAWSSNSApplicationAPNSSandBoxConfig = `
resource "aws_sns_application" "apns_test" {
	event_delivery_failure_topic_arn = "arn:aws:sns:us-east-1:638386993804:endpoint-failure-topic"
	event_endpoint_created_topic_arn = "arn:aws:sns:us-east-1:638386993804:endpoint-created-topic"
	event_endpoint_deleted_topic_arn = "arn:aws:sns:us-east-1:638386993804:endpoint-deleted-topic"
	event_endpoint_updated_topic_arn = "arn:aws:sns:us-east-1:638386993804:endpoint-updated-topic"
	name = "aws_sns_application_test"
	platform = "APNS_SANDBOX"
	platform_credential = "` + os.Getenv("APNS_SANDBOX_CREDENTIAL") + `"
	platform_principal = "` + os.Getenv("APNS_SANDBOX_PRINCIPAL") + `"
	success_feedback_sample_rate = 100
}
`

var testAccAWSSNSApplicationAPNSSandBoxConfigUpdate = `
resource "aws_sns_application" "apns_test" {
	event_delivery_failure_topic_arn = "arn:aws:sns:us-east-1:638386993804:endpoint-failure-topic-update"
	event_endpoint_created_topic_arn = "arn:aws:sns:us-east-1:638386993804:endpoint-created-topic-update"
	event_endpoint_deleted_topic_arn = "arn:aws:sns:us-east-1:638386993804:endpoint-deleted-topic-update"
	event_endpoint_updated_topic_arn = "arn:aws:sns:us-east-1:638386993804:endpoint-updated-topic-update"
	name = "aws_sns_application_test"
	platform = "APNS_SANDBOX"
	platform_credential = "` + os.Getenv("APNS_SANDBOX_CREDENTIAL") + `"
	platform_principal = "` + os.Getenv("APNS_SANDBOX_PRINCIPAL") + `"
	success_feedback_sample_rate = 99
}
`