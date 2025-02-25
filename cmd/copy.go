package cmd

import (
	"fmt"

	"github.com/JVitoroliv3ira/termotp/internal/storage"
	"github.com/JVitoroliv3ira/termotp/internal/totp"
	"github.com/JVitoroliv3ira/termotp/internal/utils"
	"github.com/spf13/cobra"
)

var copyName string

var copyCmd = &cobra.Command{
	Use:   "copy",
	Short: "Copia um código TOTP sem exibir",
	Long:  "Gera um código TOTP e copia automaticamente para a área de transferência.",
	Run: func(cmd *cobra.Command, args []string) {
		utils.HandleError(utils.ValidateServiceName(copyName))

		password, err := utils.PromptPassword()
		utils.HandleError(err)
		utils.HandleError(utils.ValidatePassword(password))

		account, err := storage.GetAccount(copyName, password)
		utils.HandleError(err)

		code, secondsRemaining, err := totp.GenerateTOTP(account.Secret)
		utils.HandleError(err)

		utils.HandleError(utils.CopyToClipboard(code))
		fmt.Printf("Código TOTP para a conta [%s]: %s\n", account.Name, code)
		fmt.Printf("Este código foi copiado com sucesso e é válido por %d %s.\n", secondsRemaining, utils.Pluralize("segundo", "segundos", secondsRemaining))
	},
}

func init() {
	rootCmd.AddCommand(copyCmd)
	copyCmd.Flags().StringVarP(&copyName, "name", "n", "", "Nome da conta (ex: gitlab)")
}
