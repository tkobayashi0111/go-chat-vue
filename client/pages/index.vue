<template>
  <v-row>
    <v-col lg="6" offset-lg="3" md="12">
      <div>
        <template v-for="(message, i) in messages">
          <conversation :key="`message-${i}`" :message="message" />
        </template>
      </div>
      <v-textarea
        v-model="message"
        append-icon="mdi-send"
        auto-grow
        autofocus
        rows="1"
        @keydown.meta.enter="onSubmit"
        @keydown.ctrl.enter="onSubmit"
        @click:append="onSubmit"
      />
    </v-col>
  </v-row>
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import io from 'socket.io-client'

import Conversation from '@/components/Conversation.vue'

@Component({
  components: {
    Conversation
  }
})
export default class Index extends Vue {
  private socket!: SocketIOClient.Socket

  private message: string = ''

  private messages: string[] = []

  mounted() {
    this.socket = io('/chat')
    this.socket.on('chat', this.onReceiveChat)
  }

  onReceiveChat(message) {
    this.messages = [...this.messages, message]
  }

  onSubmit() {
    if (this.message.length < 1) {
      return
    }

    this.socket.emit('chat', this.message)
    this.message = ''
  }
}
</script>
