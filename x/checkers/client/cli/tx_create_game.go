package cli

import (
	"fmt"
	"strconv"

	"github.com/alice/checkers/x/checkers/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCreateGame() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-game [red] [black]",
		Short: "Broadcast message createGame",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argRed := args[0]
			argBlack := args[1]
			fmt.Println(argRed)

			clientCtx, err := client.GetClientTxContext(cmd)
			fmt.Println("clientCtx")
			fmt.Println(clientCtx)
			if err != nil {
				fmt.Println("early err")
				fmt.Println(err)
				return err
			}

			msg := types.NewMsgCreateGame(
				clientCtx.GetFromAddress().String(),
				argRed,
				argBlack,
			)
			fmt.Println("msg")
			fmt.Println(msg)
			if err := msg.ValidateBasic(); err != nil {
				fmt.Println("err")
				fmt.Println(err)
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
