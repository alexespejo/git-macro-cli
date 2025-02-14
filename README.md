---

# Git Macros

These Git aliases simplify common tasks. They save time by allowing shorthand commands for frequent actions like committing, branching, and pushing.

## Setup

Run the following commands to set up the aliases:

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
```

---

## Available Aliases

### Branching & Stashing

- **`bfd <branch-name>`**  
  Stash changes, switch to `develop`, create a new branch, set upstream to the new branch, and apply the stash.  
  **Usage:**  
  ```sh
  git bfd my-feature-branch
  ```

- **`jd`**  
  Switch to the `develop` branch.  
  **Usage:**  
  ```sh
  git jd
  ```

### Commits

- **`feat <message>`**  
  Create a commit for a new feature with the emoji `✨`.  
  **Usage:**  
  ```sh
  git feat "added search functionality"
  ```

- **`bug <message>`**  
  Create a commit to fix a bug with the emoji `🐛`.  
  **Usage:**  
  ```sh
  git bug "fixed login issue"
  ```

- **`docs <message>`**  
  Create a commit for documentation updates with the emoji `📚`.  
  **Usage:**  
  ```sh
  git docs "updated README"
  ```

- **`style <message>`**  
  Create a commit for style changes with the emoji `🎨`.  
  **Usage:**  
  ```sh
  git style "formatted the code"
  ```

- **`refactor <message>`**  
  Create a commit for code refactoring with the emoji `♻️`.  
  **Usage:**  
  ```sh
  git refactor "optimized loop"
  ```

- **`test <message>`**  
  Create a commit for adding or fixing tests with the emoji `✅`.  
  **Usage:**  
  ```sh
  git test "added unit tests for login"
  ```

- **`chore <message>`**  
  Create a commit for chores or maintenance tasks with the emoji `🧹`.  
  **Usage:**  
  ```sh
  git chore "removed unused files"
  ```

### Commit with `-am` (Auto-Stages)

- **`afeat <message>`**  
  Create a commit for a new feature with the emoji `✨`, and automatically stage all modified files.  
  **Usage:**  
  ```sh
  git afeat "added new button design"
  ```

---
