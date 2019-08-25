<template>
  <v-row>
    <v-col lg="6" offset-lg="3" md="12">
      <div ref="contentArea" :style="contentAreaStyle">
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

  private contentArea: null | HTMLElement = null

  private contentAreaStyle: Partial<CSSStyleDeclaration> = {
    overflowY: 'auto'
  }

  mounted() {
    this.contentArea = this.$refs.contentArea as HTMLElement
    this.textarea = (this.$refs.textarea as Vue).$el as HTMLElement
    this.socket = io('/chat')
    this.socket.on('chat', this.onReceiveChat)

    window.addEventListener('resize', () => {
      this.updateContentAreaHeight()
    })
    this.updateContentAreaHeight()
  }

  onReceiveChat(message) {
    const shouldScroll = this.shouldScroll()
    this.messages = [...this.messages, message]
    if (shouldScroll) {
      this.$nextTick().then(() => {
        this.scrollBottom()
      })
    }
  }

  onSubmit() {
    if (this.message.length < 1) {
      return
    }

    this.socket.emit('chat', this.message)
    this.message = ''
  }

  @Watch('message')
  updateContentAreaHeight() {
    const shouldScroll = this.shouldScroll()
    const height = this.getContentAreaHeight()
    this.contentAreaStyle = {
      ...this.contentAreaStyle,
      ...{
        height
      }
    }
    if (shouldScroll) {
      this.$nextTick().then(() => {
        this.scrollBottom()
      })
    }
  }

  private getContentAreaHeight() {
    if (this.textarea) {
      const textareaHeight = this.getElementHeight(this.textarea)
      const containerPadding = 48
      const windowHeight = window.innerHeight
      return `${windowHeight - textareaHeight - containerPadding}px`
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

  private shouldScroll() {
    return this.getScrollBottom() === 0
  }

  private getScrollBottom() {
    if (this.contentArea) {
      return (
        this.contentArea.scrollHeight -
        this.contentArea.clientHeight -
        this.contentArea.scrollTop
      )
    } else {
      return 0
    }
  }

  private scrollBottom() {
    if (this.contentArea) {
      this.contentArea.scrollTop = this.contentArea.scrollHeight
    }
  }
}
</script>
