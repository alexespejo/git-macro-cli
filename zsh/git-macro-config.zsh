#!/bin/zsh

# Print a green colored terminal welcome message
echo -e "\033[35m📚 ✨ 🐛 🎨 ♻️ 🔧 🚧 📚 ✨ 🐛 🎨 ♻️ 🔧 🚧 📚 ✨\033[0m"
echo
echo -e "\033[1;32mLet's setup to Git Macros! \033[0m"

read "use_emojis?Use icons in your commit messages? 📚 (y/n): "
if [[ "$use_emojis" == "y" ]]; then
  echo "Good to go! 🚀"
  echo "Setting up ..."
  git config alias.feat '!sh -c "git commit -m \"feat: ✨ $*\""'
  git config alias.fix '!sh -c "git commit -m \"fix: 🐛 $*\""'
  git config alias.docs '!sh -c "git commit -m \"docs: 📚 $*\""'
  git config alias.style '!sh -c "git commit -m \"style: 🎨 $*\""'
  git config alias.refactor '!sh -c "git commit -m \"refactor: ♻️ $*\""'
  git config alias.test '!sh -c "git commit -m \"test: ✅ $*\""'
  git config alias.chore '!sh -c "git commit -m \"chore: 🔧 $*\""'
  git config alias.wip '!sh -c "git commit -m \"wip: 🚧 $*\""'
  git config alias.afeat '!sh -c "git commit -am \"feat: ✨ $*\""'
  git config alias.afix '!sh -c "git commit -am \"fix: 🐛 $*\""'
  git config alias.adocs '!sh -c "git commit -am \"docs: 📚 $*\""'
  git config alias.astyle '!sh -c "git commit -am \"style: 🎨 $*\""'
  git config alias.arefactor '!sh -c "git commit -am \"refactor: ♻️ $*\""'
  git config alias.atest '!sh -c "git commit -am \"test: ✅ $*\""'
  git config alias.achore '!sh -c "git commit -am \"chore: 🔧 $*\""'
  git config alias.awip '!sh -c "git commit -am \"wip: 🚧 $*\""'

  
  echo -e "\033[35mThe following git macros have been created:\033[0m"
  echo "feat:     git commit -m \"feat: ✨\""
  echo "fix:      git commit -m \"fix: 🐛\""
  echo "docs:     git commit -m \"docs: 📚\""
  echo "style:    git commit -m \"style: 🎨\""
  echo "refactor: git commit -m \"refactor: ♻️\""
  echo "test:     git commit -m \"test: ✅\""
  echo "chore:    git commit -m \"chore: 🔧\""
  echo "wip:      git commit -m \"wip: 🚧\""
else
  echo "Good to go!"
  echo "Setting up ..."
  git config alias.feat '!sh -c "git commit -m \"feat: $*\""'
  git config alias.fix '!sh -c "git commit -m \"fix: $*\""'
  git config alias.docs '!sh -c "git commit -m \"docs: $*\""'
  git config alias.style '!sh -c "git commit -m \"style: $*\""'
  git config alias.refactor '!sh -c "git commit -m \"refactor: $*\""'
  git config alias.test '!sh -c "git commit -m \"test: $*\""'
  git config alias.chore '!sh -c "git commit -m \"chore: $*\""'
  git config alias.wip '!sh -c "git commit -m \"wip: $*\""'
  git config alias.afeat '!sh -c "git commit -am \"feat: $*\""'
  git config alias.afix '!sh -c "git commit -am \"fix: $*\""'
  git config alias.adocs '!sh -c "git commit -am \"docs: $*\""'
  git config alias.astyle '!sh -c "git commit -am \"style: $*\""'
  git config alias.arefactor '!sh -c "git commit -am \"refactor: $*\""'
  git config alias.atest '!sh -c "git commit -am \"test: $*\""'
  git config alias.achore '!sh -c "git commit -am \"chore: $*\""'
  git config alias.awip '!sh -c "git commit -am \"wip: $*\""'

  
  echo -e "\033[35mThe following git macros have been created:\033[0m"
  echo "feat:     git commit -m \"feat: \""
  echo "fix:      git commit -m \"fix: \""
  echo "docs:     git commit -m \"docs: \""
  echo "style:    git commit -m \"style: \""
  echo "refactor: git commit -m \"refactor: \""
  echo "test:     git commit -m \"test: \""
  echo "chore:    git commit -m \"chore: \""
  echo "wip:      git commit -m \"wip: \""
fi