package mocks

func New() MockedTns {
	return &mockTns{
		mapDef: make(map[string]interface{}, 0),
	}
}

// Wasm Path
// wasm/project/QmZVGRba3AKTfYav4KbN45QFBYHETWCxhEV7pAySCHrVYS/modules/libraries/fakeModuleName
// RETURNS::: CURRENT PATH LIST

// Current Path
// branches/master/commit/qwertyuiopasdfghjklzxcvbnm/projects/QmZVGRba3AKTfYav4KbN45QFBYHETWCxhEV7pAySCHrVYS/applications/QmeksLEwbGdJxZ6w8k2Xjtu7NqKxGUb6NxdQ8seKvBVdvV/libraries/QmdXdwXK376cSVSkdZ4UbhS56ZuhE58qCGHX7tzbyaeKV1
// RETURNS:::: Nil

// Asset Hash
// assets/QmRNXxfoVLXQTLRKqh7NM76uiMdT8EdE7VqQH5ZNDeLatQ
