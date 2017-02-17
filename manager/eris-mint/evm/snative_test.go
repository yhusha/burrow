package vm

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
	"fmt"
)

/* Compiling the Permissions solidity contract at
(generated by Solidity() function)
https://ethereum.github.io/browser-solidity yields:

3fbf7da5 add_role(address,bytes32)
744f5998 has_base(address,uint64)
e8145855 has_role(address,bytes32)
28fd0194 rm_role(address,bytes32)
3f0ebb30 set_base(address,uint64,uint64)
67dc6f70 set_global(address,uint64,uint64)
73448c99 unset_base(address,uint64)
*/

func TestPermissionsContract(t *testing.T) {
	registerNativeContracts()
	contract := SNativeContracts()["permissions_contract"]

	assertContractFunction(t, contract, "3fbf7da5",
		"add_role(address,bytes32)")

	assertContractFunction(t, contract, "744f5998",
		"has_base(address,uint64)")

	assertContractFunction(t, contract, "e8145855",
		"has_role(address,bytes32)")

	assertContractFunction(t, contract, "28fd0194",
		"rm_role(address,bytes32)")

	assertContractFunction(t, contract, "3f0ebb30",
		"set_base(address,uint64,uint64)")

	assertContractFunction(t, contract, "67dc6f70",
		"set_global(address,uint64,uint64)")

	assertContractFunction(t, contract, "73448c99",
		"unset_base(address,uint64)")
}

func TestSNativeFuncTemplate(t *testing.T) {
	contract := SNativeContracts()["permissions_contract"]
	function, err := contract.FunctionByName("rm_role")
	if err != nil {
		t.Fatal("Couldn't get function")
	}
	solidity, err := function.Solidity()
	assert.NoError(t, err)
	fmt.Println(solidity)
}

// This test checks that we can generate the SNative contract interface and
// prints it to stdout
func TestSNativeContractTemplate(t *testing.T) {
	contract := SNativeContracts()["permissions_contract"]
	solidity, err := contract.Solidity()
	assert.NoError(t, err)
	fmt.Println(solidity)
}

// Helpers

func assertContractFunction(t *testing.T, contract SNativeContractDescription,
	funcIDHex string, expectedSignature string) {
	function, err := contract.FunctionByID(fourBytesFromHex(t, funcIDHex))
	assert.NoError(t, err,
		"Error retrieving SNativeFunctionDescription with ID %s", funcIDHex)
	if err == nil {
		assert.Equal(t, expectedSignature, function.Signature())
	}
}

func fourBytesFromHex(t *testing.T, hexString string) [4]byte {
	bs, err := hex.DecodeString(hexString)
	assert.NoError(t, err, "Could not decode hex string '%s'", hexString)
	if len(bs) != 4 {
		t.Fatalf("FuncID must be 4 bytes but '%s' is %v bytes", hexString,
			len(bs))
	}
	return firstFourBytes(bs)
}

