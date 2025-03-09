#!/bin/bash
show_loading_animation() {
  frames=("." ".." "...")
   for frame in "${frames[@]}"; do
    printf "\r%s" "$frame"
    sleep 0.5
  done
  echo
}
if [ "$1" = "y" ]; then
  show_loading_animation
  git config alias.feat '!sh -c "git commit -m \"feat: ✨ $*\""'
  git config alias.fix '!sh -c "git commit -m \"fix: 🪲 $*\""'
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
  exit 1
else 
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
 exit 1
fi

exit 2