pragma solidity ^0.6.0;
contract MiniStore {
    uint256[] internal arrayData;
    uint256 internal numberValue;
    constructor() public {
    }
    function getNumberValue() public view returns(uint256){
        return numberValue;
    }
    function setNumberValue(uint256 _value) public {
        numberValue = _value;
    }
    function getArrayValueAt(uint256 i) public view returns(uint256){
        require(i < arrayData.length,"Index out of range");
        return arrayData[i];
    }
    function getArrayDataLength() public view returns(uint256){
        return arrayData.length;
    }
    function getArrayLatestValue() public view returns(uint256){
        if(arrayData.length > 0){
            return arrayData[arrayData.length-1];
        }
        else {
            return 0;
        }
    }
    function setArrayValue(uint256 _i, uint256 _value) public returns (uint256){
        if( _i < arrayData.length){
            arrayData[_i] = _value;
            return _i;
        }else {
            arrayData.push(_value);
            return arrayData.length-1;
        }
    }
    function deleteLastArrayValue() public returns (uint256) {
        require(arrayData.length > 0,"array is empty");
        uint256 value = arrayData[arrayData.length-1];
        arrayData.pop();
        return value;
    }
    function addValue(uint256 _value) public returns (uint256) {
        arrayData.push(_value);
        return arrayData.length-1;
    }
    function insertArray(uint256 _startIndex, uint256[] memory _array) public {
        for(uint256 i = 0; i< _array.length; i++ ){
            if(i + _startIndex < arrayData.length){
                arrayData[i + _startIndex] = _array[i];
            } else{
                arrayData.push(_array[i]);
            }
        }
    }
}