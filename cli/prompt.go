package cli

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"github.com/manifoldco/promptui"
)

type PromptContent struct {
	Label    string
	errorMsg string
}

type PromptOptions struct {
	Label string
	Key   int64
}

func GetTemplates() *promptui.SelectTemplates {
	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}",
		Active:   "\U0001F336 {{ .Label | yellow }} ",
		Inactive: "  {{ .Label | cyan }} ",
		Selected: "\U0001FAD1 {{ .Label | green | cyan }}",
	}
	return templates
}

func GetSearcher(options []PromptOptions) func(input string, index int) bool {
	searcher := func(input string, index int) bool {
		option := options[index]
		name := strings.Replace(strings.ToLower(option.Label), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)
		return strings.Contains(name, input)
	}
	return searcher
}

func Select(promptTitle string, options []PromptOptions) *promptui.Select {
	prompt := promptui.Select{
		Label:     promptTitle,
		Items:     options,
		Templates: GetTemplates(),
		Size:      4,
		Searcher:  GetSearcher(options),
	}
	return &prompt
}

func PromptGetInput(pc PromptContent) string {
	validate := func(input string) error {
		if len(input) <= 0 {
			return errors.New(pc.errorMsg)
		}
		return nil
	}
	prompt := promptui.Prompt{
		Label: pc.Label,
		// Templates: templates,
		Validate: validate,
	}
	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Input: %s\n", result)
	return result
}

func PromptDownload() bool {
	willDownload := PromptGetInput(
		PromptContent{Label: "Download the sandbox credentials file in plaintext? (yes/no)"})
	if willDownload == "yes" {
		fmt.Println("Downloading Sandbox Credentials...")
		return true
	} else if willDownload == "no" {
		fmt.Println("Not downloading...")
		return false
	} else {
		fmt.Println("Invalid Answer")
		PromptDownload()
	}
	return false
}

func PromptFileName() string {
	filename := PromptGetInput(PromptContent{Label: "File Name?"})
	return filename
}


func PromptConfig() bool {
	willAppend := PromptGetInput(
		PromptContent{Label: "Would you like to append the sandbox credentials file to your AWS config file? (yes/no)"})
	if willAppend == "yes" {
		fmt.Println("Appending Sandbox Credentials to AWS configs...")
		return true

	} else if willAppend == "no" {
		fmt.Println("Not Appending to AWS configs...")
		return false
	} else {
		fmt.Println("Invalid Answer")
		PromptConfig()
	}
	return false
}

func PromptRepoOwner() (owner string, err error) {
	owner = PromptGetInput(PromptContent{Label: "What is the name of the repository owner?"})
	if owner == "" {
		err = errors.New("please enter a valid repository owner")
		PromptRepoOwner()
	}
	return owner, err
}

func PromptRepoName() (repo string, err error) {
	repo = PromptGetInput(PromptContent{Label: "What is the name of the repository?"})
	if repo == "" {
		err = errors.New("please enter a valid repository name")
		PromptRepoName()
	}
	return repo, err
}

func PromptQuery() string {
	prompt := promptui.Prompt{
		Label: "Enter the query to run",
		// Templates: templates,
		// Validate: validate,
	}
	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Query Entered : %q\n", result)
	return result
}


func PromptString(label string) (string, error){
	validate := func(input string) error {
		if len(input) <= 0 {
			return errors.New("Please enter a valid string!")
		}
		return nil
	}
	prompt := promptui.Prompt{ Label: label, Validate: validate}
	result, err := prompt.Run()
	PrintIfErr(err)
	Success("Input: %s\n", result)
	return result, err
}