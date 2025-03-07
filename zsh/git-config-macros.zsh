#!/bin/zsh
  show_loading_animation() {
    echo -ne "Setting up "
    for i in {1..5}; do
      echo -ne "."
      sleep 0.5
    done
    echo
  }
# Print a green colored terminal welcome message
echo -e "\033[1;32mLet's setup to Git Macros!\033[0m"

read "use_emojis?Use icons in your commit messages? 📚 (y/n): "
if [[ "$use_emojis" == "y" ]]; then
  echo "Good to go! ⚙️"

  show_loading_animation
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

  
  echo
  echo -e "\033[35mDone! 🎉\033[0m"
  echo -e "\033[35mThe following git macros have been created:\033[0m"
  echo -e "\033[1;32mfeat:\033[0m     git commit -m \"feat: ✨\""
  echo -e "\033[1;32mfix:\033[0m      git commit -m \"fix: 🐛\""
  echo -e "\033[1;32mdocs:\033[0m     git commit -m \"docs: 📚\""
  echo -e "\033[1;32mstyle:\033[0m    git commit -m \"style: 🎨\""
  echo -e "\033[1;32mrefactor:\033[0m git commit -m \"refactor: ♻️\""
  echo -e "\033[1;32mtest:\033[0m     git commit -m \"test: ✅\""
  echo -e "\033[1;32mchore:\033[0m    git commit -m \"chore: 🔧\""
  echo -e "\033[1;32mwip:\033[0m      git commit -m \"wip: 🚧\""
else
  echo "Good to go!"
  show_loading_animation
  
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

  echo 
  echo -e "\033[35mDone! \033[0m"
  echo -e "\033[35mThe following git macros have been created:\033[0m"
  echo -e "\033[1;32mfeat:\033[0m     git commit -m \"feat: \""
  echo -e "\033[1;32mfix:\033[0m      git commit -m \"fix: \""
  echo -e "\033[1;32mdocs:\033[0m     git commit -m \"docs: \""
  echo -e "\033[1;32mstyle:\033[0m    git commit -m \"style: \""
  echo -e "\033[1;32mrefactor:\033[0m git commit -m \"refactor: \""
  echo -e "\033[1;32mtest:\033[0m     git commit -m \"test: \""
  echo -e "\033[1;32mchore:\033[0m    git commit -m \"chore: \""
  echo -e "\033[1;32mwip:\033[0m      git commit -m \"wip: \""
fi
fi