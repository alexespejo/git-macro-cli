package main

import (
	"bufio"
	color "cli/utils/colors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const (
	YES = "y"
	NO  = "n"
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
func mnuConfigCommits() bool {
	param := promptYesNo(func() {
		printStyled("Do you want to use emojis in your commit messages? ", color.MAGENTA, color.BOLD)
		printStyled("(y/n)", color.MAGENTA)
	})

	cmd := exec.Command("/bin/zsh", "git.config.commits.zsh", param)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()

	status := cmd.ProcessState.ExitCode()
	if status == 1 {
		printlnStyled("Done! 🎉", color.MAGENTA)
		printlnStyled("The following macros have been created:", color.MAGENTA)
		if param == YES {
			printlnStyled("feat: ✨", color.BOLD, color.GREEN)
			printlnStyled("fix: 🐛", color.BOLD, color.GREEN)
			printlnStyled("docs: 📚", color.BOLD, color.GREEN)
			printlnStyled("style: 🎨", color.BOLD, color.GREEN)
			printlnStyled("refactor: ♻️", color.BOLD, color.GREEN)
			printlnStyled("test: ✅", color.BOLD, color.GREEN)
			printlnStyled("chore: 🔧", color.BOLD, color.GREEN)
			printlnStyled("wip: 🚧", color.BOLD, color.GREEN)
		} else if param == NO {
			printlnStyled("feat", color.BOLD, color.GREEN)
			printlnStyled("fix", color.BOLD, color.GREEN)
			printlnStyled("docs", color.BOLD, color.GREEN)
			printlnStyled("style", color.BOLD, color.GREEN)
			printlnStyled("refactor", color.BOLD, color.GREEN)
			printlnStyled("test", color.BOLD, color.GREEN)
			printlnStyled("chore", color.BOLD, color.GREEN)
			printlnStyled("wip", color.BOLD, color.GREEN)
		}
		return true
	}

	printlnStyled("Do you want to try again?", color.MAGENTA)
	return false
}

func main() {
	mnuConfigCommits()
}
