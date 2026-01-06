package brancher

import (
	"fmt"

	git "github.com/go-git/go-git/v5"
	plumbing "github.com/go-git/go-git/v5/plumbing"
)

type Manager struct {
	branchMap     map[string]plumbing.Hash
	repository    *git.Repository
	CurrentBranch string
	CurrentCommit string
}

func NewManager() *Manager {

	return &Manager{
		branchMap: make(map[string]plumbing.Hash),
	}

}

func (M *Manager) AcquireLocalRepo() error {

	localRepo, err := git.PlainOpen(".")
	if err != nil {
		return fmt.Errorf("opening repo on current dir: " + err.Error())
	}

	M.repository = localRepo

	return nil
}

func (M *Manager) SyncLocalBranches() error {

	// Make a fresh map on every sync
	M.branchMap = make(map[string]plumbing.Hash)

	head, err := M.repository.Head()
	if err != nil {
		return fmt.Errorf("getting local repo head: " + err.Error())
	}

	M.CurrentBranch = head.Name().Short()
	commit, _ := M.repository.CommitObject(head.Hash())
	M.CurrentCommit = commit.Message

	refrences, err := M.repository.References()
	if err != nil {
		return fmt.Errorf("getting local repo references: " + err.Error())
	}

	err = refrences.ForEach(func(r *plumbing.Reference) error {

		if r.Name().IsBranch() {
			M.branchMap[r.Name().Short()] = r.Hash()
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("iterating local references: " + err.Error())

	}

	return nil

}

func (M *Manager) GetLocalBranches() []string {

	var branches []string
	for branch := range M.branchMap {
		branches = append(branches, branch)
	}
	return branches
}
