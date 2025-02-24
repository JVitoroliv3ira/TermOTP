package cmd

import (
	"fmt"
	"time"

	"github.com/JVitoroliv3ira/termotp/internal/models"
	"github.com/JVitoroliv3ira/termotp/internal/storage"
	"github.com/JVitoroliv3ira/termotp/internal/utils"
	"github.com/spf13/cobra"
)

var setupName string

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Adiciona uma nova conta TOTP",
	Long:  "Cadastra um novo serviço TOTP, armazenando sua chave de forma segura.",
	Run: func(cmd *cobra.Command, args []string) {
		utils.HandleError(utils.ValidateServiceName(setupName))

		secret, err := utils.PromptSecret()
		utils.HandleError(err)
		utils.HandleError(utils.ValidateServiceSecret(secret))

		password, err := utils.PromptPassword()
		utils.HandleError(err)
		utils.HandleError(utils.ValidatePassword(password))

		account := models.Account{
			Name:      setupName,
			Secret:    secret,
			CreatedAt: time.Now(),
		}

		utils.HandleError(storage.SaveAccount(account, password))
		fmt.Printf("\nConta '%s' cadastrada e armazenada com segurança!\n", account.Name)
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)
	setupCmd.Flags().StringVarP(&setupName, "name", "n", "", "Nome da conta (ex: gitlab)")
}
