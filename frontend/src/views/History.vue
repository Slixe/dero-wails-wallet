<template>
<div id="history">
    <h1>HISTORY</h1>
    <v-card class="menu">
        <h1 v-if="!ready">{{ message }}</h1>
        <v-simple-table v-else dark id="table">
            <template v-slot:default>
                <thead>
                    <tr>
                        <th class="text-center">Height</th>
                        <th class="text-center">Topo Height</th>
                        <th class="text-center">Timestamp</th>
                        <th class="text-center">Tx</th>
                        <th class="text-center">Amount</th>
                    </tr>
                    </thead>
                    <tbody>
                        <tr>
                            <td>height</td>
                            <td>topo height</td>
                            <td>block date</td>
                            <td>tx id</td>
                            <td>amount</td>
                        </tr>
                    </tbody>
                </template>
            </v-simple-table>
    </v-card>
</div>
</template>

<script>
import * as wallet from '../wallet/async-wallet'

export default {
    data() {
        return {
            message: "Loading...",
            txHistory: [],
            ready: false
        }
    },
    mounted() {
        wallet.getTxHistory().then(result => {
            /* eslint-disable */
            console.log(result)
        }).catch(err => {
            this.message = err
        })
    }
}
</script>

<style scoped>
.menu {
    margin-top: 3%;
    margin-left: 2%;
    margin-right: 2%;
    margin-bottom: 3%;

    padding-top: 2%;
    padding-bottom: 2%;
    padding-left: 10%;
    padding-right: 10%;

    display: flex;
    flex-direction: column;
}
</style>