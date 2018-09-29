pragma solidity ^0.4.25;

contract ScoreToResultCode {
    modifier validateScore (uint8 _score) {
        require(_score >= 0 && _score <= 100);
        _;
    }
    
    function run (uint8 _score) public pure validateScore(_score) returns (string) {
        uint8[8] memory scores = [90, 80, 70, 65, 60, 50, 30, 0];
        string[8] memory codes = ['AA', 'BA', 'BB', 'CB', 'CC', 'DD', 'FD', 'FF'];
        
        for (uint i=0; i < scores.length; i++) {
            if (_score >= scores[i]) return codes[i];
        }
    }
}
