<template>
<div id="receive">
    <h1 class="title-page">RECEIVE DERO</h1>
    <v-card class="menu">
        <v-alert v-show="alertShow" :type="alertType">
            {{ alertMessage }}
        </v-alert>
        <h2>Wallet Address</h2>
        <v-divider class="new-div"></v-divider>
        <v-text-field v-model="walletAddress" label="Wallet Address" :color="$selectColor" filled disabled></v-text-field>
        <v-text-field v-show="paymentID != ''" v-model="integratedAddress" label="Integrated Address" :color="$selectColor" filled disabled></v-text-field>
        <v-text-field v-show="paymentID != ''" v-model="paymentID" label="Payment ID (8 bytes)" :color="$selectColor" filled disabled></v-text-field>
        <v-btn outlined class="generate-ia" @click.prevent="generateIA()">Generate Integrated Address</v-btn>
    </v-card>
</div>
</template>

<script>
import * as wallet from '../wallet/async-wallet'

export default {
    data() {
        return {
            walletAddress: "",
            integratedAddress: "",
            paymentID: "",
            alertType: "error",
            alertShow: false,
            alertMessage: "An error as occured!",
            timeOut: null
        }
    },
    async mounted() {
        this.walletAddress = wallet.getWalletAddress()
    },
    methods: {
        async generateIA() {
            this.alertMessage = "New integrated address generated!"
            this.alertType = "success"
            this.alertShow = true

            if (this.timeOut != null) {
                clearTimeout(this.timeOut)
            }

            this.timeOut = setTimeout(() => this.alertShow = false, 2500)

            let ia = await wallet.generateIntegratedAddress()

            this.integratedAddress = ia.Address
            this.paymentID = ia.PaymentId
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

.generate-ia {
    width: auto;
    margin: auto;
    margin-top: 6%;
}

.new-div {
    margin: 5%;
    margin-left: 2%;
    margin-right: 2%;
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