// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.9.0;

contract testError {
  constructor()  {
  }

  function mustRevert() public {
    revert("Error");
  }

  function mustRevertByCall() public pure{
    revert("Error when read contract");
  }

  function Error(string memory msg) public pure{ 
  }
}
