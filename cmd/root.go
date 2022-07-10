package cmd

import (
	"os"

	"github.com/jcbritobr/james/dengine"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:     "James CLI",
		Version: "0.0.1",
		Long:    "James is a CLI for create gnome desktop launchers",
		Run: func(cmd *cobra.Command, args []string) {

			setupField := func(flag, field, value string, engine *dengine.DesktopData) {
				isChanged := func(flag string) bool {
					return cmd.Flags().Lookup(flag).Changed
				}
				if isChanged(flag) {
					engine.AddField(field, field)
				}
			}

			enBuilder := dengine.NewDesktopData()
			enBuilder.AddField(dengine.NameEntry, opName)
			enBuilder.AddField(dengine.TerminalEntry, opTerminal)
			enBuilder.AddField(dengine.ExecEntry, opExec)
			enBuilder.AddField(dengine.TypeEntry, opType)
			setupField("comment", opComment, dengine.CommentEntry, enBuilder)
			setupField("icon", opIcon, dengine.IconEntry, enBuilder)

			data := enBuilder.BuildFileData()
			file, err := os.Create(opFilename)

			if err != nil {
				panic(err)
			}

			file.WriteString(data)
		},
	}

	opName     string
	opExec     string
	opComment  string
	opType     string
	opTerminal string
	opIcon     string
	opFilename string
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&opName, "name", "n", "", "The --name -n sets the name of desktop file")
	rootCmd.PersistentFlags().StringVarP(&opExec, "exec", "e", "", "The --exec -e sets the executable entry on desktop file")
	rootCmd.PersistentFlags().StringVarP(&opComment, "comment", "c", "default", "The --comment -c sets comments on desktop file")
	rootCmd.PersistentFlags().StringVarP(&opType, "type", "t", "", "The --type -t sets the type of desktop launcher")
	rootCmd.PersistentFlags().StringVarP(&opTerminal, "terminal", "r", "", "The --terminal -r sets if application will or not run in a terminal(true/false)")
	rootCmd.PersistentFlags().StringVarP(&opIcon, "icon", "i", "default", "The --icon -i sets the icon of desktop file")
	rootCmd.PersistentFlags().StringVarP(&opFilename, "filename", "f", "", "The --filename -f creates the <filename>.desktop")

	rootCmd.MarkPersistentFlagRequired("name")
	rootCmd.MarkPersistentFlagRequired("exec")
	rootCmd.MarkPersistentFlagRequired("filename")
	rootCmd.MarkPersistentFlagRequired("terminal")
	rootCmd.MarkPersistentFlagRequired("type")
}

func Execute() error {
	return rootCmd.Execute()
}
