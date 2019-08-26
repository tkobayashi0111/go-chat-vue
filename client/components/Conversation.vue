<template>
  <div class="conversation" :class="{ me: me }">
    <v-timeline :align-top="true" :dense="true" :reverse="me">
      <v-timeline-item>
        <template v-slot:icon>
          <v-avatar>
            <img :src="avatarSrc" />
          </v-avatar>
        </template>
        <v-card class="elevation-2">
          <v-card-title v-if="name">
            {{ name }}
          </v-card-title>
          <v-card-text>{{ message }}</v-card-text>
        </v-card>
      </v-timeline-item>
    </v-timeline>
  </div>
</template>

<style lang="scss">
.conversation {
  .v-timeline {
    padding-top: 12px;

    /* 縦線を消す */
    &::before {
      width: 0;
      height: 0;
    }

    .v-timeline-item {
      padding-bottom: 0;

      .v-timeline-item__body {
        display: flex;
      }
    }
  }

  &.me .v-timeline {
    .v-timeline-item {
      .v-timeline-item__body {
        justify-content: flex-end;
      }
    }
  }

  .v-card {
    min-width: 120px;

    .v-card__title {
      padding: 8px 16px;
      font-size: 0.9rem;
      line-height: 0.9rem;
      font-weight: normal;
      color: rgba(0, 0, 0, 0.54);
    }

    .v-card__text {
      font-size: 1.1rem;
      font-weight: normal;
      color: rgba(0, 0, 0, 0.87);
    }
  }

  &.me .v-card {
    .v-card__title {
      justify-content: flex-end;
    }

    .v-card__text {
      text-align: right;
    }
  }
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
