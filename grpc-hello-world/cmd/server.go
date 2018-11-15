package cmd

import (
	"log"

	"test.go.dev/grpc-hello-world/server"

	"github.com/spf13/cobra"
)

/* var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run the gRPC hello-world server",
	Run: func(cmd *cobra.Command, args []string) {
		defer func() {
			if err := recover(); err != nil {
				log.Println("Recover error : %v", err)
			}
		}()
		server.Serve()
	},
}
*/

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run the gRPC hello-world server",
	Run: func(cmd *cobra.Command, args []string) {
		defer func() {
			if err := recover(); err != nil {
				log.Println("Recover error : %v", err)
			}
		}()

		server.Run()
	},
}

func init() {
	serverCmd.Flags().StringVarP(&server.ServerPort, "port", "p", "8081", "server port")
	//serverCmd.Flags().StringVarP(&server.CertCr, "caroot", "", "/smb/tech/go-v/src/test.go.dev/grpc-hello-world/certs/ca.cer", "caroot path")
	//serverCmd.Flags().StringVarP(&server.CertPemPath, "cert-pem", "", "./certs/server.cer", "cert pem path")
	//serverCmd.Flags().StringVarP(&server.CertKeyPath, "cert-key", "", "./certs/server.key", "cert key path")
	serverCmd.Flags().StringVarP(&server.CertPemPath, "cert-pem", "", "/smb/tech/go-v/src/test.go.dev/grpc-hello-world/certs/cert.pem", "cert pem path")
	serverCmd.Flags().StringVarP(&server.CertKeyPath, "cert-key", "", "/smb/tech/go-v/src/test.go.dev/grpc-hello-world/certs/key.pem", "cert key path")
	serverCmd.Flags().StringVarP(&server.CertServerName, "cert-name", "", "localhost", "server's hostname")
	serverCmd.Flags().StringVarP(&server.SwaggerDir, "swagger-dir", "", "proto", "path to the directory which contains swagger definitions")
	rootCmd.AddCommand(serverCmd)
}
