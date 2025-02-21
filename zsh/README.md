# zsh

```sh
git config --global alias.feat '!sh -c "git commit -m \"feat: ✨ $*\""'
git config --global alias.fix '!sh -c "git commit -m \"fix: 🐛 $*\""'
git config --global alias.docs '!sh -c "git commit -m \"docs: 📚 $*\""'
git config --global alias.style '!sh -c "git commit -m \"style: 🎨 $*\""'
git config --global alias.refactor '!sh -c "git commit -m \"refactor: ♻️ $*\""'
git config --global alias.test '!sh -c "git commit -m \"test: ✅ $*\""'
git config --global alias.chore '!sh -c "git commit -m \"chore: 🔧 $*\""'
git config --global alias.wip '!sh -c "git commit -m \"wip: 🚧 $*\""'

git config --global alias.afeat '!sh -c "git commit -am \"feat: ✨ $*\""'
git config --global alias.afix '!sh -c "git commit -am \"fix: 🐛 $*\""'
git config --global alias.adocs '!sh -c "git commit -am \"docs: 📚 $*\""'
git config --global alias.astyle '!sh -c "git commit -am \"style: 🎨 $*\""'
git config --global alias.arefactor '!sh -c "git commit -am \"refactor: ♻️ $*\""'
git config --global alias.atest '!sh -c "git commit -am \"test: ✅ $*\""'
git config --global alias.achore '!sh -c "git commit -am \"chore: 🔧 $*\""'
git config --global alias.awip '!sh -c "git commit -am \"wip: 🚧 $*\""'

git config --global alias.bout "checkout -b"
git config --global alias.upush "!git push --set-upstream origin $(git rev-parse --abbrev-ref HEAD)"
```

## No emoji

```sh
git config --global alias.feat '!sh -c "git commit -m \"feat: $*\""'
git config --global alias.fix '!sh -c "git commit -m \"fix: $*\""'
git config --global alias.docs '!sh -c "git commit -m \"docs: $*\""'
git config --global alias.style '!sh -c "git commit -m \"style: $*\""'
git config --global alias.refactor '!sh -c "git commit -m \"refactor: $*\""'
git config --global alias.test '!sh -c "git commit -m \"test: $*\""'
git config --global alias.chore '!sh -c "git commit -m \"chore: $*\""'
git config --global alias.wip '!sh -c "git commit -m \"wip: $*\""'

git config --global alias.afeat '!sh -c "git commit -am \"feat: $*\""'
git config --global alias.afix '!sh -c "git commit -am \"fix: $*\""'
git config --global alias.adocs '!sh -c "git commit -am \"docs: $*\""'
git config --global alias.astyle '!sh -c "git commit -am \"style: $*\""'
git config --global alias.arefactor '!sh -c "git commit -am \"refactor: $*\""'
git config --global alias.atest '!sh -c "git commit -am \"test: $*\""'
git config --global alias.achore '!sh -c "git commit -am \"chore: $*\""'
git config --global alias.awip '!sh -c "git commit -am \"wip: $*\""'

git config --global alias.bout "checkout -b"
git config --global alias.upush "!git push --set-upstream origin $(git rev-parse --abbrev-ref HEAD)"
```

## Basic Commit

```sh
git config --global alias.com '!zsh -c "git commit -m \"$1\""'
git config --global alias.acom '!zsh -c "git commit -am \"$1\""'
```

## Quick Push (commit and push)

```sh
git config --global alias.qush '!zsh -c "git commit -am \"$1\" && git push --set-upstream origin $(git rev-parse --abbrev-ref HEAD)"'
```

## 💰🌿💥, jump to dev

_Require configuring for whatever your "develop" branch is called_

```sh
git config --global alias.bfd "!f() { git stash push -m 'bfd stash' && git switch <develop-branch> && git switch -c \"$1\" && git stash pop; }; f"
git config alias.jd 'checkout <develop-branch>'
```
