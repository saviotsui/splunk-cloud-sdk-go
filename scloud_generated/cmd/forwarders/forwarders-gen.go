// Package forwarders -- generated by scloudgen
// !! DO NOT EDIT !!
//
package forwarders

import (
	"github.com/spf13/cobra"
	impl "github.com/splunk/splunk-cloud-sdk-go/scloud_generated/pkg/forwarders"
)

// addCertificate -- Each tenant can have up to five certificates.
var addCertificateCmd = &cobra.Command{
	Use:   "add-certificate",
	Short: "Each tenant can have up to five certificates.",
	RunE:  impl.AddCertificate,
}

// deleteCertificate --
var deleteCertificateCmd = &cobra.Command{
	Use:   "delete-certificate",
	Short: "",
	RunE:  impl.DeleteCertificate,
}

// deleteCertificates --
var deleteCertificatesCmd = &cobra.Command{
	Use:   "delete-certificates",
	Short: "",
	RunE:  impl.DeleteCertificates,
}

// listCertificates --
var listCertificatesCmd = &cobra.Command{
	Use:   "list-certificates",
	Short: "",
	RunE:  impl.ListCertificates,
}

func init() {
	forwardersCmd.AddCommand(addCertificateCmd)

	var addCertificatePem string
	addCertificateCmd.Flags().StringVar(&addCertificatePem, "pem", "", "")

	forwardersCmd.AddCommand(deleteCertificateCmd)
	var deleteCertificateSlot string
	deleteCertificateCmd.Flags().StringVar(&deleteCertificateSlot, "slot", "", "")
	deleteCertificateCmd.MarkFlagRequired("slot")

	forwardersCmd.AddCommand(deleteCertificatesCmd)

	forwardersCmd.AddCommand(listCertificatesCmd)

}
