<template>
<div id="main">
    <h1 class="title-page">RECOVER WALLET</h1>
    <v-card v-if="!askPassword" class="menu">
        <v-alert v-show="alertShow" :type="alertType">
            {{ alertMessage }}
        </v-alert>
        <h2>Available Wallets</h2>
        <v-select v-model="walletSelected" :items="availableWallets" item-text="name" :color="$selectColor" :item-color="$selectColor" class="wallet-select"></v-select>
        <div class="select-buttons">
            <v-btn color="success" class="select-button" @click="askWalletPassword()" outlined>Open Selected Wallet</v-btn>
            <v-btn color="error" class="select-button" @click="deleteWallet()" outlined>Delete Selected Wallet</v-btn>
        </div>
        <v-divider class="div-space"></v-divider>
        <div class="recover-buttons">
            <v-btn outlined v-for="(button, key) in recoverButtons" :key="key"  :to="button.to" class="recover-button">{{ button.name }}</v-btn>
        </div>
        <v-btn outlined @click="wip()" class="recover-view-wallet">Recover View Only Wallet</v-btn>
    </v-card>
    <v-card v-else class="menu">
        <v-alert v-show="alertShow" :type="alertType">
            {{ alertMessage }}
        </v-alert>
        <h2 style="margin-bottom: 2%;">Password to open {{ walletSelected }}</h2>
        <div id="password">
            <v-text-field type="password" v-model="password" label="Password" :color="$selectColor" filled :disabled="btnDisabled"></v-text-field>
            <div class="select-buttons">
                <v-btn outlined class="recover-button" @click="cancelPassword()" :disabled="btnDisabled">Cancel</v-btn>
                <v-btn outlined class="recover-button" @click="openWallet()" :disabled="btnDisabled">Open Wallet</v-btn>
            </div>
        </div>
    </v-card>
</div>
</template>

<script>
import * as wallet from '../wallet/async-wallet'

export default {
    data() {
        return {
            walletSelected: null,
            availableWallets: [],
            recoverButtons: [ //["Recover using Seed", "Recover using Recovery Key", "Recover using File"],
                {
                    name: "Recover using Seed",
                    to: "/recover-seed"
                },
                {
                    name: "Recover using Recovery Key",
                    to: "/recover-key"
                },
                {
                    name: "Recover using File",
                    to: "/recover-file"
                },
            ],
            alertType: "error",
            alertShow: false,
            alertMessage: "An error as occured!",
            askPassword: false,
            password: "",
            btnDisabled: false
        }
    },
    async mounted() {
        this.availableWallets = await wallet.getWalletsNames()
    },
    methods: {
        askWalletPassword() {
            if (this.walletSelected != null)
            {
                this.askPassword = true
            }
        },
        cancelPassword() {
            this.askPassword = false
        },
        async openWallet() {
            if (this.walletSelected != null)
            {
                this.btnDisabled = true
                if (await wallet.openEncryptedWallet(this.walletSelected, this.password)) {
                    this.alertType = "success"
                    this.alertMessage = "Wallet successfully opened!"
                    this.alertShow = true

                    setTimeout(() => {
                        this.$router.push('/home')
                    } , 1500) //1.5s
                } else {
                    this.alertMessage = "Incorrect password for this wallet !"
                    this.alertShow = true
                    this.btnDisabled = false
                    setTimeout(() => this.alertShow = false, 5000) //after 5s
                }
            }
        },
        async deleteWallet() {
           if (this.walletSelected != null)
           {
                await wallet.removeEncryptedWallet(this.walletSelected)
                this.availableWallets = wallet.getEncryptedWallets()
                this.walletSelected = null
                this.alertType = "success"
                this.alertMessage = "Wallet successfully deleted!"
                this.alertShow = true
                setTimeout(() => this.alertShow = false, 2000) //2s
           } 
        },
        wip() {
            this.alertType = "error"
            this.alertMessage = "Not implemented yet!"
            this.alertShow = true
            setTimeout(() => this.alertShow = false, 2000) //2s   
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

.wallet-select {
    width: 50%;
    margin: auto;
}

.select-buttons {
    display: flex;
    justify-content: space-evenly;
}

.recover-buttons {
    margin-top: 8%;
    display: flex;
    align-items: center;
    justify-content: center;
    flex-direction: column;
}

.recover-button {
    margin-bottom: 2%;
    width: auto;
}

.recover-view-wallet {
    margin-top: 6%;
    width: auto;
}

.div-space {
    margin: 5%;
    margin-left: 15%;
    margin-right: 15%;
}

#password {
    margin-left: 15%;
    margin-right: 15%;
    margin-top: 5%;
    margin-bottom: 5%;
}

@media screen and (max-width: 960px)
{
    .menu {
        margin-top: 3%;
        margin-left: 2%;
        margin-right: 2%;
        margin-bottom: 3%;

        padding: 2%;
    }
    .select-buttons {
        display: flex;
        flex-direction: column;
        justify-content: space-around;
    }
    .recover-button {
        margin-bottom: 2%;
        width: auto;
    }
    .recover-view-wallet {
        margin-top: 6%;
        margin-bottom: 5%;
        width: auto;
    }
    .select-button {
        margin: auto;
        margin-bottom: 5%;
    }
    .title-page {
        margin-top: 5%;
    }
}
</style>