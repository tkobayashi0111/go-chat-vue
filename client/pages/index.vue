<template>
  <v-row>
    <v-col lg="6" offset-lg="3" md="12">
      <div :style="{ overflowY: 'auto', height: getContentAreaHeight() }">
        <template v-for="(message, i) in messages">
          <conversation :key="`message-${i}`" :message="message" />
        </template>
      </div>
      <v-textarea
        ref="textarea"
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
import { Vue, Component, Watch } from 'vue-property-decorator'
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

  private textarea: null | HTMLElement = null

  mounted() {
    this.socket = io('/chat')
    this.socket.on('chat', this.onReceiveChat)
    this.textarea = (this.$refs.textarea as Vue).$el as HTMLElement
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

  @Watch('message')
  private getContentAreaHeight() {
    if (this.textarea) {
      const textareaHeight = this.getElementHeight(this.textarea)
      const containerPadding = 48
      return `calc(100vh - ${textareaHeight}px - ${containerPadding}px)`
    } else {
      return '0'
    }
  }

  private getElementHeight(el: HTMLElement) {
    const style = window.getComputedStyle(el)
    const margin =
      parseInt(style.marginTop || '0') + parseInt(style.marginBottom || '0')
    return el.offsetHeight + margin
  }
}
</script>
