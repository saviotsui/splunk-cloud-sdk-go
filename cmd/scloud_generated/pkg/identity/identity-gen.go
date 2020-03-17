// Package identity -- generated by scloudgen
// !! DO NOT EDIT !!
//
package identity

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud_generated/auth"
	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud_generated/flags"
	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud_generated/jsonx"
	model "github.com/splunk/splunk-cloud-sdk-go/services/identity"
)

// AddGroupMember Adds a member to a given group.
func AddGroupMember(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var group string
	err = flags.ParseFlag(cmd.Flags(), "group", &group)
	if err != nil {
		return fmt.Errorf(`error parsing "group": ` + err.Error())
	}
	var name string
	err = flags.ParseFlag(cmd.Flags(), "name", &name)
	if err != nil {
		return fmt.Errorf(`error parsing "name": ` + err.Error())
	}
	// Form the request body
	generated_request_body := model.AddGroupMemberBody{

		Name: name,
	}

	resp, err := client.IdentityService.AddGroupMember(group, generated_request_body)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// AddGroupRole Adds a role to a given group.
func AddGroupRole(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var group string
	err = flags.ParseFlag(cmd.Flags(), "group", &group)
	if err != nil {
		return fmt.Errorf(`error parsing "group": ` + err.Error())
	}
	var name string
	err = flags.ParseFlag(cmd.Flags(), "name", &name)
	if err != nil {
		return fmt.Errorf(`error parsing "name": ` + err.Error())
	}
	// Form the request body
	generated_request_body := model.AddGroupRoleBody{

		Name: name,
	}

	resp, err := client.IdentityService.AddGroupRole(group, generated_request_body)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// AddMember Adds a member to a given tenant.
func AddMember(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var name string
	err = flags.ParseFlag(cmd.Flags(), "name", &name)
	if err != nil {
		return fmt.Errorf(`error parsing "name": ` + err.Error())
	}
	// Form the request body
	generated_request_body := model.AddMemberBody{

		Name: name,
	}

	resp, err := client.IdentityService.AddMember(generated_request_body)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// AddRolePermission Adds permissions to a role in a given tenant.
func AddRolePermission(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var body string
	err = flags.ParseFlag(cmd.Flags(), "body", &body)
	if err != nil {
		return fmt.Errorf(`error parsing "body": ` + err.Error())
	}
	var role string
	err = flags.ParseFlag(cmd.Flags(), "role", &role)
	if err != nil {
		return fmt.Errorf(`error parsing "role": ` + err.Error())
	}

	resp, err := client.IdentityService.AddRolePermission(role, body)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// CreateGroup Creates a new group in a given tenant.
func CreateGroup(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var name string
	err = flags.ParseFlag(cmd.Flags(), "name", &name)
	if err != nil {
		return fmt.Errorf(`error parsing "name": ` + err.Error())
	}
	// Form the request body
	generated_request_body := model.CreateGroupBody{

		Name: name,
	}

	resp, err := client.IdentityService.CreateGroup(generated_request_body)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// CreateRole Creates a new authorization role in a given tenant.
func CreateRole(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var name string
	err = flags.ParseFlag(cmd.Flags(), "name", &name)
	if err != nil {
		return fmt.Errorf(`error parsing "name": ` + err.Error())
	}
	// Form the request body
	generated_request_body := model.CreateRoleBody{

		Name: name,
	}

	resp, err := client.IdentityService.CreateRole(generated_request_body)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// DeleteGroup Deletes a group in a given tenant.
func DeleteGroup(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var group string
	err = flags.ParseFlag(cmd.Flags(), "group", &group)
	if err != nil {
		return fmt.Errorf(`error parsing "group": ` + err.Error())
	}

	err = client.IdentityService.DeleteGroup(group)
	if err != nil {
		return err
	}

	return nil
}

// DeleteRole Deletes a defined role for a given tenant.
func DeleteRole(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var role string
	err = flags.ParseFlag(cmd.Flags(), "role", &role)
	if err != nil {
		return fmt.Errorf(`error parsing "role": ` + err.Error())
	}

	err = client.IdentityService.DeleteRole(role)
	if err != nil {
		return err
	}

	return nil
}

// GetGroup Returns information about a given group within a tenant.
func GetGroup(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var group string
	err = flags.ParseFlag(cmd.Flags(), "group", &group)
	if err != nil {
		return fmt.Errorf(`error parsing "group": ` + err.Error())
	}

	resp, err := client.IdentityService.GetGroup(group)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// GetGroupMember Returns information about a given member within a given group.
func GetGroupMember(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var group string
	err = flags.ParseFlag(cmd.Flags(), "group", &group)
	if err != nil {
		return fmt.Errorf(`error parsing "group": ` + err.Error())
	}
	var member string
	err = flags.ParseFlag(cmd.Flags(), "member", &member)
	if err != nil {
		return fmt.Errorf(`error parsing "member": ` + err.Error())
	}

	resp, err := client.IdentityService.GetGroupMember(group, member)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// GetGroupRole Returns information about a given role within a given group.
func GetGroupRole(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var group string
	err = flags.ParseFlag(cmd.Flags(), "group", &group)
	if err != nil {
		return fmt.Errorf(`error parsing "group": ` + err.Error())
	}
	var role string
	err = flags.ParseFlag(cmd.Flags(), "role", &role)
	if err != nil {
		return fmt.Errorf(`error parsing "role": ` + err.Error())
	}

	resp, err := client.IdentityService.GetGroupRole(group, role)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// GetMember Returns a member of a given tenant.
func GetMember(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var member string
	err = flags.ParseFlag(cmd.Flags(), "member", &member)
	if err != nil {
		return fmt.Errorf(`error parsing "member": ` + err.Error())
	}

	resp, err := client.IdentityService.GetMember(member)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// GetPrincipal Returns the details of a principal, including its tenant membership.
func GetPrincipal(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClientSystemTenant()

	if err != nil {
		return err
	}
	// Parse all flags

	var principal string
	err = flags.ParseFlag(cmd.Flags(), "principal", &principal)
	if err != nil {
		return fmt.Errorf(`error parsing "principal": ` + err.Error())
	}

	resp, err := client.IdentityService.GetPrincipal(principal)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// GetRole Returns a role for a given tenant.
func GetRole(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var role string
	err = flags.ParseFlag(cmd.Flags(), "role", &role)
	if err != nil {
		return fmt.Errorf(`error parsing "role": ` + err.Error())
	}

	resp, err := client.IdentityService.GetRole(role)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// GetRolePermission Gets a permission for the specified role.
func GetRolePermission(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var permission string
	err = flags.ParseFlag(cmd.Flags(), "permission", &permission)
	if err != nil {
		return fmt.Errorf(`error parsing "permission": ` + err.Error())
	}
	var role string
	err = flags.ParseFlag(cmd.Flags(), "role", &role)
	if err != nil {
		return fmt.Errorf(`error parsing "role": ` + err.Error())
	}

	resp, err := client.IdentityService.GetRolePermission(role, permission)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// ListGroupMembers Returns a list of the members within a given group.
func ListGroupMembers(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var group string
	err = flags.ParseFlag(cmd.Flags(), "group", &group)
	if err != nil {
		return fmt.Errorf(`error parsing "group": ` + err.Error())
	}

	resp, err := client.IdentityService.ListGroupMembers(group)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// ListGroupRoles Returns a list of the roles that are attached to a group within a given tenant.
func ListGroupRoles(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var group string
	err = flags.ParseFlag(cmd.Flags(), "group", &group)
	if err != nil {
		return fmt.Errorf(`error parsing "group": ` + err.Error())
	}

	resp, err := client.IdentityService.ListGroupRoles(group)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// ListGroups List the groups that exist in a given tenant.
func ListGroups(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var access model.ListGroupsaccess
	err = flags.ParseFlag(cmd.Flags(), "access", &access)
	if err != nil {
		return fmt.Errorf(`error parsing "access": ` + err.Error())
	}
	// Form query params
	generated_query := model.ListGroupsQueryParams{}
	generated_query.Access = access

	resp, err := client.IdentityService.ListGroups(&generated_query)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// ListMemberGroups Returns a list of groups that a member belongs to within a tenant.
func ListMemberGroups(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var member string
	err = flags.ParseFlag(cmd.Flags(), "member", &member)
	if err != nil {
		return fmt.Errorf(`error parsing "member": ` + err.Error())
	}

	resp, err := client.IdentityService.ListMemberGroups(member)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// ListMemberPermissions Returns a set of permissions granted to the member within the tenant.

func ListMemberPermissions(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var member string
	err = flags.ParseFlag(cmd.Flags(), "member", &member)
	if err != nil {
		return fmt.Errorf(`error parsing "member": ` + err.Error())
	}

	resp, err := client.IdentityService.ListMemberPermissions(member)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// ListMemberRoles Returns a set of roles that a given member holds within the tenant.

func ListMemberRoles(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var member string
	err = flags.ParseFlag(cmd.Flags(), "member", &member)
	if err != nil {
		return fmt.Errorf(`error parsing "member": ` + err.Error())
	}

	resp, err := client.IdentityService.ListMemberRoles(member)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// ListMembers Returns a list of members in a given tenant.
func ListMembers(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}

	resp, err := client.IdentityService.ListMembers()
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// ListPrincipals Returns the list of principals that the Identity service knows about.
func ListPrincipals(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClientSystemTenant()

	if err != nil {
		return err
	}

	resp, err := client.IdentityService.ListPrincipals()
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// ListRoleGroups Gets a list of groups for a role in a given tenant.
func ListRoleGroups(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var role string
	err = flags.ParseFlag(cmd.Flags(), "role", &role)
	if err != nil {
		return fmt.Errorf(`error parsing "role": ` + err.Error())
	}

	resp, err := client.IdentityService.ListRoleGroups(role)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// ListRolePermissions Gets the permissions for a role in a given tenant.
func ListRolePermissions(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var role string
	err = flags.ParseFlag(cmd.Flags(), "role", &role)
	if err != nil {
		return fmt.Errorf(`error parsing "role": ` + err.Error())
	}

	resp, err := client.IdentityService.ListRolePermissions(role)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// ListRoles Returns all roles for a given tenant.
func ListRoles(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}

	resp, err := client.IdentityService.ListRoles()
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// RemoveGroupMember Removes the member from a given group.
func RemoveGroupMember(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var group string
	err = flags.ParseFlag(cmd.Flags(), "group", &group)
	if err != nil {
		return fmt.Errorf(`error parsing "group": ` + err.Error())
	}
	var member string
	err = flags.ParseFlag(cmd.Flags(), "member", &member)
	if err != nil {
		return fmt.Errorf(`error parsing "member": ` + err.Error())
	}

	err = client.IdentityService.RemoveGroupMember(group, member)
	if err != nil {
		return err
	}

	return nil
}

// RemoveGroupRole Removes a role from a given group.
func RemoveGroupRole(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var group string
	err = flags.ParseFlag(cmd.Flags(), "group", &group)
	if err != nil {
		return fmt.Errorf(`error parsing "group": ` + err.Error())
	}
	var role string
	err = flags.ParseFlag(cmd.Flags(), "role", &role)
	if err != nil {
		return fmt.Errorf(`error parsing "role": ` + err.Error())
	}

	err = client.IdentityService.RemoveGroupRole(group, role)
	if err != nil {
		return err
	}

	return nil
}

// RemoveMember Removes a member from a given tenant
func RemoveMember(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var member string
	err = flags.ParseFlag(cmd.Flags(), "member", &member)
	if err != nil {
		return fmt.Errorf(`error parsing "member": ` + err.Error())
	}

	err = client.IdentityService.RemoveMember(member)
	if err != nil {
		return err
	}

	return nil
}

// RemoveRolePermission Removes a permission from the role.
func RemoveRolePermission(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var permission string
	err = flags.ParseFlag(cmd.Flags(), "permission", &permission)
	if err != nil {
		return fmt.Errorf(`error parsing "permission": ` + err.Error())
	}
	var role string
	err = flags.ParseFlag(cmd.Flags(), "role", &role)
	if err != nil {
		return fmt.Errorf(`error parsing "role": ` + err.Error())
	}

	err = client.IdentityService.RemoveRolePermission(role, permission)
	if err != nil {
		return err
	}

	return nil
}

// RevokePrincipalAuthTokens Revoke all existing tokens issued to a principal
func RevokePrincipalAuthTokens(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClientSystemTenant()

	if err != nil {
		return err
	}
	// Parse all flags

	var principal string
	err = flags.ParseFlag(cmd.Flags(), "principal", &principal)
	if err != nil {
		return fmt.Errorf(`error parsing "principal": ` + err.Error())
	}

	err = client.IdentityService.RevokePrincipalAuthTokens(principal)
	if err != nil {
		return err
	}

	return nil
}

// ValidateToken Validates the access token obtained from the authorization header and returns the principal name and tenant memberships.

func ValidateToken(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var include model.ValidateTokenincludeEnum
	err = flags.ParseFlag(cmd.Flags(), "include", &include)
	if err != nil {
		return fmt.Errorf(`error parsing "include": ` + err.Error())
	}
	// Form query params
	generated_query := model.ValidateTokenQueryParams{}

	generated_query.Include = []model.ValidateTokenincludeEnum{include}

	resp, err := client.IdentityService.ValidateToken(&generated_query)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}
