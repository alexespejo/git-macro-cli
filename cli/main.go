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

func mnuConfigCommits() bool {
	option1 := colorize("- Yes Emoji 👍", color.MAGENTA)
	option2 := colorize("- No Emoji 👎", color.MAGENTA)
	option3 := colorize("- Back")
	req := surveyPrompt([]string{option1, option2, option3}, "Do you want to use emojis in your commits?", option1)

	param := ""
	switch req {
	case option1:
		param = YES
	case option2:
		param = NO
	default:
		return false
	}
	printlnStyled("You got it! 😎", color.MAGENTA)
	cmd := exec.Command("./scripts/git-config-commits.zsh", param)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()

	status := cmd.ProcessState.ExitCode()
	fmt.Println(status)
	if status == 1 {
		fmt.Println("")
		if param == YES {
			printDelay(colorize("git feat     → ", color.BOLD, color.GREEN) + "feat: ✨")
			printDelay(colorize("git fix      → ", color.BOLD, color.GREEN) + "fix: 🪲")
			printDelay(colorize("git docs     → ", color.BOLD, color.GREEN) + "docs: 📚")
			printDelay(colorize("git style    → ", color.BOLD, color.GREEN) + "style: 🎨")
			printDelay(colorize("git refactor → ", color.BOLD, color.GREEN) + "refactor: ♻️")
			printDelay(colorize("git test     → ", color.BOLD, color.GREEN) + "test: ✅")
			printDelay(colorize("git chore    → ", color.BOLD, color.GREEN) + "chore: 🔧")
			printDelay(colorize("git wip      → ", color.BOLD, color.GREEN) + "wip: 🚧")
		} else if param == NO {
			printDelay(colorize("git feat     → ", color.BOLD, color.GREEN) + "feat")
			printDelay(colorize("git fix      → ", color.BOLD, color.GREEN) + "fix")
			printDelay(colorize("git docs     → ", color.BOLD, color.GREEN) + "docs")
			printDelay(colorize("git style    → ", color.BOLD, color.GREEN) + "style")
			printDelay(colorize("git refactor → ", color.BOLD, color.GREEN) + "refactor")
			printDelay(colorize("git test     → ", color.BOLD, color.GREEN) + "test")
			printDelay(colorize("git chore    → ", color.BOLD, color.GREEN) + "chore")
			printDelay(colorize("git wip      → ", color.BOLD, color.GREEN) + "wip")
		}
		printlnStyled("Your commit macros have been created!", color.MAGENTA)
		return true
	}

	printlnStyled("Do you want to try again?", color.MAGENTA)
	return false
}

func main() {
	printlnStyled("===================================", color.CYAN, color.BOLD)
	printlnStyled(" 	ABE Git Macro CLI ", color.CYAN, color.BOLD)
	printlnStyled("===================================", color.CYAN, color.BOLD)
	option1 := colorize("- ☄️ Commits", color.ORANGE)
	option2 := colorize("- 🌱 Branching", color.GREEN)
	option3 := colorize("- 💫 Special", color.YELLOW)
	option4 := colorize("- ⚙️ Other", color.BLUE)
	optionCancel := colorize("Close", color.UNDERLINE)
	defaultVal := option1

	for {
		options := []string{option1, option2, option3, option4, optionCancel}
		res := surveyPrompt(options, "Choose what to configure:", defaultVal)

		if res == ERROR_FLAG {
			fmt.Println(colorize("Exiting...", color.RED))
			return
		} else if res == optionCancel {
			return
		}

		switch res {
		case option1:
			mnuConfigCommits()
		}
		defaultVal = optionCancel
	}
}
