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
import { reactive,ref } from 'vue'
const data = reactive({
  form:{
  }
})
const rules = reactive({
  username: [
    { required: true, message: '请输入账号', trigger: 'blur' },
    { min: 3, max: 5, message: 'Length should be 3 to 5', trigger: 'blur' },
  ],
  password: [
    { required: true, message: '请输入密码', trigger: 'blur' },
    { min: 3, max: 5, message: 'Length should be 3 to 5', trigger: 'blur' },
  ],
})
const formRef = ref()
const login = () => {
  formRef.value.validate((valid) => {
    console.log(valid)
  }
)}

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