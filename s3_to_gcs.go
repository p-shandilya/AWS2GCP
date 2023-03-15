package main

import (
	"fmt"
	"os"
	"strings"
)

func translateCommand(s3Cmd string) string {
	s3Args := strings.Split(s3Cmd, " ")
	gcsCmd := ""

	switch s3Args[0] {
	case "aws":
		if len(s3Args) >= 3 && s3Args[1] == "s3" {
			switch s3Args[2] {
			case "ls":
				gcsCmd = "gsutil ls"
				if len(s3Args) > 3 {
					gcsCmd += " " + strings.Replace(s3Args[3], "s3://", "gs://", 1)
				}
			case "cp":
				if len(s3Args) > 4 {
					gcsCmd = fmt.Sprintf("gsutil cp %s %s", s3Args[3], strings.Replace(s3Args[4], "s3://", "gs://", 1))
				}
			case "mv":
				if len(s3Args) > 4 {
					gcsCmd = fmt.Sprintf("gsutil mv %s %s", s3Args[3], strings.Replace(s3Args[4], "s3://", "gs://", 1))
				}
			case "rm":
				if len(s3Args) > 3 {
					gcsCmd = "gsutil rm " + strings.Replace(s3Args[3], "s3://", "gs://", 1)
					if len(s3Args) > 4 && s3Args[4] == "--recursive" {
						gcsCmd += " -r"
					}
				}
			case "sync":
				if len(s3Args) > 4 {
					gcsCmd = fmt.Sprintf("gsutil rsync %s %s", s3Args[3], strings.Replace(s3Args[4], "s3://", "gs://", 1))
				}
			}
		}
	}
	return gcsCmd
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: s3togcs \"s3 command\"")
		os.Exit(1)
	}

	s3Cmd := os.Args[1]
	gcsCmd := translateCommand(s3Cmd)

	if gcsCmd != "" {
		fmt.Println("Equivalent GCS command:")
		fmt.Println(gcsCmd)
	} else {
		fmt.Println("Unable to translate the given S3 command.")
	}
}
