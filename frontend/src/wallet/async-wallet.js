import { EventBus } from '../event-bus'

/* eslint-disable */

let hasWallet = false
let walletName = "No Wallet"
let walletAddress
let syncing

//from original dero web wallet (wallet.js)
let getRandomBytes = ((typeof self !== 'undefined' && (self.crypto || self.msCrypto)) ? function() { // Browsers
    let crypto = (self.crypto || self.msCrypto), QUOTA = 65536
    return function(n) {
        let a = new Uint8Array(n)
        for (var i = 0; i < n; i += QUOTA)
        {
            crypto.getRandomValues(a.subarray(i, i + Math.min(n - i, QUOTA)))
        }
        return a;
    }
} : function() { // Node
    return require("crypto").randomBytes;
}
)();

function toHexString(byteArray) {
    return byteArray.reduce((output, elem) => (output + ('0' + elem.toString(16)).slice(-2)), '')
}

function backend()
{
    return window.backend
}
//end

export async function createWallet(walletName, password)
{
    let result = await backend().createNewWallet(walletName, password)

    if (result === "success")
    {
        setWalletName(walletName)
        await onlineMode(true)
    }

    return result
}

export async function openEncryptedWallet(walletName, password)
{
    let result = await backend().openEncryptedWallet(walletName, password)
    result = result === "success"

    if (result)
    {
        setWalletName(walletName)
        await onlineMode(true)
    }

    return result
}

export async function recoverWalletSeed(walletName, password, seed)
{

    let result = await backend().createEncryptedWalletFromRecoveryWords(walletName, password, seed)
    
    console.log(result)
    if (result === "success") {
        setWalletName(walletName)
        await onlineMode(true)
        addEncryptedWallet(walletName)
    }

    return result
}

export async function recoverViewWallet(walletName, password, viewKey)
{
    let result = await backend().createEncryptedWalletViewOnly(walletName, password, viewKey)

    console.log(result)
    if (result === "success") {
        setWalletName(walletName)
        await onlineMode(true)
        addEncryptedWallet(walletName)
    }

    return result
}

export async function getInfos()
{
    let result = await backend().getInfos()
    result = JSON.parse(result)

    syncing = result.DaemonTopoHeight - result.WalletTopoHeight > 100 //If more than 100, so it's syncing

    return result
}

export async function generateIntegratedAddress()
{
    return JSON.parse(await backend().generateIntegratedAddress())
}

export async function getSeedInLanguage(language)
{
    return await backend().getSeedInLanguage(language)
}

export async function getTxHistory()
{
    return null //await run("getTxHistory")
}

export async function transfer(addresses, amounts, paymentId)
{
    return null//await run("transfer", addresses, amounts, paymentId)
}

export async function onlineMode(online)
{
    syncing = online

    return await backend().setOnlineMode(online)
}

export async function closeWallet()
{
    await onlineMode(false)
    let result = await backend().closeWallet()
    hasWallet = false
    walletName = "No Wallet"

    return result
}
 
export async function getWalletsNames()
{
    return await backend().getWallets() 
}

export async function removeEncryptedWallet(walletName)
{
    return await backend().removeWalletFile(walletName)
}

export function setWalletName(name)
{
    walletName = name
    hasWallet = name != null

    if (name != null) {
        getInfos().then(result => {
            walletAddress = result.WalletAddress
        })
    } else {
        walletAddress = null
    }
    
    EventBus.$emit('isWalletOpen', hasWallet)
}

export function getWalletAddress()
{
    return walletAddress
}

export function getWalletName()
{
    return walletName
}

export function hasWalletOpen()
{
    return hasWallet
}

export function isSyncing()
{
    return syncing
}