<script setup lang="ts">
import { ref } from "vue";
import logo from "./assets/images/logo-universal.png";
import { Greet } from "../wailsjs/go/main/App.js";

const resultText = ref("Please enter your name below 👇");
const name = ref("");

function greet(): void {
  Greet(name.value).then((result) => (resultText.value = result));
}
</script>

<template>
  <main>
    <img alt="Wails logo" id="logo" :src="logo" />
    <div class="result" id="result">{{ resultText }}</div>
    <div class="input-box" id="input">
      <input autocomplete="off" v-model="name" class="input" id="name" type="text" />
      <button class="btn" @click="greet">Greet</button>
    </div>
  </main>
</template>

<style lang="less">
#logo {
  display: block;
  width: 50%;
  height: 50%;
  margin: auto;
  padding: 10% 0 0;
  background-position: center;
  background-repeat: no-repeat;
  background-size: 100% 100%;
  background-origin: content-box;
}

.result {
  height: 20px;
  line-height: 20px;
  margin: 1.5rem auto;
}

.input-box {
  .btn {
    width: 60px;
    height: 30px;
    line-height: 30px;
    border-radius: 3px;
    border: none;
    margin: 0 0 0 20px;
    padding: 0 8px;
    cursor: pointer;

    &:hover {
      background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
      color: #333333;
    }
  }

  .input {
    border: none;
    border-radius: 3px;
    outline: none;
    height: 30px;
    line-height: 30px;
    padding: 0 10px;
    background-color: rgba(240, 240, 240, 1);
    -webkit-font-smoothing: antialiased;

    &:hover {
      border: none;
      background-color: rgba(255, 255, 255, 1);
    }

    &:focus {
      border: none;
      background-color: rgba(255, 255, 255, 1);
    }
  }
}
</style>