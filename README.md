---
testing upush
# Git Macros

| Command | Resulting Commit Message |
|---------|-------------------------|
| `git feat "commit message"` | ✨ feat: commit message |
| `git fix "commit message"` | 🐞 fix: commit message |
| `git docs "commit message"` | 📚 docs: commit message |
| `git style "commit message"` | 🎨 style: commit message |
| `git refactor "commit message"` | ♻️ refactor: commit message |
| `git test "commit message"` | ✅ test: commit message |
| `git chore "commit message"` | 🧹 chore: commit message |
| `git wip "commit message"` | 🚧 wip: commit message |

## commit, commit -am, checkout -b, push -u
```sh
git config --global alias.feat "!f() { git commit -m \"✨ feat: $*\"; }; f"
git config --global alias.fix "!f() { git commit -m \"🐞 fix: $*\"; }; f"
git config --global alias.docs "!f() { git commit -m \"📚 docs: $*\"; }; f"
git config --global alias.style "!f() { git commit -m \"🎨 style: $*\"; }; f"
git config --global alias.refactor "!f() { git commit -m \"♻️ refactor: $*\"; }; f"
git config --global alias.test "!f() { git commit -m \"✅ test: $*\"; }; f"
git config --global alias.chore "!f() { git commit -m \"🧹 chore: $*\"; }; f"
git config --global alias.wip "!f() { git commit -m \"🚧 wip: $*\"; }; f"

git config --global alias.afeat "!f() { git commit -am \"✨ feat: $*\"; }; f"
git config --global alias.afix "!f() { git commit -am \"🐞 fix: $*\"; }; f"
git config --global alias.adocs "!f() { git commit -am \"📚 docs: $*\"; }; f"
git config --global alias.astyle "!f() { git commit -am \"🎨 style: $*\"; }; f"
git config --global alias.arefactor "!f() { git commit -am \"♻️ refactor: $*\"; }; f"
git config --global alias.atest "!f() { git commit -am \"✅ test: $*\"; }; f"
git config --global alias.achore "!f() { git commit -am \"🧹 chore: $*\"; }; f"
git config --global alias.awip "!f() { git commit -am \"🚧 wip: $*\"; }; f"

git config --global alias.bout "checkout -b"
git config --global alias.upush "!git push --set-upstream origin $(git rev-parse --abbrev-ref HEAD)"
```

### No Emojis
```sh
git config --global alias.feat "!f() { git commit -m \"feat: $*\"; }; f"
git config --global alias.fix "!f() { git commit -m \"fix: $*\"; }; f"
git config --global alias.docs "!f() { git commit -m \"docs: $*\"; }; f"
git config --global alias.style "!f() { git commit -m \"style: $*\"; }; f"
git config --global alias.refactor "!f() { git commit -m \"refactor: $*\"; }; f"
git config --global alias.test "!f() { git commit -m \"test: $*\"; }; f"
git config --global alias.chore "!f() { git commit -m \"chore: $*\"; }; f"
git config --global alias.wip "!f() { git commit -m \"wip: $*\"; }; f"

git config --global alias.afeat "!f() { git commit -am \"feat: $*\"; }; f"
git config --global alias.afix "!f() { git commit -am \"fix: $*\"; }; f"
git config --global alias.adocs "!f() { git commit -am \"docs: $*\"; }; f"
git config --global alias.astyle "!f() { git commit -am \"style: $*\"; }; f"
git config --global alias.arefactor "!f() { git commit -am \"refactor: $*\"; }; f"
git config --global alias.atest "!f() { git commit -am \"test: $*\"; }; f"
git config --global alias.achore "!f() { git commit -am \"chore: $*\"; }; f"
git config --global alias.awip "!f() { git commit -am \"wip: $*\"; }; f"

git config --global alias.bout "checkout -b"
git config --global alias.upush "!git push --set-upstream origin $(git rev-parse --abbrev-ref HEAD)"

```

## zsh
```sh
# Aliases for `git commit -m` (commit without adding changes)
git config --global alias.feat '!sh -c "git commit -m \"feat: $*\""'
git config --global alias.fix '!sh -c "git commit -m \"fix: $*\""'
git config --global alias.docs '!sh -c "git commit -m \"docs: $*\""'
git config --global alias.style '!sh -c "git commit -m \"style: $*\""'
git config --global alias.refactor '!sh -c "git commit -m \"refactor: $*\""'
git config --global alias.test '!sh -c "git commit -m \"test: $*\""'
git config --global alias.chore '!sh -c "git commit -m \"chore: $*\""'
git config --global alias.wip '!sh -c "git commit -m \"wip: $*\""'

# Aliases for `git commit -am` (commit including all tracked changes)
git config --global alias.afeat '!sh -c "git commit -am \"feat: $*\""'
git config --global alias.afix '!sh -c "git commit -am \"fix: $*\""'
git config --global alias.adocs '!sh -c "git commit -am \"docs: $*\""'
git config --global alias.astyle '!sh -c "git commit -am \"style: $*\""'
git config --global alias.arefactor '!sh -c "git commit -am \"refactor: $*\""'
git config --global alias.atest '!sh -c "git commit -am \"test: $*\""'
git config --global alias.achore '!sh -c "git commit -am \"chore: $*\""'
git config --global alias.awip '!sh -c "git commit -am \"wip: $*\""'
```

## basic commit macro
```sh
#Example
git com "commit msg" # commit msg
git acom "commit msg" # commit msg
```


```sh
git config --global alias.com "!f() { git commit -m \"$*\"; }; f"
git config --global alias.acom "!f() { git commit -am \"$*\"; }; f"
```

## commit -am & push
```sh
# Example
git qush "commit-and-push"
```

```sh
git config --global alias.qush "!f() { git commit -am \"$*\" && git push --set-upstream origin $(git rev-parse --abbrev-ref HEAD); }; f"
```

## 💰🌿💥, jump to dev

_Require configuring for whatever your "develop" branch is called_
```sh
git config --global alias.bfd "!f() { git stash push -m 'bfd stash' && git switch <develop-branch> && git switch -c \"$1\" && git stash pop; }; f"
git config --global alias.jd "checkout <develop-branch>"
```
