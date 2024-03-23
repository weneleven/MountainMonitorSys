<template>
  <div >
    <div class="login-container">
      <div class='login-box'>
        <div style="text-align: center;margin-bottom: 30px;font-weight: bold;font-size: 24px">登 录</div>
        <el-form :model="data.form"
                 ref="formRef"
                 :rules="rules">
          <el-form-item prop="username">
            <el-input prefix-icon="User" v-model="data.form.username" placeholder="请输入账号"/>
          </el-form-item>
          <el-form-item prop="password">
            <el-input prefix-icon="Lock" v-model="data.form.password" placeholder="请输入密码"  show-password/>
          </el-form-item>
          <el-form-item >
            <el-button type="primary" style="width: 100%" @click="login">登录</el-button>
          </el-form-item>
        </el-form>
        <div style="text-align: right;font-size: small">
          还没有账号？请<a href="/register">注册</a>
        </div>
      </div>
    </div>
  </div>
</template>


<script setup>
import { reactive, ref } from 'vue'
import axios from 'axios'
import { useRouter } from 'vue-router'

const data = reactive({
  form: {}
})

const rules = reactive({
  username: [
    { required: true, message: '请输入账号', trigger: 'blur' },
  ],
  password: [
    { required: true, message: '请输入密码', trigger: 'blur' },
  ],
})

const formRef = ref()
const router = useRouter()
const login = () => {
  formRef.value.validate(async (valid) => {
    if (valid) {
      try {
        // 发送登录请求到后端 API
        const response = await axios.post('http://124.70.83.36:3000/api/v1/login', data.form)

        // 根据后端返回的状态码处理逻辑
        if (response.data.status === 200) {
          // 登录成功，提示并跳转到其他页面
          ElMessage({
            message: response.data.message,
            type: 'success',
          })
          //存储登录成功后的username到本地，之后使用
          localStorage.setItem('user', JSON.stringify(response.data.name));
          //存储登录成功后的token
          localStorage.setItem('token',JSON.stringify(response.data.token))
          router.push('/manager')
        } else {
          // 登录失败，给出提示
          ElMessage.error(response.data.message || 'Login failed');
          console.error('Login failed:', response.data.message)
        }
      } catch (error) {
        console.error('Error during login:', error)
      }
    }
  })
}
</script>

<style scoped>
.login-container{
  min-height: 100vh;
  overflow: hidden;
  display: flex;
  align-items: center;
  justify-content: center;
  background-image: url("../assets/imgs/bg.png");
  background-size: cover;
}
.login-box{
  width: 350px;
  background-color: #ffffff;
  box-shadow: 0 0 10px rgba(0,0,0,0.1);
  padding: 30px;
  border-radius: 5px;
  background-color: rgba(255,255,255,.8);
}
</style>