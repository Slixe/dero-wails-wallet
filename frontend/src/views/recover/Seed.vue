<template>
<div id="recover-seed">
    <h1 class="title-page">RECOVER WALLET USING SEED</h1>
    <v-card class="menu">
        <v-alert v-show="alertShow" :type="alertType">
            {{ alertMessage }}
        </v-alert>
        <div id="seed">
            <v-text-field autocomplete="new-password" v-model="walletName" label="Wallet Name" :color="$selectColor" filled :disabled="btnDisabled"></v-text-field>
            <v-textarea autocomplete="new-password" type="password" v-model="seed" label="Seed" :color="$selectColor" filled :disabled="btnDisabled"></v-textarea>
            <v-text-field autocomplete="new-password" type="password" v-model="password" label="Password" :color="$selectColor" filled :disabled="btnDisabled"></v-text-field>
            <v-text-field autocomplete="new-password" type="password" v-model="confirmPassword" label="Confirm Password" :color="$selectColor" filled :disabled="btnDisabled"></v-text-field>
            <div class="select-buttons">
                <v-btn outlined class="recover-button" to="/" :disabled="btnDisabled">Cancel</v-btn>
                <v-btn outlined class="recover-button" @click="recoverWallet()" :disabled="btnDisabled">Recover Wallet</v-btn>
            </div>
        </div>
    </v-card>
</div>
</template>

<script>
import * as wallet from '../../wallet/async-wallet'

export default {
    data() {
        return {
            alertType: "error",
            alertShow: false,
            alertMessage: "An error as occured!",
            seed: "",
            password: "",
            confirmPassword: "",
            walletName: "",
            btnDisabled: false
        }
    },
    methods: {
        async recoverWallet() {
            this.btnDisabled = true

            if (wallet.encryptedWalletExist(this.walletName)) {
                this.alertMessage = "This wallet name is already used!"
                this.alertShow = true
                this.btnDisabled = false
                setTimeout(() => this.alertShow = false, 5000) //after 5s
                
                return
            }

            let result = await wallet.recoverWalletSeed(this.walletName, this.password, this.seed)

            if (result === "success") {
                this.alertType = "success"
                this.alertMessage = "Wallet successfully recovered from seed!"
                this.alertShow = true

                setTimeout(() => this.$router.push('/home'), 1500) //1.5s
            } else {
                this.alertMessage = result
                this.alertShow = true
                this.btnDisabled = false
                setTimeout(() => this.alertShow = false, 5000) //after 5s
            }
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

    padding: 2%;
}

.recover-button {
    margin-bottom: 2%;
    width: auto;
}

#seed {
    margin-left: 15%;
    margin-right: 15%;
    margin-top: 5%;
    margin-bottom: 5%;
}

.select-buttons {
    display: flex;
    justify-content: space-evenly;
}
</style>