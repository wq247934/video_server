import Vue from 'vue'
import Vuex from 'vuex'//导入vuex，先导入Vue
Vue.use(Vuex)//将导入的vuex安装到项目中

//创建Vuex实例并将它暴露出去给其它页面使用
export default new Vuex.Store({
    // state提供唯一的公共数据源，所有共享的数据都要同意放到store的state中进行存储
    state:{
        promptMessage: false, // home中定义全局snackbar
        snackbarInfo: null, // home中全局snackbar提示信息
    },
    getters:{},
    //使用mutations变更store中的数据以及传递参数，可以集中监听所有数据的变化
    //只能通过mutations变更store数据，不可以直接操作store中的数据
    mutations:{
        //定义个函数，传入形参state，形参就是state:{}
        //全局消息回执后的snackbar提示
        showsnackbar(state, message) {
            state.promptMessage = true;
            state.snackbarInfo = message;
        },
    },
    actions:{}
})
