/*
Copyright © 2019 Ken'ichiro Oyama <k1lowxb@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/k1LoW/tsplt/protter"
	"github.com/k1LoW/tsplt/timeseries"
	"github.com/k1LoW/tsplt/version"
	"github.com/mattn/go-isatty"
	"github.com/spf13/cobra"
)

var (
	inPath    string
	outPath   string
	delimiter string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tsplt",
	Short: "plot time series data",
	Long:  `plot time series data.`,
	Args: func(cmd *cobra.Command, args []string) error {
		versionVal, err := cmd.Flags().GetBool("version")
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}
		if versionVal {
			fmt.Println(version.Version)
			os.Exit(0)
		}
		if len(args) > 0 {
			_, _ = fmt.Fprintf(os.Stderr, "%s\n", errors.New("invalid argument"))
			os.Exit(1)
		}
		if inPath == "" && isatty.IsTerminal(os.Stdin.Fd()) {
			_, _ = fmt.Fprintf(os.Stderr, "%s\n", errors.New("tsplt need STDIN or input file"))
			os.Exit(1)
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		var (
			in  io.Reader
			err error
		)
		if inPath != "" {
			in, err = os.Open(filepath.Clean(inPath))
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "%s\n", err)
				os.Exit(1)
			}
		} else {
			in = os.Stdin
		}
		dr := []rune(delimiter)
		if len(dr) != 1 {
			_, _ = fmt.Fprintf(os.Stderr, "%s\n", errors.New("invalid delimiter"))
			os.Exit(1)
		}

		data, err := timeseries.Build(in, dr[0])
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}
		err = protter.Plot(data, outPath)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&inPath, "in", "i", "", "input timeseries data file path")
	rootCmd.Flags().StringVarP(&outPath, "out", "o", "tsplt.png", "output png file path")
	rootCmd.Flags().StringVarP(&delimiter, "delimiter", "d", "\t", "input file delimiter")
	rootCmd.Flags().BoolP("version", "v", false, "print the version")
}
