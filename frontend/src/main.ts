
// src/main.ts
import { createApp } from 'vue'
import App from './App.vue'
import { router } from './router'

import {
    create,
    NConfigProvider,
    NMessageProvider,
    NButton,
    NLayout,
    NLayoutHeader,
    NLayoutContent,
    NLayoutFooter,
    NMenu,
    NIcon,
    NCard,
    NTag,
    NSpace,
    NDivider,
    NText
} from 'naive-ui'

const naive = create({
    components: [
        NConfigProvider,
        NMessageProvider,
        NButton,
        NLayout,
        NLayoutHeader,
        NLayoutContent,
        NLayoutFooter,
        NMenu,
        NIcon,
        NCard,
        NTag,
        NSpace,
        NDivider,
        NText
    ]
})

const app = createApp(App)
app.use(router)
app.use(naive)
app.mount('#app')