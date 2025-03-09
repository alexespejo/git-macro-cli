package main

import (
	"bufio"
	color "cli/utils/colors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/AlecAivazis/survey/v2"
)

const (
	YES        = "y"
	NO         = "n"
	ERROR_FLAG = "error"
)

func colorize(text string, styles ...string) string {
	return strings.Join(styles, "") + text + color.RESET
}
func printStyled(text string, styles ...string) {
	fmt.Print(colorize(text, styles...))
}

func printlnStyled(text string, styles ...string) {
	fmt.Println(colorize(text, styles...))
}
func promptYesNo(prompt func()) string {
	for {
		reader := bufio.NewReader(os.Stdin)
		prompt()
		param, _ := reader.ReadString('\n')
		param = strings.ToLower(strings.TrimRight(param, "\n"))

		if param == YES {
			return YES
		} else if param == NO {
			return NO
		}
	}
}
func printDelay(text string) {
	time.Sleep(50 * time.Millisecond)
	fmt.Println(text)
}
func mnuConfigCommits() bool {
	param := promptYesNo(func() {
		printStyled("Do you want to use emojis in your commit messages? ", color.MAGENTA, color.BOLD)
		printStyled("(y/n)", color.MAGENTA)
	})

	printlnStyled("You got it! 😎", color.MAGENTA)
	cmd := exec.Command("/bin/zsh", "git.config.commits.zsh", param)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()

	status := cmd.ProcessState.ExitCode()
	if status == 1 {
		fmt.Println("")
		if param == YES {
			printDelay(colorize("git feat     -> ", color.BOLD, color.GREEN) + "feat: ✨")
			printDelay(colorize("git fix      -> ", color.BOLD, color.GREEN) + "fix: 🪲")
			printDelay(colorize("git docs     -> ", color.BOLD, color.GREEN) + "docs: 📚")
			printDelay(colorize("git style    -> ", color.BOLD, color.GREEN) + "style: 🎨")
			printDelay(colorize("git refactor -> ", color.BOLD, color.GREEN) + "refactor: ♻️")
			printDelay(colorize("git test     -> ", color.BOLD, color.GREEN) + "test: ✅")
			printDelay(colorize("git chore    -> ", color.BOLD, color.GREEN) + "chore: 🔧")
			printDelay(colorize("git wip      -> ", color.BOLD, color.GREEN) + "wip: 🚧")
		} else if param == NO {
			printDelay(colorize("git feat     -> ", color.BOLD, color.GREEN) + "feat")
			printDelay(colorize("git fix      -> ", color.BOLD, color.GREEN) + "fix")
			printDelay(colorize("git docs     -> ", color.BOLD, color.GREEN) + "docs")
			printDelay(colorize("git style    -> ", color.BOLD, color.GREEN) + "style")
			printDelay(colorize("git refactor -> ", color.BOLD, color.GREEN) + "refactor")
			printDelay(colorize("git test     -> ", color.BOLD, color.GREEN) + "test")
			printDelay(colorize("git chore    -> ", color.BOLD, color.GREEN) + "chore")
			printDelay(colorize("git wip      -> ", color.BOLD, color.GREEN) + "wip")
		}
		printlnStyled("Your commit macros have been created!", color.MAGENTA)
		return true
	}

	printlnStyled("Do you want to try again?", color.MAGENTA)
	return false
}

func surveyPrompt(options []string, message string, defaultVal string) string {
	fmt.Println("")
	var selectedTemplate string
	surveyPrompt := &survey.Select{
		Message: message,
		Options: options,
		Default: defaultVal,
	}

	// Ask the prompt
	err := survey.AskOne(surveyPrompt, &selectedTemplate)
	if err != nil {
		fmt.Println("Error: ", err)
		return ERROR_FLAG
	}

	return selectedTemplate
}

func main() {

	printlnStyled("===================================", color.CYAN, color.BOLD)
	printlnStyled(" 	ABE's Git Macro CLI ", color.CYAN, color.BOLD)
	printlnStyled("===================================", color.CYAN, color.BOLD)
	option1 := colorize("Commits ✨", color.UNDERLINE, color.ORANGE)
	option2 := colorize("Branching 🌳", color.UNDERLINE, color.GREEN)
	option3 := colorize("Other ⚙️\n", color.UNDERLINE, color.BLUE)
	option4 := colorize("Cancel️", color.UNDERLINE)
	defaultVal := option1
	for {
		options := []string{option1, option2, option3, option4}
		res := surveyPrompt(options, "Choose what to configure:", defaultVal)

		if res == ERROR_FLAG {
			fmt.Println(colorize("Exiting...", color.RED))
			return
		} else if res == option4 {
			return
		}

		fmt.Printf("You selected the %s template.\n", res)
		switch res {
		case option1:
			mnuConfigCommits()
		}
		defaultVal = option4
	}
}
