package brancher

import (
	"fmt"

	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

var NO_BRANCH_ERR = fmt.Errorf("branch doesnt exist")

func (M *Manager) BranchCheckout(name string) error {

	if _, ok := M.branchMap[name]; !ok {
		return NO_BRANCH_ERR
	}

	workTree, err := M.repository.Worktree()
	if err != nil {
		return fmt.Errorf("getting worktree: " + err.Error())
	}

	status, err := workTree.Status()
	if err != nil {
		return fmt.Errorf("getting status: " + err.Error())
	}

	if !status.IsClean() {
		return fmt.Errorf("Worktree is not clean. Manage the changes before checking out.")
	}

	branchRefName := plumbing.NewBranchReferenceName(name)
	coutOpts := git.CheckoutOptions{
		Branch: plumbing.ReferenceName(branchRefName),
	}

	if err = workTree.Checkout(&coutOpts); err != nil {
		return fmt.Errorf("checking out: " + err.Error())
	}
	return nil

}

func (M *Manager) BranchDelete(name string) error {
	if _, ok := M.branchMap[name]; !ok {
		return NO_BRANCH_ERR
	}

	branchRefName := plumbing.NewBranchReferenceName(name)
	err := M.repository.Storer.RemoveReference(plumbing.ReferenceName(branchRefName))
	if err != nil {
		return fmt.Errorf("removing reference: " + err.Error())
	}

	return nil
}

func (M *Manager) BranchCreate(name string, referenceHash ...plumbing.Hash) error {

	// Create a branch off of HEAD if not given

	headRef, err := M.repository.Head()
	if err != nil {
		return fmt.Errorf("getting head reference: " + err.Error())
	}
	hash := headRef.Hash()

	// If a hash is provided
	if len(referenceHash) > 0 {
		hash = referenceHash[0]
	}

	newBranchNameRef := plumbing.ReferenceName("refs/heads/" + name)
	newBranchHashRef := plumbing.NewHashReference(newBranchNameRef, hash)

	err = M.repository.Storer.SetReference(newBranchHashRef)
	if err != nil {
		return fmt.Errorf("setting new reference: " + err.Error())
	}

	return nil
}

func (M *Manager) BranchRename(oldName string, newName string) error {

	if _, ok := M.branchMap[oldName]; !ok {
		return NO_BRANCH_ERR
	}

	// Create a branch with newName and currentBranch hash

	branchRefName := plumbing.NewBranchReferenceName(oldName)
	oldReference, err := M.repository.Reference(branchRefName, true)
	if err != nil {
		return fmt.Errorf("getting old branch reference: " + err.Error())
	}

	oldRefereceHash := oldReference.Hash()
	err = M.BranchCreate(newName, oldRefereceHash)
	if err != nil {
		return fmt.Errorf("new branch creation: " + err.Error())
	}

	// Safely delete the old branch
	err = M.BranchDelete(oldName)
	if err != nil {
		return fmt.Errorf("old branch deletion: " + err.Error())
	}

	return nil
}
