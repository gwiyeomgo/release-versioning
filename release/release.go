package release

type Release struct {
	ID string `json:"id"`
}

func Find(hash string) (*Release, error) {
	//blockBytes := db.Block(hash)
	/*	blockBytes := dbStorage.FindBlock(hash)
		if blockBytes == nil {
			return nil, ErrNotFound
		}
		block := &Block{}
		block.restore(blockBytes)
		return block, nil*/
	return nil, nil
}

func createBlock(preHash string, height int, diff int) *Release {
	/*	block := Release{
			Hash:    "",
		}

		persistBlock(&block)
		return &block */
	return nil
}
