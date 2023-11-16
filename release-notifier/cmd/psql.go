/*******************************************************************************
 * Copyright (c) 2023 Contributors to the Eclipse Foundation
 *
 * See the NOTICE file(s) distributed with this work for additional
 * information regarding copyright ownership.
 *
 * This program and the accompanying materials are made available under the
 * terms of the Apache License, Version 2.0 which is available at
 * https://www.apache.org/licenses/LICENSE-2.0.
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
 * License for the specific language governing permissions and limitations
 * under the License.
 *
 * SPDX-License-Identifier: Apache-2.0
 ******************************************************************************/
package cmd

import (
	"log"
	"os"
	"github.com/spf13/cobra"
	"release-notifier/internal/psql"
)

// psqlCmd represents the psql command
var psqlCmd = &cobra.Command{
	Use:   "psql",
	Short: "Sends out notification of new PostgresSQL release",
	Long: `Sends out notification of new PostgresSQL release
to a mailinglist:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("PostgresSQL release notifier called.")
		psqlNotifier()
	},
}

func init() {
	rootCmd.AddCommand(psqlCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// psqlCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// psqlCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func psqlNotifier() {
	latestRelease := psql.GetLatestRel()
	prevRelease := psql.GetPrevRelFromArtifact()

	if latestRelease != prevRelease {
		log.Printf("New release is out: %v\n", latestRelease)
		alignedRelease := os.Getenv("CURRENT_ALIGNED_PSQL_VER")
		log.Printf("Current aligned version: %v\n", alignedRelease)
		psql.Notify(latestRelease, alignedRelease)
		psql.SaveLatestRel(latestRelease)
	} else {
		log.Println("No new release :(")
	}
}