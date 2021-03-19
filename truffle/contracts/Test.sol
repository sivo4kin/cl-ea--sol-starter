pragma solidity >=0.4.21 < 0.7.0;

contract Test {
    uint256 public test;
//    function Test(){
//
//    }

    function setTest(uint256 val) public  {
        test = val;
    }

    function getTest() public view returns (uint256)  {
        return test;
    }

    function lowLevelGet(bytes memory b) public view returns (bytes32) {
        (bool success, bytes memory data) = address(this).staticcall(b);
        if (!success || data.length == 0) {
            return '';
        }
        return abi.decode(data, (bytes32));
    }

}
