package shareholders

import (
	"encoding/json"
	"testing"

	"github.com/xmn-services/rod-network/libs/hash"
)

func TestShareHolder_Success(t *testing.T) {
	hashAdapter := hash.NewAdapter()
	power := uint(554)
	key, _ := hashAdapter.FromBytes([]byte("to build the key hash..."))
	shareHolderIns := CreateShareHolderForTests(power, *key)

	if !shareHolderIns.Key().Compare(*key) {
		t.Errorf("the key hash is invalid")
		return
	}

	if shareHolderIns.Power() != power {
		t.Errorf("the power is invalid, expected: %d, returned: %d", power, shareHolderIns.Power())
		return
	}

	// repository and service:
	repository, service := CreateRepositoryServiceForTests()
	err := service.SaveAll([]ShareHolder{
		shareHolderIns,
	})

	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retShareHolders, err := repository.RetrieveAll([]hash.Hash{
		shareHolderIns.Hash(),
	})

	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if len(retShareHolders) <= 0 {
		t.Errorf("more than 0 shareholders were expected")
		return
	}

	if !shareHolderIns.Hash().Compare(retShareHolders[0].Hash()) {
		t.Errorf("the hash is invalid")
		return
	}

	js, err := json.Marshal(shareHolderIns)
	if err != nil {
		t.Errorf("the ShareHolder instance could not be converted to JSON: %s", err.Error())
		return
	}

	ins := new(shareHolder)
	err = json.Unmarshal(js, ins)
	if err != nil {
		t.Errorf("the JSON instance could not be converted to a ShareHolder instance: %s", err.Error())
		return
	}

	// compare:
	TestCompare(t, ins, shareHolderIns)
}
