<template>
  <div>

    <div class="card" style="line-height: 30px">
      <div>欢迎您，{{ username }} 祝您今天过得开心！</div>
    </div>

  </div>
</template>

<script setup>
  import request from "@/utils/request";
  const username = JSON.parse(localStorage.getItem('user') || '{}')
</script>