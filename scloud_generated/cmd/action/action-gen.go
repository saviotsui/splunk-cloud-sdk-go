// Package action -- generated by scloudgen
// !! DO NOT EDIT !! 
// 
package action

import (
	"github.com/spf13/cobra"
	impl "github.com/splunk/splunk-cloud-sdk-go/scloud_generated/pkg/action"
)


// createAction -- 
var createActionCmd = &cobra.Command{
	Use:   "create-action",
	Short: "",
	RunE:  impl.CreateAction,
}

// deleteAction -- 
var deleteActionCmd = &cobra.Command{
	Use:   "delete-action",
	Short: "",
	RunE:  impl.DeleteAction,
}

// getAction -- 
var getActionCmd = &cobra.Command{
	Use:   "get-action",
	Short: "",
	RunE:  impl.GetAction,
}

// getActionStatus -- 
var getActionStatusCmd = &cobra.Command{
	Use:   "get-action-status",
	Short: "",
	RunE:  impl.GetActionStatus,
}

// getActionStatusDetails -- 
var getActionStatusDetailsCmd = &cobra.Command{
	Use:   "get-action-status-details",
	Short: "",
	RunE:  impl.GetActionStatusDetails,
}

// getPublicWebhookKeys -- 
var getPublicWebhookKeysCmd = &cobra.Command{
	Use:   "get-public-webhook-keys",
	Short: "",
	RunE:  impl.GetPublicWebhookKeys,
}

// listActions -- 
var listActionsCmd = &cobra.Command{
	Use:   "list-actions",
	Short: "",
	RunE:  impl.ListActions,
}

// triggerAction -- 
var triggerActionCmd = &cobra.Command{
	Use:   "trigger-action",
	Short: "",
	RunE:  impl.TriggerAction,
}

// updateAction -- 
var updateActionCmd = &cobra.Command{
	Use:   "update-action",
	Short: "",
	RunE:  impl.UpdateAction,
}


func init() {

    actionCmd.AddCommand(createActionCmd)
    actionCmd.AddCommand(deleteActionCmd)
    actionCmd.AddCommand(getActionCmd)
    actionCmd.AddCommand(getActionStatusCmd)
    actionCmd.AddCommand(getActionStatusDetailsCmd)
    actionCmd.AddCommand(getPublicWebhookKeysCmd)
    actionCmd.AddCommand(listActionsCmd)
    actionCmd.AddCommand(triggerActionCmd)
    actionCmd.AddCommand(updateActionCmd)
    

	// subTest1Cmd.Flags().StringP("id", "i", "", "resource identifier")
	// subTest2Cmd.Flags().StringP("id", "i", "", "resource identifier")
}