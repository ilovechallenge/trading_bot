package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/c9s/bbgo/pkg/data/tsv"
	"github.com/c9s/bbgo/pkg/optimizer"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"io"
	"io/ioutil"
	"os"
	"time"
)

func init() {
	optimizeExCmd.Flags().String("optimizer-config", "optimizer.yaml", "config file")
	optimizeExCmd.Flags().String("name", "", "assign a name to the study")
	optimizeExCmd.Flags().Bool("json-keep-all", false, "keep all results of trials")
	optimizeExCmd.Flags().String("output", "output", "backtest report output directory")
	optimizeExCmd.Flags().Bool("json", false, "print optimizer metrics in json format")
	optimizeExCmd.Flags().Bool("tsv", false, "print optimizer metrics in csv format")
	RootCmd.AddCommand(optimizeExCmd)
}

var optimizeExCmd = &cobra.Command{
	Use:   "optimizeex",
	Short: "run hyperparameter optimizer (experimental)",

	// SilenceUsage is an option to silence usage when an error occurs.
	SilenceUsage: true,

	RunE: func(cmd *cobra.Command, args []string) error {
		optimizerConfigFilename, err := cmd.Flags().GetString("optimizer-config")
		if err != nil {
			return err
		}

		configFile, err := cmd.Flags().GetString("config")
		if err != nil {
			return err
		}

		studyName, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}

		jsonKeepAll, err := cmd.Flags().GetBool("json-keep-all")
		if err != nil {
			return err
		}

		outputDirectory, err := cmd.Flags().GetString("output")
		if err != nil {
			return err
		}

		printJsonFormat, err := cmd.Flags().GetBool("json")
		if err != nil {
			return err
		}

		printTsvFormat, err := cmd.Flags().GetBool("tsv")
		if err != nil {
			return err
		}

		yamlBody, err := ioutil.ReadFile(configFile)
		if err != nil {
			return err
		}
		var obj map[string]interface{}
		if err := yaml.Unmarshal(yamlBody, &obj); err != nil {
			return err
		}
		delete(obj, "notifications")
		delete(obj, "sync")

		optConfig, err := optimizer.LoadConfig(optimizerConfigFilename)
		if err != nil {
			return err
		}

		// the config json template used for patch
		configJson, err := json.MarshalIndent(obj, "", "  ")
		if err != nil {
			return err
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		_ = ctx

		if len(studyName) == 0 {
			studyName = fmt.Sprintf("bbgo-hpopt-%v", time.Now().UnixMilli())
		}
		tempDirNameFormat := fmt.Sprintf("%s-config-*", studyName)
		configDir, err := os.MkdirTemp("", tempDirNameFormat)
		if err != nil {
			return err
		}

		executor := &optimizer.LocalProcessExecutor{
			Config:    optConfig.Executor.LocalExecutorConfig,
			Bin:       os.Args[0],
			WorkDir:   ".",
			ConfigDir: configDir,
			OutputDir: outputDirectory,
		}

		optz := &optimizer.HyperparameterOptimizer{
			StudyName: studyName,
			Config:    optConfig,
		}

		if err := executor.Prepare(configJson); err != nil {
			return err
		}

		report, err := optz.Run(executor, configJson)
		if err != nil {
			return err
		}

		if printJsonFormat {
			if !jsonKeepAll {
				report.Trials = nil
			}
			out, err := json.MarshalIndent(report, "", "  ")
			if err != nil {
				return err
			}

			// print report JSON to stdout
			fmt.Println(string(out))
		} else if printTsvFormat {
			if err := formatResultsTsv(os.Stdout, report.Parameters, report.Trials); err != nil {
				return err
			}
		} else {
			color.Green("OPTIMIZER REPORT")
			color.Green("===============================================\n")
			color.Green("STUDY NAME: %s\n", report.Name)
			color.Green("OPTIMIZE OBJECTIVE: %s\n", report.Objective)
			color.Green("BEST OBJECTIVE VALUE: %s\n", report.Best.Value)
			color.Green("OPTIMAL PARAMETERS:")
			for param, val := range report.Best.Parameters {
				color.Green("  - %s: %v", param, val)
			}
		}

		return nil
	},
}

func formatResultsTsv(writer io.WriteCloser, labelPaths map[string]string, results []*optimizer.HyperparameterOptimizeTrialResult) error {
	headerLen := len(labelPaths)
	headers := make([]string, 0, headerLen)
	for label := range labelPaths {
		headers = append(headers, label)
	}

	rows := make([][]interface{}, len(labelPaths))
	for ri, result := range results {
		row := make([]interface{}, headerLen)
		for ci, columnKey := range headers {
			var ok bool
			if row[ci], ok = result.Parameters[columnKey]; !ok {
				return fmt.Errorf(`missing parameter "%s" from trial result (%v)`, columnKey, result.Parameters)
			}
		}
		rows[ri] = row
	}

	w := tsv.NewWriter(writer)
	if err := w.Write(headers); err != nil {
		return err
	}

	for _, row := range rows {
		var cells []string
		for _, o := range row {
			cell, err := castCellValue(o)
			if err != nil {
				return err
			}
			cells = append(cells, cell)
		}

		if err := w.Write(cells); err != nil {
			return err
		}
	}
	return w.Close()
}
