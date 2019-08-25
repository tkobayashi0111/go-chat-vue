<template>
  <v-timeline :align-top="true" :dense="true" :reverse="me">
    <v-timeline-item>
      <template v-slot:icon>
        <v-avatar>
          <img :src="avatarSrc" />
        </v-avatar>
      </template>
      <v-card class="elevation-2">
        <v-card-title v-if="name" class="headline">{{ name }}</v-card-title>
        <v-card-text>{{ message }}</v-card-text>
      </v-card>
    </v-timeline-item>
  </v-timeline>
</template>

<style lang="scss" scoped>
.v-timeline::before {
  width: 0;
  height: 0;
}
</style>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'

@Component
export default class Conversation extends Vue {
  @Prop({
    type: String,
    required: true
  })
  private readonly message!: string

  @Prop({
    type: String,
    default: null
  })
  private readonly name!: string

  @Prop({
    type: String,
    default: 'http://i.pravatar.cc/64'
  })
  private readonly avatarSrc!: string

  @Prop({
    type: Boolean,
    default: false
  })
  private readonly me!: boolean
}
</script>
