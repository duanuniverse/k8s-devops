import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import Antd from 'ant-design-vue'
import 'ant-design-vue/dist/antd.dark.css';
import CButton from '@/components/CButton'
import * as Icons from '@ant-design/icons-vue'
//codemirror编辑器
import { GlobalCmComponent } from "codemirror-editor-vue3";
// 引入主题 可以从 codemirror/theme/ 下引入多个
import 'codemirror/theme/dracula.css'
// 引入语言模式 可以从 codemirror/mode/ 下引入多个
import 'codemirror/mode/yaml/yaml.js'


const app = createApp(App)
for (const i in Icons) {
    app.component(i, Icons[i])
  }
app.component('c-button', CButton)
app.use(GlobalCmComponent, { componentName: "codemirror" });
app.use(Antd)
app.use(router)
app.mount('#app')