<template>
  <v-row align="center">
    <v-col lg="6" offset-lg="3" md="12">
      <v-form ref="form" v-model="valid" :lazy-validation="true">
        <v-text-field v-model="name" :rules="nameRules" label="Name" required />
        <v-btn color="primary" @click="submit">Enter</v-btn>
      </v-form>
    </v-col>
  </v-row>
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import VForm from 'vuetify/es5/components/VForm'

@Component
export default class Enter extends Vue {
  private valid: boolean = false

  private name: string = ''

  private nameRules = [(v) => !!v || 'Name is required']

  submit() {
    if ((this.$refs.form as VForm).validate()) {
      this.$router.replace({
        name: 'index',
        params: {
          name: this.name
        }
      })
    }
  }
}
</script>
