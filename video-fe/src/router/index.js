import Vue from 'vue'   //引入Vue
import Router from 'vue-router'  //引入vue-router
import Play from '../components/Play'  //引入根目录下的Hello.vue组件
import Login from '../components/Login'
import Upload from "@/components/Upload";

Vue.use(Router)  //Vue全局使用Router

const routes = [              //配置路由，这里是个数组
    {
        path: '/',
        redirect: '/login',
    },
    {                    //每一个链接都是一个对象
        path: '/play',         //链接路径
        name: 'play',     //路由名称，
        component: Play   //对应的组件模板
    },
    {                    //每一个链接都是一个对象
        path: '/login',         //链接路径
        name: 'login',     //路由名称，
        component: Login   //对应的组件模板
    },
    {                    //每一个链接都是一个对象
        path: '/upload',         //链接路径
        name: 'upload',     //路由名称，
        component: Upload   //对应的组件模板
    }
    // {
    //     path: '/hi',
    //     component: Hi,
    //     children:
    //         [        //子路由,嵌套路由 （此处偷个懒，免得单独再列一点）
    //             {path: '/', component: Hi},
    //             {path: 'hi1', component: Hi1},
    //             {path: 'hi2', component: Hi2},
    //         ]
    // }
]
const router = new Router({
    routes // routes: routes 的简写
})

// 导航守卫
// 使用 router.beforeEach 注册一个全局前置守卫，判断用户是否登陆
router.beforeEach((to, from, next) => {
    if (to.path === '/login') {
        next();
    } else {
        let token = localStorage.getItem('Authorization');
        if (token === null || token === '') {
            console.log("当前未登录")
            next('/login');
        } else {
            next();
        }
    }
});
export default router
