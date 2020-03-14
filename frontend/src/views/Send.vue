<template>
<div id="send">
    <h1>SEND DERO</h1>
    <v-card class="menu">
        <span>Please note that sending to another address is irreversible.</span>
        <v-divider class="send-div"></v-divider>
        <v-text-field label="DERO Address" :color="$selectColor" filled></v-text-field>
        <v-text-field label="Amount" :color="$selectColor" filled></v-text-field>
        <label class="new-label">PaymentID:</label>
        <div class="payment-id">
            <v-radio-group v-model="paymentId" row>
                <v-radio color="green" label="Yes" value="yes"></v-radio>
                <v-radio color="red" label="No" value="no"></v-radio>
            </v-radio-group>
        </div>
        <div v-show='paymentId === "yes"'>
            <v-text-field label="Payment ID" :color="$selectColor" filled></v-text-field>
        </div>
        <v-btn outlined class="send-transaction">Send Transaction</v-btn>
    </v-card>
</div>
</template>

<script>
export default {
    data() {
        return {
            paymentId: "no"
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

.send-transaction {
    margin: auto;
    margin-top: 5%;
    width: auto;
}

.payment-id {
    margin-top: 2%;
    
    display: flex;
    flex-wrap: wrap;
    flex-direction: row;
    justify-content: space-evenly;
    margin: auto;
}

.send-div {
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
}
</style>