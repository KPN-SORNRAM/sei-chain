require('dotenv').config({path:__dirname+'/.env'})
require("@nomicfoundation/hardhat-toolbox");

/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
  solidity: {
    version: "0.8.20",
    settings: {
      optimizer: {
        enabled: true,
        runs: 1000,
      },
    },
  },
  paths: {
    sources: "./src", // contracts are in ./src
  },
  networks: {
    goerli: {
      url: "https://eth-goerli.g.alchemy.com/v2/NHwLuOObixEHj3aKD4LzN5y7l21bopga", // Replace with your JSON-RPC URL
      address: ["0xF87A299e6bC7bEba58dbBe5a5Aa21d49bCD16D52"],
      accounts: ["0x57acb95d82739866a5c29e40b0aa2590742ae50425b7dd5b5d279a986370189e"], // Replace with your private key
    },
    // sei: {
    //   url: "https://evm-warroom-test.seinetwork.io:18545", // Replace with your JSON-RPC URL
    //   address: ["0x07dc55085b721947d5c1645a07929eac9f1cc750"],
    //   accounts: [process.env.TEST_PRIVATE_KEY], // Replace with your private key
    // },
    seilocal: {
      url: "http://127.0.0.1:8545",
    }
  },
};
