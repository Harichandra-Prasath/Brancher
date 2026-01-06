package brancher

import (
	"errors"
	"fmt"
	"os"
	"os/exec"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
)

var NO_BRANCH_ERR = fmt.Errorf("branch doesnt exist")
var ALREDY_UPTO_DATE = fmt.Errorf("already upto date")

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

	cmd := exec.Command("git", "checkout", name)
	cmd.Dir = "."

	err = cmd.Run()
	if err != nil {
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

func (M *Manager) BranchPull(name string) error {

	if _, ok := M.branchMap[name]; !ok {
		return NO_BRANCH_ERR
	}

	if name != M.CurrentBranch {
		return fmt.Errorf("Requested Branch is not the Active Branch")
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
		return fmt.Errorf("Worktree is not clean. Manage the changes before pulling the source to not lose the local changes.")
	}

	key, err := ssh.NewPublicKeysFromFile("git", os.Getenv("HOME")+"/.ssh/"+os.Getenv("PV_KEY_FILE"), "")
	if err != nil {
		return fmt.Errorf("creating public key: " + err.Error())
	}

	err = workTree.Pull(&git.PullOptions{SingleBranch: true, Auth: key})
	if err != nil {
		if errors.Is(err, git.NoErrAlreadyUpToDate) {
			return ALREDY_UPTO_DATE
		}
		return fmt.Errorf("pulling remote: " + err.Error())
	}

	return nil
}
