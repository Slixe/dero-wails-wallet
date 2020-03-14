import VueRouter from 'vue-router'
import Vue from 'vue'
import Recover from './views/Recover.vue'
import New from './views/New.vue'
import About from './views/About.vue'
import Send from './views/Send.vue'
import Receive from './views/Receive.vue'
import Home from './views/Home.vue'
import History from './views/History.vue'
import Seed from './views/recover/Seed.vue'

Vue.use(VueRouter)

const routes = [
    { path: '/', component: Recover },
    { path: '/new', component: New },
    { path: '/about', component: About },
    { path: '/send', component: Send },
    { path: '/receive', component: Receive },
    { path: '/home', component: Home },
    { path: '/history', component: History },

    { path: '/recover-seed', component: Seed }
  ];

export default new VueRouter({
    mode: 'hash',
    routes
});