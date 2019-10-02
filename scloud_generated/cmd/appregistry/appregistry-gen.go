// Package appregistry -- generated by scloudgen
// !! DO NOT EDIT !!
//
package appregistry

import (
	"github.com/spf13/cobra"
	impl "github.com/splunk/splunk-cloud-sdk-go/scloud_generated/pkg/appregistry"
)

// createApp -- Creates an app.
var createAppCmd = &cobra.Command{
	Use:   "create-app",
	Short: "Creates an app.",
	RunE:  impl.CreateApp,
}

// createSubscription -- Creates a subscription.
var createSubscriptionCmd = &cobra.Command{
	Use:   "create-subscription",
	Short: "Creates a subscription.",
	RunE:  impl.CreateSubscription,
}

// deleteApp -- Removes an app.
var deleteAppCmd = &cobra.Command{
	Use:   "delete-app",
	Short: "Removes an app.",
	RunE:  impl.DeleteApp,
}

// deleteSubscription -- Removes a subscription.
var deleteSubscriptionCmd = &cobra.Command{
	Use:   "delete-subscription",
	Short: "Removes a subscription.",
	RunE:  impl.DeleteSubscription,
}

// getApp -- Returns the metadata of an app.
var getAppCmd = &cobra.Command{
	Use:   "get-app",
	Short: "Returns the metadata of an app.",
	RunE:  impl.GetApp,
}

// getKeys -- Returns a list of the public keys used for verifying signed webhook requests.
var getKeysCmd = &cobra.Command{
	Use:   "get-keys",
	Short: "Returns a list of the public keys used for verifying signed webhook requests.",
	RunE:  impl.GetKeys,
}

// getSubscription -- Returns or validates a subscription.
var getSubscriptionCmd = &cobra.Command{
	Use:   "get-subscription",
	Short: "Returns or validates a subscription.",
	RunE:  impl.GetSubscription,
}

// listAppSubscriptions -- Returns the collection of subscriptions to an app.
var listAppSubscriptionsCmd = &cobra.Command{
	Use:   "list-app-subscriptions",
	Short: "Returns the collection of subscriptions to an app.",
	RunE:  impl.ListAppSubscriptions,
}

// listApps -- Returns a list of apps.
var listAppsCmd = &cobra.Command{
	Use:   "list-apps",
	Short: "Returns a list of apps.",
	RunE:  impl.ListApps,
}

// listSubscriptions -- Returns the tenant subscriptions.
var listSubscriptionsCmd = &cobra.Command{
	Use:   "list-subscriptions",
	Short: "Returns the tenant subscriptions.",
	RunE:  impl.ListSubscriptions,
}

// rotateSecret -- Rotates the client secret for an app.
var rotateSecretCmd = &cobra.Command{
	Use:   "rotate-secret",
	Short: "Rotates the client secret for an app.",
	RunE:  impl.RotateSecret,
}

// updateApp -- Updates an app.
var updateAppCmd = &cobra.Command{
	Use:   "update-app",
	Short: "Updates an app.",
	RunE:  impl.UpdateApp,
}

func init() {
	appregistryCmd.AddCommand(createAppCmd)
	var createAppKind string
	createAppCmd.Flags().StringVar(&createAppKind, "kind", "", " can accept values web, native, service")
	createAppCmd.MarkFlagRequired("kind")
	var createAppName string
	createAppCmd.Flags().StringVar(&createAppName, "name", "", "App name that is unique within Splunk Cloud Platform.")
	createAppCmd.MarkFlagRequired("name")
	var createAppTitle string
	createAppCmd.Flags().StringVar(&createAppTitle, "title", "", "Human-readable title for the app.")
	createAppCmd.MarkFlagRequired("title")

	var createAppAppPrincipalPermissions string
	createAppCmd.Flags().StringVar(&createAppAppPrincipalPermissions, "app-principal-permissions", "", "Array of permission templates that are used to grant permission to the app principal when a tenant subscribes.")
	var createAppDescription string
	createAppCmd.Flags().StringVar(&createAppDescription, "description", "", "Short paragraph describing the app.")
	var createAppLoginUrl string
	createAppCmd.Flags().StringVar(&createAppLoginUrl, "login-url", "", "The URL used to log in to the app.")
	var createAppLogoUrl string
	createAppCmd.Flags().StringVar(&createAppLogoUrl, "logo-url", "", "The URL used to display the app's logo.")
	var createAppRedirectUrls string
	createAppCmd.Flags().StringVar(&createAppRedirectUrls, "redirect-urls", "", "Array of URLs that can be used for redirect after logging into the app.")
	var createAppSetupUrl string
	createAppCmd.Flags().StringVar(&createAppSetupUrl, "setup-url", "", "URL to redirect to after a subscription is created.")
	var createAppUserPermissionsFilter string
	createAppCmd.Flags().StringVar(&createAppUserPermissionsFilter, "user-permissions-filter", "", "Array of permission filter templates that are used to intersect with a user's permissions when using the app.")
	var createAppWebhookUrl string
	createAppCmd.Flags().StringVar(&createAppWebhookUrl, "webhook-url", "", "URL that webhook events are sent to.")

	appregistryCmd.AddCommand(createSubscriptionCmd)
	var createSubscriptionAppName string
	createSubscriptionCmd.Flags().StringVar(&createSubscriptionAppName, "app-name", "", "")
	createSubscriptionCmd.MarkFlagRequired("app-name")

	appregistryCmd.AddCommand(deleteAppCmd)
	var deleteAppAppName string
	deleteAppCmd.Flags().StringVar(&deleteAppAppName, "app-name", "", "App name.")
	deleteAppCmd.MarkFlagRequired("app-name")

	appregistryCmd.AddCommand(deleteSubscriptionCmd)
	var deleteSubscriptionAppName string
	deleteSubscriptionCmd.Flags().StringVar(&deleteSubscriptionAppName, "app-name", "", "App name.")
	deleteSubscriptionCmd.MarkFlagRequired("app-name")

	appregistryCmd.AddCommand(getAppCmd)
	var getAppAppName string
	getAppCmd.Flags().StringVar(&getAppAppName, "app-name", "", "App name.")
	getAppCmd.MarkFlagRequired("app-name")

	appregistryCmd.AddCommand(getKeysCmd)

	appregistryCmd.AddCommand(getSubscriptionCmd)
	var getSubscriptionAppName string
	getSubscriptionCmd.Flags().StringVar(&getSubscriptionAppName, "app-name", "", "App name.")
	getSubscriptionCmd.MarkFlagRequired("app-name")

	appregistryCmd.AddCommand(listAppSubscriptionsCmd)
	var listAppSubscriptionsAppName string
	listAppSubscriptionsCmd.Flags().StringVar(&listAppSubscriptionsAppName, "app-name", "", "App name.")
	listAppSubscriptionsCmd.MarkFlagRequired("app-name")

	appregistryCmd.AddCommand(listAppsCmd)

	appregistryCmd.AddCommand(listSubscriptionsCmd)

	var listSubscriptionsKind string
	listSubscriptionsCmd.Flags().StringVar(&listSubscriptionsKind, "kind", "", "The kind of application.")

	appregistryCmd.AddCommand(rotateSecretCmd)
	var rotateSecretAppName string
	rotateSecretCmd.Flags().StringVar(&rotateSecretAppName, "app-name", "", "App name.")
	rotateSecretCmd.MarkFlagRequired("app-name")

	appregistryCmd.AddCommand(updateAppCmd)
	var updateAppAppName string
	updateAppCmd.Flags().StringVar(&updateAppAppName, "app-name", "", "App name.")
	updateAppCmd.MarkFlagRequired("app-name")
	var updateAppTitle string
	updateAppCmd.Flags().StringVar(&updateAppTitle, "title", "", "Human-readable title for the app.")
	updateAppCmd.MarkFlagRequired("title")

	var updateAppAppPrincipalPermissions string
	updateAppCmd.Flags().StringVar(&updateAppAppPrincipalPermissions, "app-principal-permissions", "", "Array of permission templates that are used to grant permission to the app principal when a tenant subscribes.")
	var updateAppDescription string
	updateAppCmd.Flags().StringVar(&updateAppDescription, "description", "", "Short paragraph describing the app.")
	var updateAppLoginUrl string
	updateAppCmd.Flags().StringVar(&updateAppLoginUrl, "login-url", "", "The URL used to log in to the app.")
	var updateAppLogoUrl string
	updateAppCmd.Flags().StringVar(&updateAppLogoUrl, "logo-url", "", "The URL used to display the app's logo.")
	var updateAppRedirectUrls string
	updateAppCmd.Flags().StringVar(&updateAppRedirectUrls, "redirect-urls", "", "Array of URLs that can be used for redirect after logging into the app.")
	var updateAppSetupUrl string
	updateAppCmd.Flags().StringVar(&updateAppSetupUrl, "setup-url", "", "URL to redirect to after a subscription is created.")
	var updateAppUserPermissionsFilter string
	updateAppCmd.Flags().StringVar(&updateAppUserPermissionsFilter, "user-permissions-filter", "", "Array of permission filter templates that are used to intersect with a user's permissions when using the app.")
	var updateAppWebhookUrl string
	updateAppCmd.Flags().StringVar(&updateAppWebhookUrl, "webhook-url", "", "URL that webhook events are sent to.")

}
