---

# Git Macros

```sh
git config --global alias.feat "!f() { git commit -m \"✨ feat: $*\"; }; f"
git config --global alias.fix "!f() { git commit -m \"🐞 fix: $*\"; }; f"
git config --global alias.docs "!f() { git commit -m \"📚 docs: $*\"; }; f"
git config --global alias.style "!f() { git commit -m \"🎨 style: $*\"; }; f"
git config --global alias.refactor "!f() { git commit -m \"♻️ refactor: $*\"; }; f"
git config --global alias.test "!f() { git commit -m \"✅ test: $*\"; }; f"
git config --global alias.chore "!f() { git commit -m \"🧹 chore: $*\"; }; f"
git config --global alias.afeat "!f() { git commit -am \"✨ feat: $*\"; }; f"
git config --global alias.afix "!f() { git commit -m \"🐞 fix: $*\"; }; f"
git config --global alias.adocs "!f() { git commit -m \"📚 docs: $*\"; }; f"
git config --global alias.astyle "!f() { git commit -m \"🎨 style: $*\"; }; f"
git config --global alias.arefactor "!f() { git commit -m \"♻️ refactor: $*\"; }; f"
git config --global alias.atest "!f() { git commit -m \"✅ test: $*\"; }; f"
git config --global alias.achore "!f() { git commit -m \"🧹 chore: $*\"; }; f"

alias.bout checkout -b
alias.upush !git push --set-upstream origin $(git rev-parse --abbrev-ref HEAD)
```
