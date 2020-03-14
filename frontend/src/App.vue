<template>
  <div id="app">
    <v-app>
      <header>
        <v-toolbar id="navbar">
          <v-btn v-if="mobile" text icon @click="menu = !menu">
            <v-icon>menu</v-icon>
          </v-btn>
          <div v-else class="logo">
            <v-img src="https://dero.io/img/logo.png"/>
          </div>
          <h1 class="navbar-title">{{ $name }}</h1>
        </v-toolbar>
      </header>
      <v-content>
        <div class="wallet-content">
          <div v-show="menu" id="side-menu">
            <ul class="buttons" v-for="(btn, i) in (walletOpened ? buttonsWalletOpen : buttons)" :key="i">
              <li>
                <v-btn class="button" :to="btn.to" @click="changeState(btn)" text :color="btn.color != null ? btn.color : ''">{{ btn.name }}</v-btn>
              </li>
            <v-divider></v-divider>
            </ul>
          </div>
          <div id="router-content" v-show="mobile ? !menu : true">
            <transition name="fade">
              <router-view></router-view>
            </transition>
          </div>
        </div>
      </v-content>
      <v-footer elevation="10" padless>
        <v-card flat tile width="100%" class="text-center">
          <v-card-text v-if="$donations">
            dERokevAZEZVJ2N7o39VH81BXBqX9ojtncnPTDMyiVbmYiTXQY93AUCLcor9xsWCKWhYy25ja89ikZWXWab9kXRB7LYfUmbQyS
          </v-card-text>
          <v-divider v-if="$donations"></v-divider>
          <v-card-text class="white--text">
            {{ new Date().getFullYear() }} — <strong>{{ $name }}</strong>
          </v-card-text>
        </v-card>
      </v-footer>
    </v-app>
  </div>
</template>

<script>
import * as wallet from './wallet/async-wallet'
import { EventBus } from './event-bus';

export default {
  name: 'app',
  components: {
  },
  data() {
    return {
      menu: true,
      mobile: false,
      walletOpened: false,
      buttons: [
        {
          name: "Recover Wallet",
          to: "/"
        },
        {
          name: "Create New Wallet",
          to: "/new"
        },
        {
          name: "About",
          to: "/about"
        }
      ],
      buttonsWalletOpen: [
        {
          name: "Home",
          to: "/home"
        },
        {
          name: "Receive",
          to: "/receive"
        },
        {
          name: "Send",
          to: "/send"
        },
        {
          name: "History",
          to: "/history"
        },
        {
          name: "Logout",
          color: "red",
          logout: true,
        },
        {
          name: "About",
          to: "/about"
        }
      ]
    }
  },
  mounted() {
    this.mobile = (typeof window.orientation !== "undefined") || (navigator.userAgent.indexOf('IEMobile') !== -1)

    EventBus.$on('isWalletOpen', walletOpen => {
      this.walletOpened = walletOpen
    })
  },
  methods: {
    changeState(btn)
    {
      if (this.mobile) {
        this.menu = !this.menu
      }
      
      if (btn.logout) {
        this.walletOpened = false
        this.$router.push('/')
        wallet.closeWallet()
      }
    }
  }
}
</script>

<style>
#app {
  text-align: center;
  color: white; /*#2c3e50;*/
  background-color: #4d4d4d;
}

#navbar {
  width: 100%;
  position: relative;
}

#side-menu {
  overflow-y: auto;
  flex: 15%;
  height: 100%;
  width: 255px;
  background-color: #525252;
  position: absolute;
  z-index: 1;
}

#router-content {
  flex: 85%;
  padding-left: 1%;
  padding-right: 1%;
  padding-top: 2%;
  padding-bottom: 2%;
  margin-left: 255px;
}

.wallet-content {
  display: flex;
  flex-wrap: wrap;
  height: 100%;
}

.logo {
  width: 50px;
  height: 50px;
}

.buttons {
  margin-top: 20px;
  padding-right: 35px;
  padding-left: 10px;
  text-align: left;
}

.navbar-title {
  margin-left: 10px;
}

.button {
  margin: auto;
  margin-bottom: 20px;
  width: 100%;
}

* {
    transition: background-color 200ms ease, color 150ms ease-in-out;
}

lu, li {
  list-style: none;
}

.fade-enter-active,
.fade-leave-active {
  transition-property: opacity;
  transition-duration: 0.4s;
}

.fade-enter-active {
  transition-delay: 0.4s;
}

.fade-enter,
.fade-leave-active {
  opacity: 0;
}

@media screen and (max-width: 960px) {
  #side-menu {
    width: 100%;
    position: fixed;
    transition: transform 300ms ease-in-out;
  }
  #router-content {
    margin-left: 0px;
  }
}
</style>
