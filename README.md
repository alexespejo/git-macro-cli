---

# Git Macros

| Command                         | Resulting Commit Message    |
| ------------------------------- | --------------------------- |
| `git feat "commit message"`     | ✨ feat: commit message     |
| `git fix "commit message"`      | 🐞 fix: commit message      |
| `git docs "commit message"`     | 📚 docs: commit message     |
| `git style "commit message"`    | 🎨 style: commit message    |
| `git refactor "commit message"` | ♻️ refactor: commit message |
| `git test "commit message"`     | ✅ test: commit message     |
| `git chore "commit message"`    | 🧹 chore: commit message    |
| `git wip "commit message"`      | 🚧 wip: commit message      |

## basic commit macro

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
