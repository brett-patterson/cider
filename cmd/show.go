package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/brett-patterson/cider/lib"
	"github.com/spf13/cobra"
	"net"
	"os"
	"text/template"
)

var describeTemplate *template.Template
var describeTemplateString = `{{bold .CIDR}}
Subnet Mask: {{.SubnetMask}}
IP Range:    {{.FirstIp}} - {{.LastIp}}
`

type CIDRDescription struct {
	CIDR       string
	SubnetMask string
	FirstIp    net.IP
	LastIp     net.IP
}

func describeCIDRBlock(c *net.IPNet) *CIDRDescription {
	return &CIDRDescription{
		CIDR:       c.String(),
		SubnetMask: lib.SubnetMask(c.Mask),
		FirstIp:    c.IP,
		LastIp:     lib.LastIp(c),
	}
}

var showCmd = &cobra.Command{
	Use:   "show cidr [cidr] ...",
	Short: "Describe CIDR blocks",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		cidrBlocks := make([]*net.IPNet, 0)

		for _, arg := range args {
			if c, err := lib.ParseCIDRBlock(arg); err == nil {
				cidrBlocks = append(cidrBlocks, c)
			} else {
				return err
			}
		}

		descriptions := make([]*CIDRDescription, len(cidrBlocks))
		for i, c := range cidrBlocks {
			descriptions[i] = describeCIDRBlock(c)
		}

		switch OutputType {
		case Text:
			for i, d := range descriptions {
				if err := describeTemplate.Execute(os.Stdout, d); err != nil {
					return err
				}

				if i < len(descriptions)-1 {
					fmt.Println()
				}
			}
			return nil
		case Json:
			return json.NewEncoder(os.Stdout).Encode(descriptions)
		}

		return errors.New("Unsupported output type: " + OutputType)
	},
}

func init() {
	rootCmd.AddCommand(showCmd)

	describeTemplate = lib.ParseTemplate("DescribeCIDRBlock", describeTemplateString)
}
