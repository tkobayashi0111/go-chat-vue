<template>
  <v-layout column justify-center align-center>
    <v-flex xs12 sm8 md6>
      <v-text-field v-model="message" />
      <v-btn color="primary" @click="onClickSend">送信</v-btn>
      <v-list v-show="messages.length > 0">
        <v-list-item v-for="(message, i) in messages" :key="`message-${i}`">
          {{ message }}
        </v-list-item>
      </v-list>
    </v-flex>
  </v-layout>
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import io from 'socket.io-client'

@Component
export default class Index extends Vue {
  private socket!: SocketIOClient.Socket

  private message: string = ''

  private messages: string[] = []

  mounted() {
    this.socket = io('/chat')
    this.socket.on('chat', this.onReceiveChat)
  }

  onClickSend() {
    if (this.socket && this.socket.connected) {
      this.socket.emit('chat', this.message)
      this.message = ''
    }
  }

  onReceiveChat(message) {
    this.messages = [...this.messages, message]
  }
}
</script>
