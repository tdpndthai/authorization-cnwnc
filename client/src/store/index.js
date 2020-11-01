import Vue from 'vue';
import Vuex from 'vuex';
import user from './user/user'
import * as getters from "./getters.js";
import * as actions from "./actions.js";
import * as mutations from "./mutations.js";

Vue.use(Vuex);

export default new Vuex.Store({
    state: {},
    getters,
    mutations,
    actions,
    modules: {
        user
    }
})