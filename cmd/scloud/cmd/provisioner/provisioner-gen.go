// Package provisioner -- generated by scloudgen
// !! DO NOT EDIT !!
//
package provisioner

import (
	"github.com/spf13/cobra"
	impl "github.com/splunk/splunk-cloud-sdk-go/cmd/scloud/pkg/provisioner"
)

// createInvite -- Creates an invitation for a person to join the tenant using their email address.
var createInviteCmd = &cobra.Command{
	Use:   "create-invite",
	Short: "Creates an invitation for a person to join the tenant using their email address.",
	RunE:  impl.CreateInvite,
}

// createProvisionJob -- Creates a new job that provisions a new tenant and subscribes apps to the tenant.
var createProvisionJobCmd = &cobra.Command{
	Use:   "create-provision-job",
	Short: "Creates a new job that provisions a new tenant and subscribes apps to the tenant.",
	RunE:  impl.CreateProvisionJob,
}

// deleteInvite -- Removes an invitation in the given tenant.
var deleteInviteCmd = &cobra.Command{
	Use:   "delete-invite",
	Short: "Removes an invitation in the given tenant.",
	RunE:  impl.DeleteInvite,
}

// getInvite -- Returns an invitation in the given tenant.
var getInviteCmd = &cobra.Command{
	Use:   "get-invite",
	Short: "Returns an invitation in the given tenant.",
	RunE:  impl.GetInvite,
}

// getProvisionJob -- Returns details of a specific provision job.
var getProvisionJobCmd = &cobra.Command{
	Use:   "get-provision-job",
	Short: "Returns details of a specific provision job.",
	RunE:  impl.GetProvisionJob,
}

// getTenant -- Returns a specific tenant.
var getTenantCmd = &cobra.Command{
	Use:   "get-tenant",
	Short: "Returns a specific tenant.",
	RunE:  impl.GetTenant,
}

// listInvites -- Returns a list of invitations in a given tenant.
var listInvitesCmd = &cobra.Command{
	Use:   "list-invites",
	Short: "Returns a list of invitations in a given tenant.",
	RunE:  impl.ListInvites,
}

// listProvisionJobs -- Returns a list of all provision jobs created by the user.
var listProvisionJobsCmd = &cobra.Command{
	Use:   "list-provision-jobs",
	Short: "Returns a list of all provision jobs created by the user.",
	RunE:  impl.ListProvisionJobs,
}

// listTenants -- Returns all tenants that the user can read.
var listTenantsCmd = &cobra.Command{
	Use:   "list-tenants",
	Short: "Returns all tenants that the user can read.",
	RunE:  impl.ListTenants,
}

// updateInvite -- Modifies an invitation in the given tenant.
var updateInviteCmd = &cobra.Command{
	Use:   "update-invite",
	Short: "Modifies an invitation in the given tenant.",
	RunE:  impl.UpdateInvite,
}

func init() {
	provisionerCmd.AddCommand(createInviteCmd)

	var createInviteEmail string
	createInviteCmd.Flags().StringVar(&createInviteEmail, "email", "", "This is a required parameter. ")
	createInviteCmd.MarkFlagRequired("email")

	var createInviteComment string
	createInviteCmd.Flags().StringVar(&createInviteComment, "comment", "", "")

	var createInviteGroups []string
	createInviteCmd.Flags().StringSliceVar(&createInviteGroups, "groups", nil, "")

	provisionerCmd.AddCommand(createProvisionJobCmd)

	var createProvisionJobApps []string
	createProvisionJobCmd.Flags().StringSliceVar(&createProvisionJobApps, "apps", nil, "")

	var createProvisionJobTenant string
	createProvisionJobCmd.Flags().StringVar(&createProvisionJobTenant, "tenant", "", "")

	provisionerCmd.AddCommand(deleteInviteCmd)

	var deleteInviteInviteId string
	deleteInviteCmd.Flags().StringVar(&deleteInviteInviteId, "invite-id", "", "This is a required parameter. ")
	deleteInviteCmd.MarkFlagRequired("invite-id")

	provisionerCmd.AddCommand(getInviteCmd)

	var getInviteInviteId string
	getInviteCmd.Flags().StringVar(&getInviteInviteId, "invite-id", "", "This is a required parameter. ")
	getInviteCmd.MarkFlagRequired("invite-id")

	provisionerCmd.AddCommand(getProvisionJobCmd)

	var getProvisionJobJobId string
	getProvisionJobCmd.Flags().StringVar(&getProvisionJobJobId, "job-id", "", "This is a required parameter. ")
	getProvisionJobCmd.MarkFlagRequired("job-id")

	provisionerCmd.AddCommand(getTenantCmd)

	var getTenantTenantName string
	getTenantCmd.Flags().StringVar(&getTenantTenantName, "tenant-name", "", "This is a required parameter. ")
	getTenantCmd.MarkFlagRequired("tenant-name")

	provisionerCmd.AddCommand(listInvitesCmd)

	provisionerCmd.AddCommand(listProvisionJobsCmd)

	provisionerCmd.AddCommand(listTenantsCmd)

	provisionerCmd.AddCommand(updateInviteCmd)

	var updateInviteAction string
	updateInviteCmd.Flags().StringVar(&updateInviteAction, "action", "", "This is a required parameter.  can accept values accept, reject, resend")
	updateInviteCmd.MarkFlagRequired("action")

	var updateInviteInviteId string
	updateInviteCmd.Flags().StringVar(&updateInviteInviteId, "invite-id", "", "This is a required parameter. ")
	updateInviteCmd.MarkFlagRequired("invite-id")

}