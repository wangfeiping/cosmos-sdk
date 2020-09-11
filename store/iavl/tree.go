package iavl

var (
	_ Tree = (*immutableTree)(nil)
)

type (
	// Tree defines an interface that both mutable and immutable IAVL trees
	// must implement. For mutable IAVL trees, the interface is directly
	// implemented by an iavl.MutableTree. For an immutable IAVL tree, a wrapper
	// must be made.
	Tree interface {
		Has(key []byte) bool
		Get(key []byte) (index int64, value []byte)
		Set(key, value []byte) bool
		Remove(key []byte) ([]byte, bool)
		SaveVersion() ([]byte, int64, error)
		DeleteVersion(version int64) error
		DeleteVersions(versions ...int64) error
		Version() int64
		Hash() []byte
		VersionExists(version int64) bool
		GetVersioned(key []byte, version int64) (int64, []byte)
		SetInitialVersion(version uint64)
	}

	// immutableTree is a simple wrapper around a reference to an iavl.ImmutableTree
	// that implements the Tree interface. It should only be used for querying
	// and iteration, specifically at previous heights.
	immutableTree struct {
	}
)

func (it *immutableTree) Set(_, _ []byte) bool {
	panic("cannot call 'Set' on an immutable IAVL tree")
}

func (it *immutableTree) Remove(_ []byte) ([]byte, bool) {
	panic("cannot call 'Remove' on an immutable IAVL tree")
}

func (it *immutableTree) SaveVersion() ([]byte, int64, error) {
	panic("cannot call 'SaveVersion' on an immutable IAVL tree")
}

func (it *immutableTree) DeleteVersion(_ int64) error {
	panic("cannot call 'DeleteVersion' on an immutable IAVL tree")
}

func (it *immutableTree) DeleteVersions(_ ...int64) error {
	panic("cannot call 'DeleteVersions' on an immutable IAVL tree")
}

func (it *immutableTree) SetInitialVersion(_ uint64) {
	panic("cannot call 'SetInitialVersion' on an immutable IAVL tree")
}

func (it *immutableTree) VersionExists(version int64) bool {
	return it.Version() == version
}

func (it *immutableTree) GetVersioned(key []byte, version int64) (int64, []byte) {
	if it.Version() != version {
		return -1, nil
	}

	return it.Get(key)
}

func (it *immutableTree) Has(key []byte) bool {
	return false
}

func (it *immutableTree) Get(key []byte) (index int64, value []byte) {
	return -1, nil
}

func (it *immutableTree) Version() int64 {
	return -1
}

func (it *immutableTree) Hash() []byte {
	return nil
}
