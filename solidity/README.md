# ERC-20 swapping contract

The task is to create a simple Solidity contract for exchanging Ether to an arbitrary ERC-20.

## Requirements

1. Implement the following interface as a Solidity Contract

   ```solidity
   interface ERC20Swapper {
       /// @dev swaps the `msg.value` Ether to at least `minAmount` of tokens in `address`, or reverts
       /// @param token The address of ERC-20 token to swap
       /// @param minAmount The minimum amount of tokens transferred to msg.sender
       /// @return The actual amount of transferred tokens
       function swapEtherToToken(address token, uint minAmount) public payable returns (uint);
   }
   ```

2. Deploy the contract to a public Ethereum testnet (e.g. Sepolia)
3. Send the address of deployed contract and the source code to us

### Non-requirements

- Feel free to implement the contract by integrating to whatever DEX you feel comfortable - the exchange implementation is not required.

## Evaluation

Following properties of the contract implementation will be evaluated in this exercise:

- **Safety and trust minimization**. Are user's assets kept safe during the exchange transaction? Is the exchange rate fair and correct? Does the contract have an owner?
- **Performance**. How much gas will the `swapEtherToToken` execution and the deployment take?
- **Upgradeability**. How can the contract be updated if e.g. the DEX it uses has a critical vulnerability and/or the liquidity gets drained?
- **Usability and interoperability**. Is the contract usable for EOAs? Are other contracts able to interoperate with it?
- **Readability and code quality**. Are the code and design understandable and error-tolerant? Is the contract easily testable?

# Solution

## Upgradeability

Using OpenZeppelin upgrades, Transparent proxy solution implemented with [Foundry plug-in](https://docs.openzeppelin.com/upgrades-plugins/1.x/foundry-upgrades)

The swapper contract is deployed at sepolia address 0xD0B98e56F19f9b5c2Ef0de72c0bc799a99555261

The transparent proxy contract is deployed at sepolia address 0xdaE97900D4B184c5D2012dcdB658c008966466DD

For future upgrades, the swapper contract needs to follow the guide by Openzeppelin about upgradeability, basically to keep the state backwards compatible.

## Safety 

I have presumed that the swapper contract owns tokens at the ERC20 contracts and when it receives ETHER, it transfers its own tokens to the sender. In this case, the transaction is safe, as in case the swapper contract doesn't have enough tokens, it would be reverted and sender would not lose his ether. 
A more complex approach would be to consider that the tokens belong to someone else. Then contract would need to facilitate the transaction between two parties, using the approve() and transferFrom() functions from ERC20 (similar to how Uniswap works). 

_Is the exchange rate fair?_ 

The exchange rate is provided by Uniswap v3 pools. Each pair token/weth may have up to three pools with different fees (they call them tiers). The swapper contract makes use of Uniswapv3 Factory contract to search for a token/weth pool in any of the tiers. From that pool, the exchange rate is obtained. This is not the most reliable way to get the exchange rate, although it is reasonably good. The price in principle is fair, if the token/weth pool has enough liquidity. Otherwise it might be subject to sudden variations. 

The best way would be to use a decentralized oracle (for instance, chainlink). This solution provides a more reliable price, not subject to sudden and intentionally provoked fluctuations. The downside is that it does not offers the functionality of just using a pair of contract addresses to search for a price feed (price feeds addresses for each supported pair of tokens are published on the web, that is, off-chain). Besides, using chainlink price feeds costs extra fees.

## Performance

The swapper contract has taken: 

Total Paid: 0.002522603083589364 ETH (783382 gas * avg 3.220144302 gwei)


## Usuability and Interoperativity

Yes, Externally Owned Accounts can interact/transact with the swapper contract, sending ether and receiving tokens. 

Also, other contracts can interact in the same way. 

## Readability and code quality

Comments have been added to help understanding the logic behind it. 

Most potential errored situations have been taken into accout (addess not being zero, only-owner transactions, etc.). 
Test cases are available for the self contained logic. However, the functions that rely on calls to external contracts (either Uniswap v3 or the ERC20 token) need a more complex setup. A possible solution (just a hint implemented) is to use mocks based on virtual functions for the external contracts. This is not an easy/straightforward solution, it can get messy if the number of contracts/functions to mock is reasonably large. 
