// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface IERC20 {

    function totalSupply() external view returns (uint256);
    function balanceOf(address account) external view returns (uint256);
    function allowance(address owner, address spender) external view returns (uint256);

    function transfer(address recipient, uint256 amount) external returns (bool);
    function approve(address spender, uint256 amount) external returns (bool);
    function transferFrom(address sender, address recipient, uint256 amount) external returns (bool);


    event Transfer(address indexed from, address indexed to, uint256 value);
    event Approval(address indexed owner, address indexed spender, uint256 value);
}

contract Gacha is IERC20 {
  string public constant name = "GACHA coin";
  string public constant symbol = "GAC";
  uint256 public totalSupply_ = 1000*10**18;
  mapping (address => uint256) public balances;
  mapping (address => bool) public account;
  mapping(address => mapping (address => uint256)) allowed;

  event Bought(uint256 amount);

  function totalSupply() public override view returns (uint256) {
    return totalSupply_;
  }

  function balanceOf(address tokenOwner) public override view returns (uint256) {
    return balances[tokenOwner];
  }

  function createAccount() public returns (bool) {
    return account[msg.sender] = true;
  }

  function transfer(address receiver, uint256 numTokens) public override returns (bool) {
    require(numTokens <= balances[msg.sender], "not enough tokens to send");
    balances[msg.sender] = balances[msg.sender]-numTokens;
    balances[receiver] = balances[receiver]+numTokens;
    emit Transfer(msg.sender, receiver, numTokens);
    return true;
  }

  function approve(address delegate, uint256 numTokens) public override returns (bool) {
      allowed[msg.sender][delegate] = numTokens;
      emit Approval(msg.sender, delegate, numTokens);
      return true;
    }

  function allowance(address owner, address delegate) public override view returns (uint) {
    return allowed[owner][delegate];
  }
  function transferFrom(address owner, address buyer, uint256 numTokens) public override returns (bool) {
      require(numTokens <= balances[owner]);
      require(numTokens <= allowed[owner][msg.sender]);
      balances[owner] = balances[owner]-numTokens;
      allowed[owner][msg.sender] = allowed[owner][msg.sender]-numTokens;
      balances[buyer] = balances[buyer]+numTokens;
      emit Transfer(owner, buyer, numTokens);
      return true;
  }

  function buy() payable public {
    uint256 amountTobuy = msg.value;
    uint256 dexBalance = balanceOf(address(this));
    require(amountTobuy > 0, "You need to send some ether");
    require(amountTobuy <= dexBalance, "Not enough tokens in the reserve");
    transfer(msg.sender, amountTobuy);
    emit Bought(amountTobuy);
  }
}
