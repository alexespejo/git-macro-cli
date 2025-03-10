if git show-ref --verify --quiet refs/heads/$1; then
  echo "Branch '$1' exists. Use a different branch name before proceeding"
  exit 1
else
  git stash push -m "Stash before deleting branch '$1'"
  git checkout -b $1
  git stash pop
  exit 2
fi


