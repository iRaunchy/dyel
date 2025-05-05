import {createRouter, createWebHistory, RouteRecordRaw} from 'vue-router'
import ProgramList from '../components/ProgramList.vue'
import ProgramDetail from '../components/ProgramDetail.vue'
import ProgramCreate from '../components/ProgramCreate.vue'

const routes: RouteRecordRaw[] = [
    {path: '/', redirect: '/programs'},
    {path: '/programs', component: ProgramList},
    {path: '/programs/new', component: ProgramCreate},
    {path: '/programs/:id', component: ProgramDetail, props: true},
]

export const router = createRouter({
    history: createWebHistory(),
    routes,
})
