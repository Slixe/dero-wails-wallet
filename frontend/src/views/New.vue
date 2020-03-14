<template>
<div id="new">
    <h1 class="title-page">CREATE NEW WALLET</h1>
    <v-card v-if="!showSeed" class="menu">
        <v-alert v-show="alertShow" :type="alertType">
            {{ alertMessage }}
        </v-alert>
        <span>Please fill in all the blanks</span>
        <v-divider class="new-div"></v-divider>
        <v-text-field autocomplete="new-password" v-model="walletName" label="Wallet name for quick reference" :color="$selectColor" filled :disabled="btnDisabled"></v-text-field>
        <v-text-field autocomplete="new-password" type="password" v-model="password" label="Password" :color="$selectColor" filled :disabled="btnDisabled"></v-text-field>
        <v-text-field autocomplete="new-password" type="password" label="Confirm password" :color="$selectColor" filled :disabled="btnDisabled"></v-text-field>
        <v-btn outlined class="create-wallet" @click="newWallet()" :disabled="btnDisabled">Create Wallet</v-btn>
    </v-card>
    <v-card v-else class="menu">
        <h1>WALLET SEED</h1>
        <v-divider class="new-div"></v-divider>
        <span><font color="#D32F2F">Your wallet can be recovered using the below seed.</font></span>
        <span><font color="#D32F2F">You must save the SEED in a safe secure location.</font></span>
        <span><font color="#D32F2F">Sharing your SEED is equal to sharing your wallet. If your SEED is lost, consider your wallet as lost.</font></span>
        <v-divider class="new-div"></v-divider>
        <v-select label="Language" :color="$selectColor" :item-color="$selectColor" :items="seedLanguages" v-model="seedSelected" @change="seed()" outlined></v-select>
        <v-textarea label="SEED" auto-grow :value="walletSeed" disabled filled></v-textarea>
        <v-btn outlined class="create-wallet" to="/home">I saved it!</v-btn>
    </v-card>
</div>
</template>

<script>
import * as wallet from '../wallet/async-wallet'

export default {
    data() {
        return {
            walletName: "",
            password: "",
            alertType: "error",
            alertShow: false,
            alertMessage: "An error as occured!",
            btnDisabled: false,
            walletSeed: "",
            showSeed: false,
            seedSelected: "English",
            seedLanguages: ["English", "Français", "Deutsch", "Italiano", "Español", "Português", "Nederlands", "Esperanto", "русский язык", "日本語", "简体中文 (中国)"]
        }
    },
    methods: {
        async newWallet() {
            this.btnDisabled = true

            wallet.createWallet(this.walletName, this.password).then(async result => {
                if (result === "success") {
                    this.alertType = "success"
                    this.alertMessage = "Wallet successfully created!"
                    this.alertShow = true
                    setTimeout(() => this.alertShow = false, 5000) //after 5s

                    this.walletSeed =  await wallet.getSeedInLanguage(this.seedSelected)
                    setTimeout(() => {
                        this.showSeed = true
                    } , 1500) //1.5s
                } else {
                    this.alertType = "error"
                    this.alertMessage = result
                    this.alertShow = true
                    this.btnDisabled = false
                    setTimeout(() => this.alertShow = false, 5000) //after 5s
                }
            })
        },
        async seed() {
            this.walletSeed = await wallet.getSeedInLanguage(this.seedSelected)
        }
    }
}
</script>

<style scoped>
.menu {
    margin-top: 3%;
    margin-left: 20%;
    margin-right: 20%;
    margin-bottom: 3%;

    padding-top: 2%;
    padding-bottom: 2%;
    padding-left: 10%;
    padding-right: 10%;

    display: flex;
    flex-direction: column;
}

.new-label {
    text-align: center;
}

.create-wallet {
    width: auto;
    margin: auto;
    margin-top: 6%;
}

.new-div {
    margin: 5%;
}

@media screen and (max-width: 960px) {
    .menu {
        margin-top: 3%;
        margin-left: 2%;
        margin-right: 2%;
        margin-bottom: 3%;

        padding-top: 5%;
        padding-bottom: 5%;
    }

    .create-wallet {
        width: auto;
    }

    .title-page {
        margin-top: 5%;
    }
}
</style>